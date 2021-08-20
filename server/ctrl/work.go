package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "server/mdl"
	"server/srvc"
	"strconv"
)

func WorkCreate(c *gin.Context) {
	loginUser := GetLoginUser(c)
	var req CreateWorkParam
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
		"Work": work,
	})
}

func WorkRead(c *gin.Context) {
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	work, err := service.WorkRead(uint(workId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkInProgressRead(c *gin.Context) {
	loginUser := GetLoginUser(c)
	work, err := service.WorkInProgressRead(loginUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkUpdate(c *gin.Context) {
	loginUser := GetLoginUser(c)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	var req UpdateWorkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	work, err := service.WorkUpdate(loginUser, uint(workId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkDelete(c *gin.Context) {
	loginUser := GetLoginUser(c)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	work, err := service.WorkDelete(loginUser, uint(workId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
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
		"Works": works,
	})
}

func DrinkCreate(c *gin.Context) {
	loginUser := GetLoginUser(c)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	var req CreateDrinkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	drink, err := service.DrinkCreate(loginUser, uint(workId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drink": drink,
	})
}

func DrinkRead(c *gin.Context) {
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinkId, _ := strconv.ParseUint(c.Param("drink"), 10, 64)
	drink, err := service.DrinkRead(uint(workId), uint(drinkId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drink": drink,
	})
}

func DrinkUpdate(c *gin.Context) {
	loginUser := GetLoginUser(c)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinkId, _ := strconv.ParseUint(c.Param("drink"), 10, 64)
	var req UpdateDrinkParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
	}
	drink, err := service.DrinkUpdate(loginUser, uint(workId), uint(drinkId), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drink": drink,
	})
}

func DrinkList(c *gin.Context) {
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinks, err := service.DrinkList(uint(workId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drinks": drinks,
	})
}
