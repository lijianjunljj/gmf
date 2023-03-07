package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/lijianjunljj/gmfcommon/db"
	"gmf/src/servers/user/model"
	"gmf/src/servers/user/services"
	"gorm.io/gorm"
)

type UserService struct {
}

func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := db.GetDB().Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		fmt.Println("err l:", err)
		if err == gorm.ErrRecordNotFound {
			resp.Code = 400
			fmt.Println("resp.Code:", resp.Code)
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	var count int64 = 0
	if err := db.GetDB().Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := db.GetDB().Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}
