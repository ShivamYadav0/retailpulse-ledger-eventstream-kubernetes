package repository

import (
    "context"
    "fmt"
    "strings"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/shivamyadav0/vani-ledger-platform/services/ledger-service/internal/types"
)

type PostgresRepository struct {
    pool *pgxpool.Pool
}

func NewRepository(conn string) (*PostgresRepository, error) {
    pool, err := pgxpool.New(context.Background(), conn)
    if err != nil {
        return nil, err
    }
    return &PostgresRepository{pool: pool}, nil
}

func (r *PostgresRepository) GetTransactions(ctx context.Context, userID string) ([]types.Transaction, error) {
    rows, err := r.pool.Query(ctx, `SELECT id, user_id, amount, description, created_at FROM transactions WHERE user_id=$1 ORDER BY created_at DESC LIMIT 100`, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []types.Transaction
    for rows.Next() {
        var tx types.Transaction
        if err := rows.Scan(&tx.ID, &tx.UserID, &tx.Amount, &tx.Description, &tx.CreatedAt); err != nil {
            return nil, err
        }
        results = append(results, tx)
    }
    return results, nil
}

func (r *PostgresRepository) BulkInsert(ctx context.Context, txs []types.Transaction) error {
    if len(txs) == 0 {
        return nil
    }

    builder := strings.Builder{}
    builder.WriteString("INSERT INTO transactions (id, user_id, amount, description, created_at) VALUES ")
    params := make([]interface{}, 0, len(txs)*5)

    for i, txn := range txs {
        if i > 0 {
            builder.WriteString(",")
        }
        idx := i*5 + 1
        builder.WriteString(fmt.Sprintf("($%d,$%d,$%d,$%d,$%d)", idx, idx+1, idx+2, idx+3, idx+4))
        params = append(params, txn.ID, txn.UserID, txn.Amount, txn.Description, txn.CreatedAt)
    }

    _, err := r.pool.Exec(ctx, builder.String(), params...)
    return err
}
