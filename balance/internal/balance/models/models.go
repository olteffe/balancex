package models

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        uint64
	UserID    uuid.UUID
	Currency  string
	Amount    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID            uint64
	TransactionID uuid.UUID
	Source        string
	Description   string
	SenderID      uuid.UUID
	RecipientID   uuid.UUID
	Currency      string
	Amount        int64
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type TransactionList struct {
	TotalCount   int64
	TotalPages   int64
	Page         int64
	Size         int64
	HasMore      bool
	Transactions []*Transaction
}
