package controller

import (
	"errors"
	"z-admin/model"

	"gorm.io/gorm"
)

type User struct {
	UserName  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	NickName  string `json:"nickname"`
	AvatarUrl string `json:"avatar"`
}

// 注册用户
func (u *User) Register() (user *model.UserModel, err error) {

	oldUser, err := model.FindUserByName(u.UserName)
	if err != gorm.ErrRecordNotFound {
		return oldUser, errors.New("用户已存在")
	}

	newUser, err := model.NewUserModel(u.UserName, u.Password, u.NickName, u.AvatarUrl)
	if err != nil {
		return nil, err
	}

	newUser, err = newUser.CreateUser()

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
