package api

import (
	"z-admin/controller"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Engine) {
	// User Router Group (api/user)
	user := c.Group("/api/user/register")
	{
		user.POST("/", CreateUser)
	}
}

func CreateUser(c *gin.Context) {
	var user controller.User
	c.ShouldBindJSON(&user)
	u, err := user.Register()

	if err != nil {
		response.ErrorResponse(err).Send(c)
		return
	}

	response.SuccessResponse(u).Send(c)
}
