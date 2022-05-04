package router

import (
	"z-admin/api"
	"z-admin/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// User Router Group (api/)
	r.Use(middleware.Cors()) // 全局中间件
	api.UserRouter(r)
	api.OtherRouter(r)
	api.Auth(r)
	api.MenuRouter(r)

	return r
}
