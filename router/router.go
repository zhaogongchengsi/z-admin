package router

import (
	"github.com/gin-gonic/gin"
	"z-admin/controller"
)

func CreateRoute(c *gin.Engine) {

	user := new(controller.User)

	routerGroup := c.Group("/home")
	{
		routerGroup.GET("/", user.Home)
	}
}
