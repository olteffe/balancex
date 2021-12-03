package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"

	"github.com/olteffe/balancex/balance/internal/balance/models"
)

// BalanceRepository Images Balance Repository
type BalanceRepository struct {
	db *pgxpool.Pool
}

// NewBalanceRepository Images repository constructor
func NewBalanceRepository(db *pgxpool.Pool) *BalanceRepository {
	return &BalanceRepository{db: db}
}

// CreateBalance Create Balance
func (r *BalanceRepository) CreateBalance(ctx context.Context, user *models.Balance) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BalanceRepository.Create")
	defer span.Finish()

	var createdBalance string
	if err := r.db.QueryRow(
		ctx,
		createBalanceQuery,
		user.UserID,
		user.Currency,
		user.Amount,
	).Scan(&createdBalance); err != nil {
		return "", fmt.Errorf("CreateBalance.QueryRowContext: %w", err)
	}
	return createdBalance, nil
}

// GetBalance get balance
func (r *BalanceRepository) GetBalance(ctx context.Context, user *models.Balance) (*models.Balance, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "BalanceRepository.Get")
	defer span.Finish()

	var gotBalance models.Balance
	if err := r.db.QueryRow(ctx, getBalanceQuery, user.UserID).Scan(
		&gotBalance.UserID,
		&gotBalance.Currency,
		&gotBalance.Amount,
		&gotBalance.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("GetBalance.QueryRowContext: %w", err)
	}
	return &gotBalance, nil
}
