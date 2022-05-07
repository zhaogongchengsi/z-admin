package api

import (
	"strconv"
	"z-admin/controller"
	"z-admin/pak/response"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Engine) {
	// role Router Group (api/role)
	Auth := c.Group("/api/role")
	{
		Auth.POST("/createRole", CreateRole)
		Auth.POST("/setPermis", SetPermis)
		Auth.GET("/getRolebyid", FindRoleById)
		Auth.GET("/getRoleList", FindRoleList)
		Auth.DELETE("/deleteRole", DeleteRole)
	}
}

func CreateRole(c *gin.Context) {

	role := new(controller.RoleController)

	err := c.ShouldBindJSON(role)

	if err != nil {
		response.ErrorResponse(err).Send(c)
		return
	}

	r, err := role.CreateRule()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(r).Send(c)
}

func SetPermis(c *gin.Context) {

	role := new(controller.RoleController)

	err := c.ShouldBindJSON(role)

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	err = role.SetPermis()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return

	}

	response.SuccessResponse(nil).Send(c)

}

func FindRoleById(c *gin.Context) {
	role := new(controller.RoleController)

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
	}

	role.RoleId = id

	r, err := role.FindRoleById()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(r).Send(c)
}

func FindRoleList(c *gin.Context) {

	role := new(controller.RoleController)

	r, err := role.FindRoleList()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(r).Send(c)

}

func DeleteRole(c *gin.Context) {
	role := new(controller.RoleController)

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	role.RoleId = id

	err = role.DeleteRole()

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(nil).Send(c)
}
