package entity

import "github.com/ctfrancia/go-code-review/review/internal/service/entity"

// ApplicationRequest is a struct that holds the request body for the Apply endpoint
type ApplicationRequest struct {
	Code   string
	Basket entity.Basket
}
