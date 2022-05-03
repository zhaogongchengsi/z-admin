package api

import (
	"z-admin/controller"
	"z-admin/pak/response"
	"z-admin/utils"

	"github.com/gin-gonic/gin"
)

func UserRouter(c *gin.Engine) {
	// User Router Group (api/user)
	user := c.Group("/api/user")
	{
		user.POST("/register", CreateUser)
		user.POST("/login", Login)
		user.POST("/changepass", ChangePass)
	}
}

// 注册
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

// Login 登陆
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

	token, err := utils.CreateJwt(u.UId, u.UserName)

	if err != nil {
		response.FailureResponse("token 签发失败").Send(c)
		return
	}

	res := map[string]interface{}{
		"z_token": token,
		"user":    u,
	}

	response.SuccessResponse(res).Send(c)
}

type ChangePassStruct struct {
	Password    string `json:"password"`
	Newpassword string `json:"newpassword"`
	Username    string `json:"username"`
}

// 修改密码
func ChangePass(c *gin.Context) {
	var user ChangePassStruct
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.ErrorResponse(err).Send(c)
		return
	}

	newuser := controller.User{
		UserName: user.Username,
		Password: user.Password,
	}

	err = newuser.ChangePass(user.Newpassword)

	if err != nil {
		response.FailureResponse(err.Error()).Send(c)
		return
	}

	response.SuccessResponse(nil).Send(c)
}
