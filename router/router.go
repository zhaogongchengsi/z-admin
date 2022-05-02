package router

import (
	"z-admin/api"
	"z-admin/middleware"

	"github.com/gin-gonic/gin"
)

func CreateRoute(c *gin.Engine) {
	// User Router Group (api/)
	c.Use(middleware.Cors())
	api.UserRouter(c)
	api.OtherRouter(c)
}
