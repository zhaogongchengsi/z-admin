package api

import (
	"z-admin/controller"
	"z-admin/middleware"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Engine) {
	// role Router Group (api/role)
	Auth := c.Group("/api/role").Use(middleware.Authorization())
	{
		Auth.GET("/index", Role)
	}
}

func Role(c *gin.Context) {

	role := new(controller.RoleController)
	err := role.Rule()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(nil).Send(c)
}
