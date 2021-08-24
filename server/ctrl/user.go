package controller

import (
	"net/http"
	"server/excp"
	. "server/mdl"
	"server/srvc"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/render"
)

func UserRead(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "user")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	var user User
	if err != nil {
		user, err = service.UserReadByScreenName(userIdStr)
	}else{
		user, err = service.UserRead(uint(userId))
	}
	if err != nil {
		render.Status(r, 404)
		render.Render(w, r, CreateErrorResponse(err))
		return
	}
	render.Render(w, r, CreateUserResponse(user))
}

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User) 
	var req UpdateUserParam
	if err := render.Bind(r, &req); err != nil {
		render.Render(r, w, CreateErrorResponse(err))
	}
	user, err := service.UserUpdate(loginUser, req)
	if err != nil {
		render.Render(r, w, CreateErrorResponse(err))
		return
	}
	render.Render(r, w, CreateUserResponse(user))
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	loginUser, _ := c.Value("LoginUser").(User) 
	user, err := service.UserDelete(loginUser)
	if err != nil {
		render.Render(r, w, CreateErrorResponse(err))
		return
	}
	render.Render(r, w, CreateUserResponse(user))
}

func UserList(w http.ResponseWriter, r *http.Request) {
	users, err := service.UserList()
	if err != nil {
		render.Render(r, w, CreateErrorResponse(err))
	}else{
		render.Render(r, w, CreateUsersResponse(users))
	}
}

func Session(w http.ResponseWriter, r *http.Request) {
	c := r.Context()
	user, _ := c.Value("LoginUser").(User)
	render.Render(r, w, CreateUserResponse(user))
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginParams
	if err := render.Bind(r, &req); err != nil {
		render.Render(r, w, CreateErrorResponse(err))
		return
	}
	token, err := service.Login(req)
	if err != nil {
		switch err {
			case exception.Unauthorized:
				render.Render(r, w, CreateErrorResponse(err))
			default:
				render.Render(r, w, CreateErrorResponse(err))
		}
	}else{
		render.Render(r, w, CreateTokenResponse(token))
	}
}
