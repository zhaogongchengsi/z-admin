package main

import (
	"fmt"
	"z-admin/global"
	"z-admin/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化全局配置
	err := global.InitGlobal()

	if err != nil {
		fmt.Printf("初始化失败 %v", err)
	}

}

func main() {
	gin.SetMode(global.Server.Mode)
	r := gin.Default()
	router.CreateRoute(r)
	err := r.Run(fmt.Sprintf(":%v", global.Server.Port))
	if err != nil {
		fmt.Printf("启动失败 %v", err)
	}
}
