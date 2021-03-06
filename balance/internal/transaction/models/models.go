package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID            uint64    `db:"id" validate:"omitempty"`
	TransactionID uuid.UUID `db:"transaction_id" validate:"required,uuid"`
	Source        string    `db:"source" validate:"required,gt=0,max=250"`
	Description   string    `db:"description" validate:"required,gt=0,max=250"`
	SenderID      uuid.UUID `db:"sender_id" validate:"required,uuid"`
	RecipientID   uuid.UUID `db:"recipient_id" validate:"required,uuid"`
	Currency      string    `db:"currency" validate:"required,len=3,uppercase"`
	Amount        int64     `db:"amount" validate:"required,numeric"`
	CreatedAt     time.Time `db:"created_at" validate:"omitempty"`
}

type TransactionList struct {
	TotalCount   uint32         `db:"total_count" validate:"omitempty,gt=0"`
	TotalPages   uint32         `db:"total_page" validate:"omitempty,gt=0"`
	Page         uint32         `db:"page" validate:"omitempty,gt=0"`
	Size         uint32         `db:"size" validate:"omitempty,gt=0"`
	HasMore      bool           `db:"has_more" validate:"omitempty,gt=0"`
	Transactions []*Transaction `validate:"omitempty"`
}
