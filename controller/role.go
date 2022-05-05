package controller

type RoleController struct {
	RoleName string `json:"role_name"`
	RoleId   int    `json:"role_id"`
	RolePId  int    `json:"role_pid"`
}

func (a *RoleController) Rule() error {
	return nil
}
