package router

import (
	"web_framework/controlle"
	"web_framework/logger"
	"web_framework/middleware"
	"web_framework/settings"

	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))
	
	r.GET("/",middleware.JWTAuthMiddleware(),func(ctx *gin.Context) {

		userid,err := controlle.GetUserId(ctx)
		ctx.JSON(200,gin.H{
			"msg":settings.Conf.APP.Port,
			"username":userid,
			"err": err,
		})
	})
	r.POST("/register", controlle.RegisterHandler)
	r.POST("/login", controlle.Login)
	return r
}