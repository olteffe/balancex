package service

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/olteffe/balancex/balance/internal/transaction"
	"github.com/olteffe/balancex/balance/internal/transaction/models"
	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/olteffe/balancex/balance/pkg/logger"
)

// transactionService transaction service
type transactionService struct {
	tranRepo  transaction.PGRepository
	redisRepo transaction.RedisRepository
	logger    logger.Logger
}

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
		return "", err
	}
	if exist != true {
		return "", grpc_errors.ErrUserExists
	}

	convert, err := t.redisRepo.ConvertTransaction(ctx, transaction)
	if err != nil {
		return "", grpc_errors.ErrConvertCurrency
	}
	return t.tranRepo.CreateTransaction(ctx, convert)
}

// GetTransactions get user transactions
func (t *transactionService) GetTransactions(ctx context.Context, transaction *models.TransactionsRequest) (*models.TransactionList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactionService.GetTransactions")
	defer span.Finish()

	exist, err := t.tranRepo.FindUserID(ctx, transaction)
	if err != nil {
		return nil, err
	}
	if exist != true {
		return nil, grpc_errors.ErrUserExists
	}

	convert, err := t.redisRepo.ConvertTransactions(ctx, transaction)
	if err != nil {
		return nil, grpc_errors.ErrConvertCurrency
	}
	return t.tranRepo.GetTransactions(ctx, convert)
}
