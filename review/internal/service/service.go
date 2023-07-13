package service

import (
	"fmt"

	"github.com/ctfrancia/go-code-review/review/internal/service/entity"
	"github.com/google/uuid"
)

// Repository is an interface that defines the methods that a repository must implement
type Repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}

// Service is the struct that holds the repository
type Service struct {
	repo Repository
}

// New creates a new service instance
func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

// ApplyCoupon applies a coupon to a basket
func (s Service) ApplyCoupon(basket entity.Basket, code string) (b *entity.Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
	}
	if b.Value == 0 {
		return
	}

	if b.Value < 0 {
		return b, fmt.Errorf("Tried to apply discount to negative value")
	}

	return b, nil
}

// CreateCoupon creates a new coupon
func (s Service) CreateCoupon(discount int, code string, minBasketValue int) (entity.Coupon, error) {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	if err := s.repo.Save(coupon); err != nil {
		return entity.Coupon{}, err
	}
	return coupon, nil
}

// GetCoupons gets coupons by codes
func (s Service) GetCoupons(codes []string) ([]entity.Coupon, error) {
	coupons := make([]entity.Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, e
}
