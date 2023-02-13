package router

import (
	"github.com/gin-gonic/gin"
	"gmf/src/common/router"
	"gmf/src/servers/user/handler"
)

type Router struct {
	router.AbstractRouter
}

func (r *Router) InAuthentic(group *gin.RouterGroup) {
	router := group.Group("user")
	{
		router.POST("/register", handler.UserRegister)
		router.POST("/login", handler.UserLogin)
	}
}
func (r *Router) Authentic(group *gin.RouterGroup) {
}
