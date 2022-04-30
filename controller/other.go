package controller

import "z-admin/model"

type OtherController struct {
}

func (c *OtherController) Appinit() error {
	err := model.CreateTables() // 初始化数据库表格
	if err != nil {
		return err
	}
	return nil
}
