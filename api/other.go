package api

import (
	"z-admin/controller"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func OtherRouter(c *gin.Engine) {
	// User Router Group (api/user)
	user := c.Group("/api/other")
	{
		user.GET("/appinit", InitApp)
	}
}

func InitApp(c *gin.Context) {
	o := new(controller.OtherController)
	err := o.Appinit()
	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(nil).Send(c)
}
