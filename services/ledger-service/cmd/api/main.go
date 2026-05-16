package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"github.com/shivamyadav0/vani-ledger-platform/pkg/logger"
	"github.com/shivamyadav0/vani-ledger-platform/pkg/messaging"
	"github.com/shivamyadav0/vani-ledger-platform/services/ledger-service/internal/cache"
	"github.com/shivamyadav0/vani-ledger-platform/services/ledger-service/internal/ledger"
	"github.com/shivamyadav0/vani-ledger-platform/services/ledger-service/internal/repository"
	"github.com/shivamyadav0/vani-ledger-platform/services/ledger-service/internal/voice"
	"go.uber.org/zap"
)

type transactionRequest struct {
	UserID      string   `json:"user_id"`
	Amount      *float64 `json:"amount,omitempty"`
	Description string   `json:"description,omitempty"`
	Text        string   `json:"text,omitempty"`
	AudioBase64 string   `json:"audio_base64,omitempty"`
}

type apiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	redisAddr := envOrDefault("REDIS_ADDR", "redis:6379")
	kafkaBrokers := envOrDefault("KAFKA_BROKERS", "kafka:9092")
	pgConn := envOrDefault("PG_CONN", "postgresql://ledger:ledgerpass@pgbouncer:6432/ledger?sslmode=disable")

	redisClient := cache.NewRedis(redisAddr)
	brokers := strings.Split(kafkaBrokers, ",")

	// --- Pre-warm & Ensure Topic Initialization ---
	voiceRequestTopic := envOrDefault("VOICE_REQUEST_TOPIC", "voice-requests")
	voiceResponseTopic := envOrDefault("VOICE_RESPONSE_TOPIC", "voice-results")

	log.Info("syncing kafka topics metadata...")
	topics := []string{"transactions", voiceRequestTopic, voiceResponseTopic}
	for _, t := range topics {
		if err := messaging.EnsureTopicExists(brokers, t, 1, 1); err != nil {
			log.Fatal("failed to verify/create core kafka topics", zap.String("topic", t), zap.Error(err))
		}
	}

	producer, err := messaging.NewProducer(brokers, "transactions")
	if err != nil {
		log.Fatal("failed to initialize kafka producer", zap.Error(err))
	}
	defer producer.Close()

	repo, err := repository.NewRepository(pgConn)
	if err != nil {
		log.Fatal("failed to connect postgres", zap.Error(err))
	}

	ledgerService := ledger.NewLedgerService(redisClient, producer, repo, log)

	voiceProducer, err := messaging.NewProducer(brokers, voiceRequestTopic)
	if err != nil {
		log.Fatal("failed to initialize voice producer", zap.Error(err))
	}
	defer voiceProducer.Close()

	voiceClient := voice.NewClient(voiceProducer, log)

	// Async consumer routine initialized safely
	go consumeVoiceResults(redisClient, brokers, voiceResponseTopic, ledgerService, log)

	r := chi.NewRouter()
	r.Post("/v1/transaction", makePostHandler(log, ledgerService, voiceClient))
	r.Get("/v1/voice-status/{requestID}", makeVoiceStatusHandler(redisClient))
	r.Get("/v1/ledger/{userID}", makeGetHandler(log, ledgerService))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := ":8080"
	log.Info("starting ledger service", zap.String("addr", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("http server failed", zap.Error(err))
	}
}

func consumeVoiceResults(
	redisClient *redis.Client,
	brokers []string,
	topic string,
	service *ledger.LedgerService,
	log *zap.Logger,
) {
	log.Info("voice consumer goroutine started", zap.String("topic", topic))

	groupID := "ledger-voice-results-consumer"
	
	// Create consumer with explicit FirstOffset fallback configuration
	consumer := messaging.NewConsumerWithOffset(brokers, groupID, topic, kafka.FirstOffset)
	
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Error("failed to close voice consumer", zap.Error(err))
		}
	}()

	for {
		log.Info("waiting for voice results...")
		msg, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Error("failed reading kafka message, refreshing context connection...", zap.Error(err))
			_ = consumer.Close()
			
			time.Sleep(3 * time.Second)
			consumer = messaging.NewConsumerWithOffset(brokers, groupID, topic, kafka.FirstOffset)
			continue
		}

		log.Info("received kafka voice result",
			zap.String("topic", msg.Topic),
			zap.Int("partition", msg.Partition),
			zap.Int64("offset", msg.Offset),
		)

		var response voice.VoiceResultResponse
		if err := json.Unmarshal(msg.Value, &response); err != nil {
			log.Error("failed unmarshalling voice response", zap.Error(err), zap.ByteString("payload", msg.Value))
			continue
		}

		key := fmt.Sprintf("voice:request:%s", response.RequestID)
		payload, err := json.Marshal(response)
		if err != nil {
			log.Error("failed marshaling redis payload", zap.Error(err))
			continue
		}

		if err := redisClient.Set(context.Background(), key, payload, 10*time.Minute).Err(); err != nil {
			log.Error("failed storing voice result in redis", zap.Error(err))
			continue
		}

		if response.Status != "completed" || response.Intent != "transaction.create" {
			log.Warn("skipping execution: incomplete or unsupported voice intent", zap.String("intent", response.Intent))
			continue
		}

		if response.Data.Amount <= 0 {
			log.Warn("invalid transaction amount", zap.Float64("amount", response.Data.Amount))
			continue
		}

		transaction := ledger.Transaction{
			ID:          generateID(),
			UserID:      response.UserID,
			Amount:      response.Data.Amount,
			Description: response.Data.Description,
			CreatedAt:   time.Now().UTC(),
		}

		if err := service.CreateTransaction(context.Background(), transaction); err != nil {
			log.Error("failed creating transaction from voice result", zap.Error(err))
			continue
		}

		log.Info("transaction created cleanly from voice engine payload", zap.String("transaction_id", transaction.ID))
	}
}

func makePostHandler(log *zap.Logger, service *ledger.LedgerService, voiceClient *voice.VoiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req transactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondJSON(w, http.StatusBadRequest, apiResponse{Success: false, Error: "invalid payload"})
			return
		}

		if req.UserID == "" {
			respondJSON(w, http.StatusBadRequest, apiResponse{Success: false, Error: "missing user_id"})
			return
		}

		if req.Amount != nil && req.Description != "" {
			txn := ledger.Transaction{
				ID:          generateID(),
				UserID:      req.UserID,
				Amount:      *req.Amount,
				Description: req.Description,
				CreatedAt:   time.Now().UTC(),
			}

			if err := service.CreateTransaction(r.Context(), txn); err != nil {
				respondJSON(w, http.StatusInternalServerError, apiResponse{Success: false, Error: "failed to persist transaction"})
				return
			}
			respondJSON(w, http.StatusOK, apiResponse{Success: true, Data: txn})
			return
		}

		requestID, err := voiceClient.Process(r.Context(), req.UserID, req.Text, req.AudioBase64)
		if err != nil {
			respondJSON(w, http.StatusBadGateway, apiResponse{Success: false, Error: err.Error()})
			return
		}

		respondJSON(w, http.StatusAccepted, apiResponse{Success: true, Data: map[string]interface{}{"request_id": requestID, "status": "processing"}})
	}
}

func makeVoiceStatusHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := chi.URLParam(r, "requestID")
		key := fmt.Sprintf("voice:request:%s", requestID)

		result, err := redisClient.Get(r.Context(), key).Result()
		if err == redis.Nil {
			respondJSON(w, http.StatusAccepted, apiResponse{Success: true, Data: map[string]interface{}{"status": "processing"}})
			return
		}
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, apiResponse{Success: false, Error: "redis error"})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(result))
	}
}

func makeGetHandler(log *zap.Logger, service *ledger.LedgerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		rows, err := service.GetLedger(r.Context(), userID)
		if err != nil {
			respondJSON(w, http.StatusInternalServerError, apiResponse{Success: false, Error: "failed to read ledger"})
			return
		}
		respondJSON(w, http.StatusOK, apiResponse{Success: true, Data: rows})
	}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func envOrDefault(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func generateID() string {
	buffer := make([]byte, 16)
	_, _ = rand.Read(buffer)
	return hex.EncodeToString(buffer)
}