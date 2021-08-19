package controller

import (
	"net/http"
	"server/mdl"
	"server/srvc"
	"server/excp"
	"github.com/gin-gonic/gin"
)

func UserCreate(c *gin.Context) {

}

func UserRead(c *gin.Context) {

}

func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}

func UserList(c *gin.Context) {
	users, err := service.UserList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"users": users,
		})
	}
}

func Login(c *gin.Context) {
	token, err := service.Login()
	switch err {
		case nil:
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"token": token,
			})
		case exception.Unauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "failure",
				"error": "unauthorized",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "failure",
				"error": "something wrong",
			})
	}
}
