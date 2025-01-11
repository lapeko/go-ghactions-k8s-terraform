package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lapeko/go-ghactions-k8s-terraform/auth/internal/logger"
)

type Api struct {
	router *gin.Engine
}

var api *Api
var ErrReinitialization = errors.New("API is already initialized")

func New() *Api {
	log := logger.GetLogger()

	if api == nil {
		log.Info("API instance creation.")
		api = &Api{}
	} else {
		log.Warn("Attempt to create an API duplicate instance. Returning created API instance...")
	}

	return api
}

func (a *Api) Init(port string) error {
	if a.router == nil {
		a.router = gin.Default()
	} else {
		return ErrReinitialization
	}

	a.router.POST("/register", func(context *gin.Context) {})
	a.router.POST("/login", func(ctx *gin.Context) {})
	a.router.POST("/refresh", func(context *gin.Context) {})

	return a.router.Run(fmt.Sprintf(":%s", port))
}
