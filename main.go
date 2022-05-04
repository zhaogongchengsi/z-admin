package main

import (
	"fmt"
	"net/http"
	"time"
	"z-admin/global"
	"z-admin/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化全局配置
	err := global.InitGlobal()

	if err != nil {
		fmt.Printf("初始化失败 %v", err)
		panic(err)
	}

}

func main() {

	gin.SetMode(global.Server.Mode)

	router := router.CreateRoute()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.Server.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(global.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(global.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// utils.Print()

	if err := s.ListenAndServe(); err != nil {
		fmt.Printf("服务器启动失败 %v", err)
	}

}
