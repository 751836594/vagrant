package routers

import (
	"github.com/gin-gonic/gin"
	"ucx/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/s", api.GetStripePlan)
	r.GET("/si", api.CanSubscribe)
	return r
}
