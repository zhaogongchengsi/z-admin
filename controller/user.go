package controller

import (
	"errors"
	"z-admin/global"
	"z-admin/model"

	"gorm.io/gorm"
)

type User struct {
	UserName  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	NickName  string `json:"nickname"`
	AvatarUrl string `json:"avatar"`
	Verify    string `json:"verify"`
	VerifyId  string `json:"verify_Id"`
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

func (u *User) Login() (*model.UserModel, error) {

	if len(u.Verify) == 0 && len(u.VerifyId) == 0 {
		return nil, errors.New("验证码不能为空")
	}

	if !global.VerifyStore.Verify(u.VerifyId, u.Verify, true) {
		return nil, errors.New("验证码错误")
	}

	user, err := model.CreateUserModel(u.UserName, u.Password, "", "")

	if err != nil {
		return nil, err
	}

	userp, err := user.FindUserByPasswordAndUsername()

	if err != nil {
		return nil, errors.New("密码错误或用户不存在")
	}

	userp.PassWord = ""

	return userp, nil

}

func (u *User) ChangePass(newpassword string) error {

	user, err := model.CreateUserModel(u.UserName, u.Password, "", "")

	if err != nil {
		return err
	}

	oldUser, usererr := user.FindUserByPasswordAndUsername()

	if usererr != nil {
		return errors.New("用户不存在或者密码错误(请检查用户名和密码)")
	}

	newUser, err := model.CreateUserModel(oldUser.UserName, newpassword, oldUser.NickName, oldUser.AvatarUrl)

	if err != nil {
		return err
	}

	err = oldUser.UpdateUser(newUser)

	if err != nil {
		return err
	}

	return nil
}
