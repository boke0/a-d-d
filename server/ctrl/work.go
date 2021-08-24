package controller

import (
	"net/http"
	. "server/mdl"
	"server/srvc"
	"strconv"
	"github.com/go-chi/chi/render"
)

func WorkCreate(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User)
	var req CreateWorkParam
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, CreateErrorResponse(err))
	}
	work, err := service.WorkCreate(loginUser, req)
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkRead(w http.ResponseWriter, r *http.Request) {
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	work, err := service.WorkRead(uint(workId))
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkInProgressRead(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User)
	work, err := service.WorkInProgressRead(loginUser)
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkUpdate(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	var req UpdateWorkParam
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	work, err := service.WorkUpdate(loginUser, uint(workId), req)
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkDelete(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	work, err := service.WorkDelete(loginUser, uint(workId))
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Work": work,
	})
}

func WorkList(w http.ResponseWriter, r *http.Request) {
	works, err := service.WorkList()
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Works": works,
	})
}

func DrinkCreate(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	var req CreateDrinkParam
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	drink, err := service.DrinkCreate(loginUser, uint(workId), req)
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drink": drink,
	})
}

func DrinkRead(w http.ResponseWriter, r *http.Request) {
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinkId, _ := strconv.ParseUint(c.Param("drink"), 10, 64)
	drink, err := service.DrinkRead(uint(workId), uint(drinkId))
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drink": drink,
	})
}

func DrinkUpdate(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User)
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinkId, _ := strconv.ParseUint(c.Param("drink"), 10, 64)
	var req UpdateDrinkParam
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	drink, err := service.DrinkUpdate(loginUser, uint(workId), uint(drinkId), req)
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drink": drink,
	})
}

func DrinkList(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	workId, _ := strconv.ParseUint(c.Param("work"), 10, 64)
	drinks, err := service.DrinkList(uint(workId))
	if err != nil {
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"Drinks": drinks,
	})
}
