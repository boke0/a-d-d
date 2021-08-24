package controller

import (
	"json"
	"errors"
)

type Response struct {
	Status: string `json:"status"`
}

type ErrorResponse struct {
	*Response
	Error: string `json:"error"`
}

type UserResponse struct {
	*Response
	User: User
}

type WorkResponse struct {
	*Response
	Work: Work
}

type WorksResponse struct {
	*Response
	Works: []Work
}

type TokenResponse struct {
	*Response
	Token: string `json:"token"`
}

func CreateErrorResponse(err error) ErrorResponse {
	return ErrorResponse {
		Status: "failure"
		Error: err.Error()
	}
}

func CreateUserResponse(user User) UserResponse {
	return UserResponse {
		Status: "success"
		User: user
	}
}

func CreateWorkResponse(work Work) WorkResponse {
	return WorkResponse {
		Status: "success"
		Work: work
	}
}

func CreateWorksResponse(works []Work) WorksResponse {
	return WorksResponse {
		Status: "success"
		Works: works
	}
}

func CreateTokenResponse(token string) TokenResponse {
	return TokenResponse {
		Status: "success"
		Token: token
	}
}
