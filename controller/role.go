package controller

import (
	"z-admin/model"
)

type RoleController struct {
	RoleName string       `json:"role_name"`
	RoleId   int          `json:"role_id"`
	RolePId  int          `json:"role_pid"`
	MenuTree []model.Menu `json:"menu_tree"`
}

func (a *RoleController) CreateRule() (*model.Role, error) {
	r := model.NewRole(a.RoleName, a.RoleId, a.RolePId)
	r.MenuTree = a.MenuTree
	return r.CreateRole()
}

func (a *RoleController) FindRoleList() ([]model.Role, error) {
	return model.NewRole(a.RoleName, a.RoleId, a.RolePId).FindRoleList()
}

func (a *RoleController) FindRoleById() (model.Role, error) {
	r, err := model.NewRole(a.RoleName, a.RoleId, a.RolePId).FindRoleById()
	if err != nil {
		return model.Role{}, err
	}
	return *r, nil
}

func (a *RoleController) SetPermis() error {
	r := model.Role{
		RoleId:   a.RoleId,
		MenuTree: a.MenuTree,
	}
	return r.SetPermis()
}
