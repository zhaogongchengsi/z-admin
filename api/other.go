package api

import (
	"z-admin/controller"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func OtherRouter(c *gin.Engine) {
	// User Router Group (api/user)
	other := c.Group("/api/other")

	{
		other.GET("/appinit", InitApp)
		other.GET("/verify", Verify)
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

type VerifyType struct {
	Id   string `json:"id"`
	Code string `json:"code"`
}

func Verify(c *gin.Context) {
	o := new(controller.OtherController)
	var v = VerifyType{}

	id, code, err := o.Verify()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	v.Id = id
	v.Code = code

	response.SuccessResponse(v).Send(c)
}
