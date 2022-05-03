package router

import (
	"z-admin/api"
	"z-admin/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute() *gin.Engine {
	r := gin.Default()
	// User Router Group (api/)
	r.Use(middleware.Cors()) // 全局中间件
	api.UserRouter(r)
	api.OtherRouter(r)
	api.Auth(r)
	return r
}
