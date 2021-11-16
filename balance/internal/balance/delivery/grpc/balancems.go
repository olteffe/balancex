package grpc

import (
	"github.com/olteffe/balancex/balance/config"
	"github.com/olteffe/balancex/balance/internal/balance"
	"github.com/olteffe/balancex/balance/pkg/logger"
)

// BalanceMicroservice gRPC microservice
type BalanceMicroservice struct {
	cfg        *config.Config
	logger     logger.Logger
	balService balance.BalanceService
}

// NewBalanceMicroservice gRPC microservice constructor
func NewBalanceMicroservice(balService balance.BalanceService, logger logger.Logger, cfg *config.Config) *BalanceMicroservice {
	return &BalanceMicroservice{balService: balService, logger: logger, cfg: cfg}
}
