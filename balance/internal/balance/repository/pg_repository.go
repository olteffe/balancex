package repository

import "github.com/jackc/pgx/v4/pgxpool"

// BalanceRepository Images Balance Repository
type BalanceRepository struct {
	db *pgxpool.Pool
}

// NewBalanceRepository Images repository constructor
func NewBalanceRepository(db *pgxpool.Pool) *BalanceRepository {
	return &BalanceRepository{db: db}
}
