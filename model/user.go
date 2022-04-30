package model

import (
	"z-admin/global"
	"z-admin/utils"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UId       uuid.UUID `gorm:"type:varchar(36);unique_index" json:"uid"`
	UserName  string    `grom:"type:varchar(100)" json:"username"`
	PassWord  string    `grom:"type:varchar(100);not null" json:"password"`
	NickName  string    `grom:"type:varchar(100)" json:"nickname"`
	AvatarUrl string    `grom:"type:varchar(100)" json:"avatar_url"`
}

// 创建用户表
func CreateUserTable() error {
	return global.DBEngine.AutoMigrate(&UserModel{})
}

func NewUserModel(u, p, n, a string) (*UserModel, error) {
	id := uuid.NewV4()        // 用户唯一标识
	newp, err := utils.Md5(p) // 加密密码 不可逆
	if err != nil {
		return nil, err
	}
	return &UserModel{
		UId:       id,
		UserName:  u,
		PassWord:  string(newp),
		NickName:  n,
		AvatarUrl: a,
	}, nil
}

// 创建用户
func (m *UserModel) CreateUser() (*UserModel, error) {
	err := global.DBEngine.Create(m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// 根据用户账号查找用户
func FindUserByName(name string) (*UserModel, error) {
	var user UserModel
	err := global.DBEngine.Where("username = ?", name).First(&user).Error
	return &user, err
}

// 更新用户信息
func (m *UserModel) UpdateUser() error {
	return global.DBEngine.Model(m).Updates(m).Error
}

func (m *UserModel) GetUserList() ([]UserModel, error) {
	var users []UserModel
	err := global.DBEngine.Find(&users).Error
	return users, err
}

func (m *UserModel) DeletelUser() error {
	return global.DBEngine.Delete(m).Error
}
