package middleware

import (
	"z-admin/utils"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {

	return func(c *gin.Context) {
		zToken := c.Request.Header.Get("ToKen")

		if zToken == "" {
			c.AbortWithStatus(401)
			return
		}

		claims, err := utils.ParseToken(zToken)

		if err != nil {
			c.AbortWithStatus(401)
			return
		}

		c.Set("username", claims.UserName)
		c.Set("uid", claims.Uid)

		c.Next()
	}

}
