package controller

import (
	"z-admin/global"
	"z-admin/model"

	"github.com/mojocn/base64Captcha"
)

type OtherController struct {
}

func (c *OtherController) Appinit() error {
	err := model.CreateTables() // 初始化数据库表格
	if err != nil {
		return err
	}
	return nil
}

// 生成验证码 并返回验证码ID 和 图片 base64
func (c *OtherController) Verify() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(
		global.CaptchaConf.ImgHeight,
		global.CaptchaConf.ImgWidth,
		global.CaptchaConf.KeyLong,
		global.CaptchaConf.MaxSkew,
		global.CaptchaConf.DotCount,
	)
	cp := base64Captcha.NewCaptcha(driver, global.VerifyStore)
	return cp.Generate()
}
