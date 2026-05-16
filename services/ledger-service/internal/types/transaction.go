package types

import "time"

type Transaction struct {
	ID          string    `json:"id,omitempty"`
	UserID      string    `json:"user_id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
