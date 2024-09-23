package router

import (
	"web_framework/logger"

	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	
	r.GET("/",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"msg":"test OK!!!!",
		})
	})
	return r
}