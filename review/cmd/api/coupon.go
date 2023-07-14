package api

import (
	"net/http"

	"github.com/ctfrancia/go-code-review/review/cmd/api/dto"

	"github.com/gin-gonic/gin"
)

// Apply is the endpoint for applying a coupon to a basket
func (a *API) Apply(c *gin.Context) {
	apiReq := dto.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, basket)
}

// Create is the endpoint for creating a coupon
func (a *API) Create(c *gin.Context) {
	apiReq := dto.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	_, err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		return
	}
	c.Status(http.StatusOK)
}

// Get is the endpoint for getting a coupon
func (a *API) Get(c *gin.Context) {
	apiReq := dto.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return
	}
	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, coupons)
}
