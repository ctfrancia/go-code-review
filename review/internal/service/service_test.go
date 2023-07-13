package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ctfrancia/go-code-review/review/internal/repository/memdb"
	"github.com/ctfrancia/go-code-review/review/internal/service/entity"
)

func TestNew(t *testing.T) {
	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"initialize service", args{repo: nil}, Service{repo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{"Apply 10%", fields{memdb.New()}, args{10, "Superdiscount", 55}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			err := s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			if err != nil {
				t.Errorf("CreateCoupon() error = %v", err)
				return
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {

	type fields struct {
		repo Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}{
		{"Apply 10%", fields{memdb.New()}, args{entity.Basket{Value: 100}, "Superdiscount"}, &entity.Basket{Value: 100, AppliedDiscount: 10, ApplicationSuccessful: true}, false},
		{"Apply 10% to 0", fields{memdb.New()}, args{entity.Basket{Value: 0}, "Superdiscount"}, &entity.Basket{Value: 0, AppliedDiscount: 0, ApplicationSuccessful: false}, false},
		{"Apply 10% to -1", fields{memdb.New()}, args{entity.Basket{Value: -1}, "Superdiscount"}, &entity.Basket{Value: -1, AppliedDiscount: 0, ApplicationSuccessful: false}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			err := s.CreateCoupon(10, "Superdiscount", 55)
			if err != nil {
				t.Errorf("CreateCoupon() error = %v", err)
				return
			}

			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_GetCoupons(t *testing.T) {
	type fields struct {
		repo Repository
	}
	tests := []struct {
		name    string
		fields  fields
		coupons []entity.Coupon
		want    []entity.Coupon
		wantErr bool
	}{
		{"Get Coupons", fields{memdb.New()}, []entity.Coupon{
			{ID: "123", Code: "Superdiscount", Discount: 10, MinBasketValue: 55},
		}, []entity.Coupon{{ID: "123", Code: "Superdiscount", Discount: 10, MinBasketValue: 55}}, false},
		{"Get Coupons multiple", fields{memdb.New()}, []entity.Coupon{
			{ID: "123", Code: "Superdiscount2", Discount: 10, MinBasketValue: 55},
			{ID: "456", Code: "Superdiscount3", Discount: 20, MinBasketValue: 85},
		}, []entity.Coupon{
			{ID: "123", Code: "Superdiscount2", Discount: 10, MinBasketValue: 55},
			{ID: "456", Code: "Superdiscount3", Discount: 20, MinBasketValue: 85},
		}, false},
	}
	xc := make([]string, 0)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			for _, coupon := range tt.coupons {
				fmt.Println("coupon", coupon)
				xc = append(xc, coupon.Code)
				err := s.CreateCoupon(coupon.Discount, coupon.Code, coupon.MinBasketValue)
				if err != nil {
					t.Errorf("CreateCoupon() error = %v", err)
					return
				}

			}

			fmt.Println("xc", xc)
			got, err := s.GetCoupons(xc)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoupons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoupons() got = %v, want %v", got, tt.want)
			}
		})
	}
}
