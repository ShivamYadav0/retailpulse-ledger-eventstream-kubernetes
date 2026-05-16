package ledger

import (
    "context"
    "encoding/json"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/shivamyadav0/vani-ledger-platform/services/ledger-service/internal/types"
    "github.com/shivamyadav0/vani-ledger-platform/pkg/messaging"
    "go.uber.org/zap"
)

// Re-export Transaction for convenience
type Transaction = types.Transaction

type LedgerService struct {
    cache    *redis.Client
    producer *messaging.KafkaProducer
    repo     interface {
        GetTransactions(ctx context.Context, userID string) ([]types.Transaction, error)
    }
    logger *zap.Logger
}

func NewLedgerService(cacheClient *redis.Client, producer *messaging.KafkaProducer, repo interface {
    GetTransactions(ctx context.Context, userID string) ([]types.Transaction, error)
}, log *zap.Logger) *LedgerService {
    return &LedgerService{cache: cacheClient, producer: producer, repo: repo, logger: log}
}

func (s *LedgerService) CreateTransaction(ctx context.Context, tx types.Transaction) error {
    key := fmt.Sprintf("ledger:%s", tx.UserID)
    payload, err := json.Marshal(tx)
    if err != nil {
        return err
    }

    if err := s.cache.LPush(ctx, key, payload).Err(); err != nil {
        s.logger.Warn("redis write failed", zap.Error(err), zap.String("user_id", tx.UserID))
        return err
    }
    if err := s.cache.LTrim(ctx, key, 0, 999).Err(); err != nil {
        s.logger.Warn("redis trim failed", zap.Error(err), zap.String("user_id", tx.UserID))
    }

    if err := s.producer.Produce(ctx, tx.UserID, payload); err != nil {
        s.logger.Error("kafka produce failed", zap.Error(err), zap.String("user_id", tx.UserID))
        return err
    }
    s.logger.Info("transaction created", zap.String("user_id", tx.UserID), zap.Float64("amount", tx.Amount))
    return nil
}

func (s *LedgerService) GetLedger(ctx context.Context, userID string) ([]types.Transaction, error) {
    key := fmt.Sprintf("ledger:%s", userID)
    raw, err := s.cache.LRange(ctx, key, 0, 99).Result()
    if err == nil && len(raw) > 0 {
        return s.unmarshalTransactions(raw)
    }
    if err != nil && err != redis.Nil {
        s.logger.Warn("redis read failed", zap.Error(err), zap.String("user_id", userID))
    }

    if s.repo == nil {
        return nil, fmt.Errorf("no repository available")
    }
    txs, err := s.repo.GetTransactions(ctx, userID)
    if err != nil {
        return nil, err
    }
    if len(txs) == 0 {
        return txs, nil
    }

    pipeline := s.cache.Pipeline()
    for _, tx := range txs {
        data, _ := json.Marshal(tx)
        pipeline.RPush(ctx, key, data)
    }
    pipeline.Expire(ctx, key, 5*time.Minute)
    _, _ = pipeline.Exec(ctx)

    return txs, nil
}

func (s *LedgerService) unmarshalTransactions(raw []string) ([]types.Transaction, error) {
    txs := make([]types.Transaction, 0, len(raw))
    for _, item := range raw {
        var tx types.Transaction
        if err := json.Unmarshal([]byte(item), &tx); err != nil {
            return nil, err
        }
        txs = append(txs, tx)
    }
    return txs, nil
}
