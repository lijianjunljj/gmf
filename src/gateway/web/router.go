package web

import (
	"github.com/gin-gonic/gin"
	"gmf/src/gateway/web/handlers"
)

func UnAuthedRouter(v1 *gin.RouterGroup) {
	// 用户服务
	v1.POST("/user/register", handlers.UserRegister)
	v1.POST("/user/login", handlers.UserLogin)
}

func AuthedRouter(authed *gin.RouterGroup) {
	authed.GET("tasks", handlers.GetTaskList)
	authed.POST("task", handlers.CreateTask)
	authed.GET("task/:id", handlers.GetTaskDetail) // task_id
	authed.PUT("task/:id", handlers.UpdateTask)    // task_id
	authed.DELETE("task/:id", handlers.DeleteTask) // task_id
}
