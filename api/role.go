package api

import (
	"fmt"
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

	username, nok := c.Get("username")
	uid, uok := c.Get("uid")

	if !nok || !uok {
		response.FailureResponse("获取用户信息失败").Send(c)
		return
	}

	fmt.Printf("username: %s, uid: %s", username, uid)

	role := new(controller.RoleController)
	err := role.Rule()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(nil).Send(c)
}
