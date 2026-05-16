package main

import (
    "context"
    "encoding/json"
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/shivamyadav0/vani-ledger-platform/pkg/logger"
    "github.com/shivamyadav0/vani-ledger-platform/pkg/messaging"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/segmentio/kafka-go"
    "go.uber.org/zap"
)

type transactionMessage struct {
    ID          string    `json:"id,omitempty"`
    UserID      string    `json:"user_id"`
    Amount      float64   `json:"amount"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
}

func main() {
    log, err := logger.NewLogger()
    if err != nil {
        panic(err)
    }
    defer log.Sync()

    brokers := envOrDefault("KAFKA_BROKERS", "kafka:9092")
    pgConn := envOrDefault("PG_CONN", "postgresql://ledger:ledgerpass@pgbouncer:6432/ledger?sslmode=disable")

    pool, err := pgxpool.New(context.Background(), pgConn)
    if err != nil {
        log.Fatal("failed to connect postgres", zap.Error(err))
    }
    defer pool.Close()

    consumer := messaging.NewConsumer([]string{brokers}, "batch-worker", "transactions")
    buffer := make([]kafka.Message, 0, 1000)
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            if len(buffer) > 0 {
                flushBatch(context.Background(), log, pool, buffer)
                buffer = buffer[:0]
            }
        default:
            msg, err := consumer.ReadMessage(context.Background())
            if err != nil {
                log.Error("kafka read failed", zap.Error(err))
                time.Sleep(time.Second)
                continue
            }
            buffer = append(buffer, msg)
            if len(buffer) >= 1000 {
                flushBatch(context.Background(), log, pool, buffer)
                _ = consumer.CommitMessages(context.Background(), buffer...)
                buffer = buffer[:0]
            }
        }
    }
}

func flushBatch(ctx context.Context, log *zap.Logger, pool *pgxpool.Pool, messages []kafka.Message) {
    txs := make([]transactionMessage, 0, len(messages))
    for _, msg := range messages {
        var txn transactionMessage
        if err := json.Unmarshal(msg.Value, &txn); err != nil {
            log.Warn("failed to unmarshal kafka message", zap.Error(err))
            continue
        }
        txs = append(txs, txn)
    }
    if len(txs) == 0 {
        return
    }

    builder := make([]string, 0, len(txs))
    args := make([]interface{}, 0, len(txs)*5)
    for i, txn := range txs {
        base := i*5 + 1
        builder = append(builder, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d)", base, base+1, base+2, base+3, base+4))
        args = append(args, txn.ID, txn.UserID, txn.Amount, txn.Description, txn.CreatedAt)
    }

    query := fmt.Sprintf("INSERT INTO transactions (id, user_id, amount, description, created_at) VALUES %s ON CONFLICT DO NOTHING", strings.Join(builder, ","))
    if _, err := pool.Exec(ctx, query, args...); err != nil {
        log.Error("batch insert failed", zap.Error(err))
        return
    }

    log.Info("batch inserted transactions", zap.Int("count", len(txs)))
}

func envOrDefault(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
