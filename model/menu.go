package model

import (
	"z-admin/global"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name       string `json:"name"`
	Path       string `json:"path"`
	Meta       string `json:"meta"`
	Components string `json:"component"`
	Icon       string `json:"icon"`
	Hidden     bool   `json:"hidden"`
	ParentId   int    `json:"parent_id"`
	Redirect   string `json:"redirect"`
	Label      string `json:"label"`
}

var initMenuList = []Menu{
	{
		Name:       "system",
		Path:       "system",
		Meta:       "",
		Components: "views/system/index.vue",
		Icon:       "Apple",
		Hidden:     false,
		ParentId:   0,
		Redirect:   "",
		Label:      "系统管理",
	},
	{
		Name:       "menu",
		Path:       "menu",
		Meta:       "",
		Components: "views/system/menu/index.vue",
		Icon:       "Apple",
		Hidden:     false,
		ParentId:   1,
		Redirect:   "",
		Label:      "菜单管理",
	},
}

// 创建用户表
func CreateMenuTable() error {
	if err := global.DBEngine.AutoMigrate(&Menu{}); err != nil {
		return err
	}

	return global.DBEngine.Create(initMenuList).Error
}

// 创建一个用户
func CreateMenuModel(name, path, meta, components, icon string, hidden bool, parentId int, redirect, Label string) *Menu {
	return &Menu{
		Name:       name,
		Path:       path,
		Meta:       meta,
		Components: components,
		Icon:       icon,
		Hidden:     hidden,
		ParentId:   parentId,
		Redirect:   redirect,
		Label:      Label,
	}
}

func (c *Menu) CreateMenu() (menu Menu, err error) {
	err = global.DBEngine.Create(c).Error
	return *c, err
}

type MenuRes struct {
	ID         int    `json:"ID"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Meta       string `json:"meta"`
	Components string `json:"component"`
	Icon       string `json:"icon"`
	Hidden     bool   `json:"hidden"`
	ParentId   int    `json:"parent_id"`
	Redirect   string `json:"redirect"`
	Label      string `json:"label"`
}

func (c *Menu) GetMenuList() (menu []MenuRes, err error) {
	err = global.DBEngine.Model(&Menu{}).Find(&menu).Error
	return menu, err
}

func (c *Menu) DeleteMenu() (err error) {
	err = global.DBEngine.Delete(c).Error
	return err
}

func (c *Menu) FindMenuByID(id string) (menu Menu, err error) {
	err = global.DBEngine.Where("id = ?", id).First(&menu).Error
	return menu, err
}
