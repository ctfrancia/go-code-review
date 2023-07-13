package entity

//Coupon defines the coupon entity
type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}
