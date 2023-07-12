package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ctfrancia/go-code-review/review/internal/service/entity"

	"github.com/gin-gonic/gin"
)

// Service is the interface that provides coupon methods.
type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) interface{} // FIXME: do correct type should return error
	GetCoupons([]string) ([]entity.Coupon, error)
}

// Config is the configuration for api.
type Config struct {
	Host string
	Port int
}

// API provides coupon api.
type API struct {
	srv *http.Server
	MUX *gin.Engine
	svc Service
	CFG Config
}

// New creates new api instance.
func New(cfg Config, svc Service) API {
	// func (s Service) New(cfg Config, svc Service) API {
	gin.SetMode(gin.ReleaseMode)
	r := new(gin.Engine)
	r = gin.New()
	r.Use(gin.Recovery())

	return API{
		MUX: r,
		CFG: cfg,
		svc: svc,
	}.withServer()
}

func (a API) withServer() API {
	ch := make(chan API)
	go func() {
		a.srv = &http.Server{
			Addr:    fmt.Sprintf(":%d", a.CFG.Port),
			Handler: a.MUX,
		}
		ch <- a
	}()

	return <-ch
}

func (a API) withRoutes() API {
	apiGroup := a.MUX.Group("/api")
	apiGroup.POST("/apply", a.Apply)
	apiGroup.POST("/create", a.Create)
	apiGroup.GET("/coupons", a.Get)
	return a
}

// Start runs http server.
func (a API) Start() {
	if err := a.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// Close closes http server.
func (a API) Close() {
	<-time.After(5 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
