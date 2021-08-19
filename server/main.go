package main

import (
	"github.com/gin-gonic/gin"
	"server/ctrl"
	"server/mdlw"
)

func main() {
	engine := gin.Default()
	v1 := engine.Group("/v1")
	{
		v1.GET("/login", controller.Login)
		v1.GET("/users", controller.UserList)
		v1.GET("/users/:id", controller.UserRead)
		v1.GET("/works", controller.WorkList)
		v1.GET("/works/:id", controller.WorkRead)
		v1.GET("/works/:work/drinks", controller.DrinkList)
		v1.GET("/works/:work/drinks/:id", controller.DrinkRead)
		auth := v1.Group("/")
		{
			auth.Use(middleware.Auth)
			auth.PUT("/users/:id", controller.UserUpdate)
			auth.DELETE("/users/:id", controller.UserDelete)
			auth.POST("/works", controller.WorkCreate)
			auth.PUT("/works/:id", controller.WorkUpdate)
			auth.DELETE("/works/:id", controller.WorkDelete)
			auth.POST("/works/:work/drinks/:id", controller.DrinkCreate)
			auth.PUT("/works/:work/drinks/:id", controller.DrinkUpdate)
		}
	}
	engine.Run(":3000")
}
