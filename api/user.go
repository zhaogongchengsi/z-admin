package api

import (
	"z-admin/controller"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Engine) {
	// User Router Group (api/user)
	user := c.Group("/api/user")
	{
		user.POST("/register", CreateUser)
		user.POST("/login", Login)
	}
}

func CreateUser(c *gin.Context) {
	var user controller.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		response.ErrorResponse(err).Send(c)
		return
	}

	u, err := user.Register()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(u).Send(c)
}

func Login(c *gin.Context) {
	var user controller.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		response.ErrorResponse(err).Send(c)
		return
	}

	u, err := user.Login()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(u).Send(c)
}
