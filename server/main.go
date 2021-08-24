package main

import (
	"github.com/go-chi/chi/v5"
	"server/ctrl"
	"server/mdlw"
	"net/http"
	"os"
)

func main() {
	r := chi.NewRouter()
	r.use(cors.Handler(cors.Options{
		AllowOrigins: []string {
			os.Getenv("CLIENT_BASE_URL"),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
	}))

	v1 := r.Group("/v1")
	{
		v1.POST("/login", controller.Login)
		v1.GET("/users", controller.UserList)
		v1.GET("/users/:user", controller.UserRead)
		v1.GET("/works", controller.WorkList)
		v1.GET("/works/:work", controller.WorkRead)
		v1.GET("/works/:work/drinks", controller.DrinkList)
		v1.GET("/works/:work/drinks/:drink", controller.DrinkRead)
		auto := v1.Group("/")
		{
			auth.Use(middleware.Auth)
			auth.GET("/session", controller.Session)
			auth.PUT("/users/:user", controller.UserUpdate)
			auth.DELETE("/users/:user", controller.UserDelete)
			auth.POST("/works", controller.WorkCreate)
			auth.GET("/works/in_progress", controller.WorkInProgressRead)
			auth.PUT("/works/:work", controller.WorkUpdate)
			auth.DELETE("/works/:work", controller.WorkDelete)
			auth.POST("/works/:work/drinks", controller.DrinkCreate)
			auth.PUT("/works/:work/drinks/:drink", controller.DrinkUpdate)
		}
	}
	http.ListenAndServe(":8000", r)
}
