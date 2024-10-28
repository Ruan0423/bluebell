package router

import (
	"web_framework/controlle"
	"web_framework/logger"
	"web_framework/middleware"

	"github.com/gin-gonic/gin"
)

func SetUprouter() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(),logger.GinRecovery(true))

	v1 := r.Group("api/v1")

	v1.POST("/register", controlle.RegisterHandler)
	v1.POST("/login", controlle.Login)
	v1.GET("login", func(ctx *gin.Context) {
		ctx.String(200,"please Login")
	})

	v1.Use(middleware.JWTAuthMiddleware())
	

	{
		v1.GET("/community", controlle.CommunityHandler)
		v1.GET("/community/:id", controlle.CommunityDetailHandler)
		v1.POST("/post", controlle.CreatePostHandler)
		v1.GET("/post/:id", controlle.GetPostHandle)
	}

	return r
}