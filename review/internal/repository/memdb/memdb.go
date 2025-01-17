package memdb

import (
	"fmt"

	"github.com/ctfrancia/go-code-review/review/internal/service/entity"
)

type repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
	Close() error
}

// Repository is the struct that holds the repository
type Repository struct {
	entries map[string]entity.Coupon
}

// New creates a new repository instance
func New() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon),
	}
}

// FindByCode finds a coupon by code
func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("Coupon not found")
	}
	return &coupon, nil
}

// Save saves a coupon
func (r *Repository) Save(coupon entity.Coupon) error {
	r.entries[coupon.Code] = coupon
	return nil
}

// Close closes the database connection
func (r *Repository) Close() error {
	// immitate graceful shutdown
	return nil
}
