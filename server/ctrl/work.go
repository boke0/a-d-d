package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/mdl"
	"server/srvc"
	"strconv"
)

func WorkCreate(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	var req model.CreateWorkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	work, err := service.WorkCreate(loginUser, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"work": work,
	})
}

func WorkRead(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	work, err := service.WorkRead(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"work": work,
	})
}

func WorkUpdate(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req model.UpdateWorkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	work, err := service.WorkUpdate(loginUser, uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"work": work,
	})
}

func WorkDelete(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	work, err := service.WorkDelete(loginUser, uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"work": work,
	})
}

func WorkList(c *gin.Context) {
	works, err := service.WorkList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"works": works,
	})
}

func DrinkCreate(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req model.CreateDrinkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	drink, err := service.DrinkCreate(loginUser, uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"drink": drink,
	})
}

func DrinkRead(c *gin.Context) {
	work, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	drink, err := service.DrinkRead(uint(work), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"drink": drink,
	})
}

func DrinkUpdate(c *gin.Context) {
	loginUser := model.GetLoginUser(c)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req model.UpdateWorkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	work, err := service.WorkUpdate(loginUser, uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"work": work,
	})
}

func DrinkList(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinks, err := service.DrinkList(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"drinks": drinks,
	})
}
