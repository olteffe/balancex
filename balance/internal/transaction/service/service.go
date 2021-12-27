package service

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"github.com/olteffe/balancex/balance/internal/transaction"
	"github.com/olteffe/balancex/balance/internal/transaction/models"
	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/olteffe/balancex/balance/pkg/logger"
	"github.com/olteffe/balancex/balance/pkg/utils"
)

// transactionService transaction service
type transactionService struct {
	tranRepo  transaction.PGRepository
	redisRepo transaction.RedisRepository
	logger    logger.Logger
}

const (
	mainCurrency = "RUB"
)

// NewTransactionService Transaction service constructor
func NewTransactionService(tranRepo transaction.PGRepository, redisRepo transaction.RedisRepository,
	logger logger.Logger) *transactionService {
	return &transactionService{tranRepo: tranRepo, redisRepo: redisRepo, logger: logger}
}

// CreateTransaction Create Transaction
func (t *transactionService) CreateTransaction(ctx context.Context, transaction *models.Transaction) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactionService.CreateTransaction")
	defer span.Finish()

	// find senderID & recipientID in repository
	exist, err := t.tranRepo.FindUsersID(ctx, transaction)
	if err != nil {
		return "", fmt.Errorf("CreateTransaction.FindUsersID: %w", err)
	}
	if !exist {
		return "", fmt.Errorf("CreateTransaction.FindUsersID: %w", grpc_errors.ErrUserExists)
	}

	convert, err := t.convertRate(ctx, transaction)
	if err != nil {
		return "", grpc_errors.ErrConvertCurrency
	}
	return t.tranRepo.CreateTransaction(ctx, convert)
}

// GetTransactions get user transactions
func (t *transactionService) GetTransactions(ctx context.Context, transaction *utils.TransactionsRequest) (*models.TransactionList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactionService.GetTransactions")
	defer span.Finish()

	exist, err := t.tranRepo.FindUserID(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("GetTransactions.FindUserID: %w", err)
	}
	if !exist {
		return nil, fmt.Errorf("GetTransactions.FindUsersID: %w", grpc_errors.ErrUserExists)
	}
	return t.tranRepo.GetTransactions(ctx, transaction)
}

// convertRate converts currency
func (t *transactionService) convertRate(ctx context.Context, transaction *models.Transaction) (*models.Transaction, error) {
	if transaction.Currency == mainCurrency {
		return transaction, nil
	}
	rate, err := t.redisRepo.GetRate(ctx, transaction.Currency)
	if err != nil {
		return nil, fmt.Errorf("convertRate: %w", err)
	}
	transaction.Amount = int64(rate * float32(transaction.Amount)) // problems with law
	return transaction, nil
}
