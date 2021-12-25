package utils

import (
	"fmt"
	"math"
	"strconv"

	"github.com/google/uuid"
)

const (
	defaultSize = 10
)

// TransactionsRequest Pagination query params
type TransactionsRequest struct {
	UserID   uuid.UUID `db:"user_id" validate:"required,uuid"`
	Currency string    `db:"currency" validate:"required,len=3,uppercase"`
	Size     uint32    `db:"size" validate:"omitempty,gt=0"`
	Page     uint32    `db:"page" validate:"omitempty,gt=0"`
	OrderBy  string    `validate:"omitempty,len=3"`
}

// NewPaginationQuery Pagination query constructor
func NewPaginationQuery(userID uuid.UUID, currency string, size uint32, page uint32, orderBy string) *TransactionsRequest {
	return &TransactionsRequest{
		UserID:   userID,
		Currency: currency,
		Size:     size,
		Page:     page,
		OrderBy:  orderBy,
	}
}

// SetStringSize Set page size
func (t *TransactionsRequest) SetStringSize(sizeQuery string) error {
	if sizeQuery == "" {
		t.Size = defaultSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	t.Size = uint32(n)

	return nil
}

// SetStringPage Set page number
func (t *TransactionsRequest) SetStringPage(pageQuery string) error {
	if pageQuery == "" {
		t.Size = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	t.Page = uint32(n)

	return nil
}

// SetOrderBy Set order by
func (t *TransactionsRequest) SetOrderBy(orderByQuery string) {
	t.OrderBy = orderByQuery
}

// GetOffset Get offset
func (t *TransactionsRequest) GetOffset() uint32 {
	if t.Page == 0 {
		return 0
	}
	return (t.Page - 1) * t.Size
}

// GetLimit Get limit
func (t *TransactionsRequest) GetLimit() uint32 {
	return t.Size
}

// GetOrderBy Get OrderBy
func (t *TransactionsRequest) GetOrderBy() string {
	return t.OrderBy
}

// GetPage Get page
func (t *TransactionsRequest) GetPage() uint32 {
	return t.Page
}

// GetSize Get size
func (t *TransactionsRequest) GetSize() uint32 {
	return t.Size
}

// GetQueryString Get query string
func (t *TransactionsRequest) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%s", t.GetPage(), t.GetSize(), t.GetOrderBy())
}

// GetTotalPages Get total pages int
func (t *TransactionsRequest) GetTotalPages(totalCount uint32) uint32 {
	d := float64(totalCount) / float64(t.GetSize())
	return uint32(math.Ceil(d))
}

// GetHasMore Get has more
func (t *TransactionsRequest) GetHasMore(totalCount uint32) bool {
	return t.GetPage() < totalCount/t.GetSize()
}
