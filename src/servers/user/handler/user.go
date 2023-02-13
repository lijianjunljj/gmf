package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gmf/src/gateway/common/logging"
	"gmf/src/gateway/common/utils"
	"gmf/src/servers/user/services"
	"net/http"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
}
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logging.Info(err)
		panic(err)
	}
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["UserService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	fmt.Println("UserLogin:", userResp, err)
	PanicIfUserError(err)
	fmt.Println("11111:")
	if userResp.UserDetail != nil {
		token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
		fmt.Println("token, err:", token, err)
		ginCtx.JSON(http.StatusOK, gin.H{
			"code": userResp.Code,
			"msg":  "成功",
			"data": gin.H{
				"user":  userResp.UserDetail,
				"token": token,
			},
		})
	} else {
		ginCtx.JSON(http.StatusOK, gin.H{
			"code": userResp.Code,
			"msg":  "失败",
		})
	}
}
