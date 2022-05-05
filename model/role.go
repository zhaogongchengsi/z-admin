package model

import "z-admin/global"

type Role struct {
	RoleName string  `json:"role_name"`
	RoleId   int     `json:"role_id"`
	RolePid  int     `json:"role_pid"`
	MenuTree []*Menu `json:"menu_tree" gorm:"many2many:role_menu"` // 菜单树
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

func (r *Role) FindRoleList() ([]Role, error) {

	var rs []Role
	err := global.DBEngine.Where("role_id = ?", r.RoleId).Find(rs).Error

	if err != nil {
		return nil, err
	}

	return rs, nil
}
