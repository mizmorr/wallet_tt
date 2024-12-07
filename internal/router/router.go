package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mizmorr/wallet/internal/controller"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(handler *gin.Engine, c *controller.WalletController) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := handler.Group("api/v1")
	{
		v1.POST("/wallet", c.Operate)
		v1.GET("/wallets/", c.Get)
	}
}
