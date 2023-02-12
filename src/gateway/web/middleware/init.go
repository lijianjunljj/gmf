package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

// 接受服务实例，并存到gin.Key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	fmt.Println("service:", service)
	return func(context *gin.Context) {
		// 将实例存在gin.Keys中
		context.Keys = make(map[string]interface{})
		for _, s := range service {
			fmt.Println("s:", s)
			v := reflect.ValueOf(s)
			fmt.Println("v:", v)
			Name := v.Elem().FieldByName("name").String()
			fmt.Println("Name:", Name)
			context.Keys[Name] = s
		}
		context.Next()
	}
}

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				context.Abort()
			}
		}()
		context.Next()
	}
}
