package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/olteffe/balancex/balance/config"
	"github.com/olteffe/balancex/balance/internal/balance"
	"github.com/olteffe/balancex/balance/internal/balance/models"
	protobalance "github.com/olteffe/balancex/balance/internal/balance/proto"
	"github.com/olteffe/balancex/balance/pkg/grpc_errors"
	"github.com/olteffe/balancex/balance/pkg/logger"
	"github.com/olteffe/balancex/balance/pkg/utils"
)

// balanceService gRPC microservice
type balanceService struct {
	cfg        *config.Config
	logger     logger.Logger
	balService balance.Service
}

// NewBalanceService gRPC microservice constructor
func NewBalanceService(balService balance.Service, logger logger.Logger, cfg *config.Config) *balanceService {
	return &balanceService{balService: balService, logger: logger, cfg: cfg}
}

// CreateBalance Create new user balance
func (b *balanceService) CreateBalance(ctx context.Context, r *protobalance.CreateBalanceRequest) (*protobalance.CreateBalanceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "balance.Create")
	defer span.Finish()

	newBalance, err := b.createBalanceReqToBalanceModel(r)
	if err != nil {
		b.logger.Errorf("registerReqToBalanceModel: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "registerReqToBalanceModel: %v", err)
	}

	if err := utils.ValidateStruct(ctx, newBalance); err != nil {
		b.logger.Errorf("ValidateStruct: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "ValidateStruct: %v", err)
	}

	createdBalance, err := b.balService.CreateBalance(ctx, newBalance)
	if err != nil {
		b.logger.Errorf("balService.CreateBalance: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "Create: %v", err)
	}

	return &protobalance.CreateBalanceResponse{UserId: createdBalance}, nil
}

// createBalanceReqToBalanceModel mapping in model
func (b *balanceService) createBalanceReqToBalanceModel(r *protobalance.CreateBalanceRequest) (*models.Balance, error) {
	userID, err := uuid.Parse(r.Balance.GetUserId())
	if err != nil {
		b.logger.Errorf("uuidParse: %v", err)
		return nil, err
	}
	candidate := &models.Balance{
		UserID:   userID,
		Currency: r.Balance.GetCurrency(),
		Amount:   r.Balance.GetAmount(),
	}
	return candidate, nil
}

// GetBalance get user balance
func (b *balanceService) GetBalance(ctx context.Context, r *protobalance.GetBalanceRequest) (*protobalance.GetBalanceResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "balance.Get")
	defer span.Finish()

	model, err := b.getBalanceReqToBalanceModel(r)
	if err != nil {
		b.logger.Errorf("getBalanceReqToBalanceModel: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "getBalanceReqToBalanceModel: %v", err)
	}

	getBalance, err := b.balService.GetBalance(ctx, model)
	if err != nil {
		b.logger.Errorf("balService.GetBalance: %v", err)
		return nil, status.Errorf(grpc_errors.ParseGRPCErrStatusCode(err), "balService.GetBalance: %v", err)
	}

	return &protobalance.GetBalanceResponse{Balance: b.balanceModelToProto(getBalance)}, nil
}

// getBalanceReqToBalanceModel mapping in model
func (b *balanceService) getBalanceReqToBalanceModel(r *protobalance.GetBalanceRequest) (*models.Balance, error) {
	userID, err := uuid.Parse(r.GetUserId())
	if err != nil {
		b.logger.Errorf("uuidParse: %v", err)
		return nil, err
	}
	candidate := &models.Balance{
		UserID:   userID,
		Currency: r.GetCurrency(),
	}
	return candidate, nil
}

func (b *balanceService) balanceModelToProto(user *models.Balance) *protobalance.Balance {
	balanceProto := &protobalance.Balance{
		UserId:    user.UserID.String(),
		Currency:  user.Currency,
		Amount:    user.Amount,
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	return balanceProto
}
