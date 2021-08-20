package controller

import (
	"net/http"
	"server/excp"
	"server/mdl"
	"server/srvc"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRead(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := service.UserRead(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"status": "success",
		"user": user,
	})
}

func UserUpdate(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	var req model.UpdateUserParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	user, err := service.UserUpdate(loginUser, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"status": "success",
		"user": user,
	})
}

func UserDelete(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	user, err := service.UserDelete(loginUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H {
		"status": "success",
		"user": user,
	})
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
	var req model.LoginParams
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	token, err := service.Login(req)
	if err != nil {
		switch err {
			case exception.Unauthorized:
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "failure",
					"error": err.Error(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "failure",
					"error": err.Error(),
				})
		}
	}else{
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"token": token,
		})
	}
}
