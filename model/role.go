package model

import (
	"errors"
	"z-admin/global"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `json:"role_name"`
	RoleId   int    `json:"role_id"`
	RolePid  int    `json:"role_pid"`
	MenuTree []Menu `json:"menu_tree" gorm:"many2many:role_menu"` // 菜单树
}

// 创建角色表
func CreateRoleTable() error {
	return global.DBEngine.AutoMigrate(&Role{})
}

// 创建角色
func NewRole(roleName string, roleId, rolePid int) *Role {
	return &Role{
		RoleName: roleName,
		RoleId:   roleId,
		RolePid:  rolePid,
	}
}

func (r *Role) CreateRole() (role *Role, err error) {

	err = global.DBEngine.Create(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil

}

func (r *Role) FindRoleById() (rs *Role, err error) {
	var role Role
	err = global.DBEngine.Model(role).First(&rs, "role_id = ?", r.RoleId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("角色不存在")
		}
		return nil, err
	}
	return rs, nil
}

func (r *Role) FindRoleList() (rs []Role, err error) {
	err = global.DBEngine.Find(&rs).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("无数据")
	}
	return rs, err
}

func (r *Role) FindRoleListByPid(pid int) (rs []Role, err error) {
	err = global.DBEngine.Where("role_pid = ?", pid).Find(&rs).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return nil, errors.New("无数据")
	}
	return rs, err
}
