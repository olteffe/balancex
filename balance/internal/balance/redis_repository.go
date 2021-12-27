//go:generate mockgen -source redis_repository.go -destination mock/redis_repository.go -package mock
package balance

import (
	"context"
)

// RedisRepository Redis balance repository interface
type RedisRepository interface {
	GetRate(ctx context.Context, rate string) (float32, error)
}
