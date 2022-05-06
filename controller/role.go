package controller

import (
	"z-admin/model"
)

type RoleController struct {
	RoleName string           `json:"role_name"`
	RoleId   int              `json:"role_id"`
	RolePId  int              `json:"role_pid"`
	MenuTree []MenuController `json:"menu_tree"`
}

func (a *RoleController) CreateRule() (*model.Role, error) {
	return model.NewRole(a.RoleName, a.RoleId, a.RolePId).CreateRole()
}

func (a *RoleController) FindRoleList() ([]model.Role, error) {
	return model.NewRole(a.RoleName, a.RoleId, a.RolePId).FindRoleList()
}
