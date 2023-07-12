package service

import (
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
		// {"Apply 10% to 0", fields{memdb.New()}, args{entity.Basket{Value: 0}, "Superdiscount"}, &entity.Basket{Value: 0, AppliedDiscount: 0, ApplicationSuccessful: false}, false},
		// {"Apply 10% to -1", fields{memdb.New()}, args{entity.Basket{Value: -1}, "Superdiscount"}, &entity.Basket{Value: -1, AppliedDiscount: 0, ApplicationSuccessful: false}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
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

			s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
		})
	}
}
