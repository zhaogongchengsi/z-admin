package controller

import "z-admin/model"

type MenuController struct {
	Name       string `json:"name"`
	Path       string `json:"path"`
	Meta       string `json:"meta"`
	Components string `json:"components"`
	Icon       string `json:"icon"`
	Hidden     bool   `json:"hidden"`
	ParentId   int    `json:"parent_id"`
	Redirect   string `json:"redirect"`
}

func (m *MenuController) GetMenu() (menu []model.Menu, err error) {
	me := new(model.Menu)
	return me.GetMenuList()
}

func (m *MenuController) CreateMenu() (menu model.Menu, err error) {
	return model.CreateMenuModel(m.Name, m.Path, m.Meta, m.Components, m.Icon, m.Hidden, m.ParentId, m.Redirect).CreateMenu()
}
