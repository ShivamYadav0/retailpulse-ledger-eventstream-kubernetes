package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "time"

    "github.com/shivamyadav0/vani-ledger-platform/pkg/logger"
    "github.com/shivamyadav0/vani-ledger-platform/pkg/messaging"
    "go.uber.org/zap"
)

type notificationMessage struct {
    UserID      string  `json:"user_id"`
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
}

func main() {
    log, err := logger.NewLogger()
    if err != nil {
        panic(err)
    }
    defer log.Sync()

    brokers := envOrDefault("KAFKA_BROKERS", "kafka:9092")
    consumer := messaging.NewConsumer([]string{brokers}, "notification-service", "transactions")

    for {
        msg, err := consumer.ReadMessage(context.Background())
        if err != nil {
            log.Error("kafka read failed", zap.Error(err))
            time.Sleep(time.Second)
            continue
        }

        var payload notificationMessage
        if err := json.Unmarshal(msg.Value, &payload); err != nil {
            log.Warn("invalid kafka payload", zap.Error(err))
            continue
        }

        log.Info("notification queued", zap.String("user_id", payload.UserID), zap.Float64("amount", payload.Amount))
        fmt.Printf("sending notification for user %s: %0.2f %s\n", payload.UserID, payload.Amount, payload.Description)
        _ = consumer.CommitMessages(context.Background(), msg)
    }
}

func envOrDefault(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
