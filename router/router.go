package router

import (
	"z-admin/api"

	"github.com/gin-gonic/gin"
)

func CreateRoute(c *gin.Engine) {
	// User Router Group (api/user)
	api.UserRouter(c)
	api.OtherRouter(c)
}
