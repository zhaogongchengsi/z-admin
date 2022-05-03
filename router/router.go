package router

import (
	"z-admin/api"
	"z-admin/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute(c *gin.Engine) {
	// User Router Group (api/)
	c.Use(middleware.Cors()) // 全局中间件
	api.UserRouter(c)
	api.OtherRouter(c)
	api.Auth(c)
}
