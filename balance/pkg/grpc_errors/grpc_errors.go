package grpc_errors

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrNoCtxMetaData   = errors.New("no ctx metadata")
	ErrConvertCurrency = errors.New("conversion error")
	ErrUserExists      = errors.New("user already exists")
	ErrDB              = errors.New("DB error")
)

// ParseGRPCErrStatusCode Parse error and get code
func ParseGRPCErrStatusCode(err error) codes.Code {
	switch {
	case errors.Is(err, ErrNotFound):
		return codes.NotFound
	case errors.Is(err, context.Canceled):
		return codes.Canceled
	case errors.Is(err, context.DeadlineExceeded):
		return codes.DeadlineExceeded
	case errors.Is(err, ErrUserExists):
		return codes.AlreadyExists
	case errors.Is(err, ErrNoCtxMetaData):
		return codes.Unauthenticated
	case errors.Is(err, ErrConvertCurrency):
		return codes.Unavailable
	case strings.Contains(err.Error(), "Validate"):
		return codes.InvalidArgument
	case strings.Contains(err.Error(), "redis"):
		return codes.NotFound
	case errors.Is(err, ErrDB):
		return codes.Unavailable
	}
	return codes.Internal
}

// MapGRPCErrCodeToHttpStatus Map GRPC errors codes to http status
func MapGRPCErrCodeToHttpStatus(code codes.Code) int {
	switch code {
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.AlreadyExists:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.InvalidArgument:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
