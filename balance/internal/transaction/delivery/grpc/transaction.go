package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/olteffe/balancex/balance/config"
	"github.com/olteffe/balancex/balance/internal/transaction"
	"github.com/olteffe/balancex/balance/internal/transaction/models"
	pt "github.com/olteffe/balancex/balance/internal/transaction/proto"
	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/olteffe/balancex/balance/pkg/logger"
	"github.com/olteffe/balancex/balance/pkg/utils"
)

// transactionService gRPC service
type transactionService struct {
	cfg         *config.Config
	logger      logger.Logger
	tranService transaction.Service
}

// NewTransactionService gRPC service constructor
func NewTransactionService(tranService transaction.Service, logger logger.Logger, cfg *config.Config) *transactionService {
	return &transactionService{tranService: tranService, logger: logger, cfg: cfg}
}

// CreateTransaction Create new transaction
func (t *transactionService) CreateTransaction(ctx context.Context, r *pt.TransactionRequest) (*pt.TransactionResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transaction.Create")
	defer span.Finish()

	newTransaction, err := t.createTransactionReqToModel(r)
	if err != nil {
		t.logger.Errorf("transactionReqToModel: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "transactionReqToModel: %v", err)
	}

	if err := utils.ValidateStruct(ctx, newTransaction); err != nil {
		t.logger.Errorf("CreateTransaction.ValidateStruct: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateTransaction.ValidateStruct: %v", err)
	}

	createdTransaction, err := t.tranService.CreateTransaction(ctx, newTransaction)
	if err != nil {
		t.logger.Errorf("tranService.CreateTransaction: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "CreateTransaction: %v", err)
	}
	return &pt.TransactionResponse{TransactionId: createdTransaction}, nil
}

// createTransactionReqToModel mapping in model
func (t *transactionService) createTransactionReqToModel(r *pt.TransactionRequest) (*models.Transaction, error) {
	transactionID, err := uuid.Parse(r.Transaction.GetTransactionId())
	if err != nil {
		t.logger.Errorf("transactionIDParse: %v", err)
		return nil, err
	}
	senderID, err := uuid.Parse(r.Transaction.GetSenderId())
	if err != nil {
		t.logger.Errorf("senderIDParse: %v", err)
		return nil, err
	}
	recipientID, err := uuid.Parse(r.Transaction.GetRecipientId())
	if err != nil {
		t.logger.Errorf("recipientIDParse: %v", err)
		return nil, err
	}
	candidate := &models.Transaction{
		TransactionID: transactionID,
		Source:        r.Transaction.GetSource(),
		Description:   r.Transaction.GetDescription(),
		SenderID:      senderID,
		RecipientID:   recipientID,
		Currency:      r.Transaction.GetCurrency(),
		Amount:        r.Transaction.GetAmount(),
		CreatedAt:     r.Transaction.GetCreatedAt().AsTime(),
	}
	return candidate, nil
}

// GetTransactions get history of transactions
func (t *transactionService) GetTransactions(ctx context.Context, r *pt.TransactionsRequest) (*pt.TransactionsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "transactions.Get")
	defer span.Finish()

	newTransactions, err := t.getTransactionsReqToModel(r)
	if err != nil {
		t.logger.Errorf("transactionsReqToModel: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "transactionsReqToModel: %v", err)
	}

	if err := utils.ValidateStruct(ctx, newTransactions); err != nil {
		t.logger.Errorf("GetTransactions.ValidateStruct: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetTransactions.ValidateStruct: %v", err)
	}

	createdTransactions, err := t.tranService.GetTransactions(ctx, newTransactions)
	if err != nil {
		t.logger.Errorf("transService.GetTransactions: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "GetTransactions: %v", err)
	}

	return t.transactionListToProto(createdTransactions), nil
}

// getTransactionsReqToModel mapping in model
func (t *transactionService) getTransactionsReqToModel(r *pt.TransactionsRequest) (*models.TransactionsRequest, error) {
	userID, err := uuid.Parse(r.GetUserId())
	if err != nil {
		t.logger.Errorf("transactionsIDParse: %v", err)
		return nil, err
	}
	candidate := &models.TransactionsRequest{
		UserID:   userID,
		Currency: r.GetCurrency(),
		Page:     r.GetPage(),
		Size:     r.GetSize(),
	}
	return candidate, nil
}

// transactionModelToProto mapping transaction to proto
func transactionModelToProto(transaction *models.Transaction) *pt.Transaction {
	transactionProto := &pt.Transaction{
		TransactionId: transaction.TransactionID.String(),
		Source:        transaction.Source,
		Description:   transaction.Description,
		SenderId:      transaction.SenderID.String(),
		RecipientId:   transaction.RecipientID.String(),
		Currency:      transaction.Currency,
		Amount:        transaction.Amount,
		CreatedAt:     timestamppb.New(transaction.CreatedAt),
	}
	return transactionProto
}

// transactionListToProto mapping transaction list to proto
func (t *transactionService) transactionListToProto(transaction *models.TransactionList) *pt.TransactionsResponse {
	list := make([]*pt.Transaction, 0, len(transaction.Transactions))
	for _, tran := range transaction.Transactions {
		list = append(list, transactionModelToProto(tran))
	}

	return &pt.TransactionsResponse{
		TotalCount:   transaction.TotalCount,
		TotalPages:   transaction.TotalPages,
		Page:         transaction.Page,
		Size:         transaction.Size,
		HasMore:      transaction.HasMore,
		Transactions: list,
	}
}
