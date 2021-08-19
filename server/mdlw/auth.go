package middleware

import (
	"net/http"
	"server/srvc"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	user, err := service.Authorize(auth)
	if err != nil {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	}else{
		c.Set("LoginUser", user)
		c.Next()
	}
}
