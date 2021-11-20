package models

import (
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID        uint64    `db:"id" validate:"omitempty"`
	UserID    uuid.UUID `db:"user_id" validate:"required,uuid"`
	Currency  string    `db:"currency" validate:"required,len=3,uppercase"`
	Amount    uint64    `db:"amount" validate:"required,numeric,gte=0"`
	CreatedAt time.Time `db:"created_at" validate:"omitempty"`
	UpdatedAt time.Time `db:"updated_at" validate:"omitempty"`
}
