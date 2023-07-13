package main

import (
	"fmt"
	"time"

	"github.com/ctfrancia/go-code-review/review/cmd/api"
	"github.com/ctfrancia/go-code-review/review/internal/config"
	"github.com/ctfrancia/go-code-review/review/internal/repository/memdb"
	"github.com/ctfrancia/go-code-review/review/internal/service"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func init() {
	// this is a hack to make sure the tests run on 32 core machines
	// also this should be dynamic
	/*
		if 32 != runtime.NumCPU() {
			panic("this api is meant to be run on 32 core machines")
		}
	*/
}

func main() {
	svc := service.New(repo)
	an := api.New(cfg.API, svc)
	an.Start()
	fmt.Println("Starting Coupon service server")
	<-time.After(1 * time.Hour * 24 * 365)
	fmt.Println("Coupon service server alive for a year, closing")
	an.Close()
}
