package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"z-admin/global"
)

type User struct {}

func (u *User) Home(c *gin.Context)  {
	data := map[string]interface{}{
		"foo": "bar",
		"port": global.Server.Port,
	}

	c.JSONP(http.StatusOK, data)
}