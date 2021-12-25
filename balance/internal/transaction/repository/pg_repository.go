package repository

import (
	"context"
	"fmt"
	"github.com/olteffe/balancex/balance/pkg/utils"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/opentracing/opentracing-go"

	"github.com/olteffe/balancex/balance/internal/transaction/models"
	gErrors "github.com/olteffe/balancex/balance/pkg/grpc_errors"
)

// TransactionRepository Images Balance Repository
type TransactionRepository struct {
	db *pgxpool.Pool
}

// NewTransactionRepository repository constructor
func NewTransactionRepository(db *pgxpool.Pool) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// FindUsersID find SenderID and RecipientID in repository
func (r *TransactionRepository) FindUsersID(ctx context.Context, transaction *models.Transaction) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TransactionRepository.FindUsersID")
	defer span.Finish()

	var count int64
	if err := r.db.QueryRow(
		ctx,
		findUsersIDQuery,
		transaction.SenderID,
		transaction.RecipientID,
	).Scan(&count); err != nil {
		return false, fmt.Errorf("FindUsersID.QueryRowContext: %w", err)
	}
	if count != 2 {
		return false, fmt.Errorf("FindUsersID.CountUsers: %w", gErrors.ErrUserExists)
	}
	return true, nil
}

// CreateTransaction Create Transaction
func (r *TransactionRepository) CreateTransaction(ctx context.Context, transaction *models.Transaction) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TransactionRepository.CreateTransaction")
	defer span.Finish()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("CreateTransaction.Begin: %w", err)
	}
	defer tx.Rollback(ctx)
	decrease, err := tx.Exec(ctx, createTransactionDecreaseQuery, transaction.Amount, transaction.SenderID)
	if err != nil {
		return "", fmt.Errorf("CreateTransaction.Decrease: %w", err)
	}
	if decrease.RowsAffected() == 0 {
		return "", fmt.Errorf("CreateTransaction.DecreaseRowsAffected: %w", err)
	}

	increase, err := tx.Exec(ctx, createTransactionIncreaseQuery, transaction.Amount, transaction.RecipientID)
	if err != nil {
		return "", fmt.Errorf("CreateTransaction.Increase: %w", err)
	}
	if increase.RowsAffected() == 0 {
		return "", fmt.Errorf("CreateTransaction.IncreaseRowsAffected: %w", err)
	}

	var transactionID string
	if err := tx.QueryRow(
		ctx,
		createTransactionHistoryQuery,
		transaction.TransactionID,
		transaction.Source,
		transaction.Description,
		transaction.SenderID,
		transaction.RecipientID,
		transaction.Amount,
		transaction.CreatedAt,
	).Scan(&transactionID); err != nil {
		return "", fmt.Errorf("CreateTransaction.History: %w", err)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return "", fmt.Errorf("CreateTransaction.Commit: %w", err)
	}
	return transactionID, nil
}

// GetTransactions Get Transactions
func (r *TransactionRepository) GetTransactions(ctx context.Context, transaction *utils.TransactionsRequest) (list *models.TransactionList, err error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "TransactionRepository.GetTransactions")
	defer span.Finish()

	var total uint32
	if err := r.db.QueryRow(ctx, getTransactionsCountQuery).Scan(&total); err != nil {
		return nil, fmt.Errorf("GetTransactions.Count: %w", err)
	}
	if total == 0 {
		return &models.TransactionList{
			TotalCount:   total,
			TotalPages:   0,
			Page:         0,
			Size:         0,
			HasMore:      false,
			Transactions: make([]*models.Transaction, 0),
		}, nil
	}

	rows, err := r.db.Query(ctx, getTransactionsQuery, transaction.UserID, transaction.GetOffset(), transaction.GetLimit())
	if err != nil {
		return nil, fmt.Errorf("GetTransactions.Rows: %w", err)
	}
	defer rows.Close()

	transactions := make([]*models.Transaction, 0, transaction.GetLimit())
	for rows.Next() {
		var tran models.Transaction
		if err := rows.Scan(
			&tran.ID,
			&tran.TransactionID,
			&tran.Source,
			&tran.Description,
			&tran.SenderID,
			&tran.RecipientID,
			&tran.Currency,
			&tran.Amount,
			&tran.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("GetTransactions.RowsScan: %w", err)
		}
		transactions = append(transactions, &tran)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetTransactions.RowsErr: %w", err)
	}
	return &models.TransactionList{
		TotalCount:   total,
		TotalPages:   transaction.GetTotalPages(total),
		Page:         transaction.GetPage(),
		Size:         transaction.GetSize(),
		HasMore:      transaction.GetHasMore(total),
		Transactions: transactions,
	}, nil
}
