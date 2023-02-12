package web

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gmf/src/gateway/web/middleware"
)

func NewRouter(service []interface{}) *gin.Engine {
	fmt.Println("service2:", service)
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	root := ginRouter.Group("/")
	{
		root.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})

		UnAuthedRouter(root)
		// 需要登录保护
		authed := root.Group("/")
		authed.Use(middleware.JWT())
		{
			AuthedRouter(authed)
		}
	}
	return ginRouter
}
