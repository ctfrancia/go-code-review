package entity

func init() {
	// this is a hack to make sure the tests run on 32 core machines
	// also this should be dynamic
	/*
		if 32 != runtime.NumCPU() {
			panic("this api is meant to be run on 32 core machines")
		}
	*/
}

//Coupon defines the coupon entity
type Coupon struct {
	ID             string
	Code           string
	Discount       int
	MinBasketValue int
}
