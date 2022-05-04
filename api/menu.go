package api

import (
	"z-admin/controller"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func MenuRouter(c *gin.Engine) {
	// Menu Router Group (api/menu)
	other := c.Group("/api/menu")

	{
		other.GET("/getmemu", GetMenu)
		other.POST("/createmenu", CreateMenu)
	}
}

func GetMenu(c *gin.Context) {
	m := new(controller.MenuController)
	menulist, err := m.GetMenu()
	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}
	response.SuccessResponse(menulist).Send(c)
}

func CreateMenu(c *gin.Context) {

	m := new(controller.MenuController)
	err := c.ShouldBindJSON(m)

	if err != nil {
		response.ErrorResponse(err).Send(c)
		return
	}

	menu, err := m.CreateMenu()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(menu).Send(c)
}
