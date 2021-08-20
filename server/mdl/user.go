package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"server/excp"
)

type User struct {
	gorm.Model
	ScreenName string
	Name string
	Description string
	GithubUserId uint
	GithubAccessToken string
	Icon string
	Works []Work
}

type CreateUserParam struct {
	ScreenName string
	Name string
	Description string
	Icon string
	GithubUserId uint
	GithubAccessToken string
}

type UpdateUserParam struct {
	ScreenName string
	Name string
	Description string
	Icon string
}

type UpdateUserAccessTokenParam struct {
	GithubAccessToken string
}

type LoginParams struct {
	Code string
}

type GithubOAuthResponse struct {
	AccessToken string
}

type GithubUserResponse struct {
	Id uint
	Login string
	Name string
	AvatarUrl string
	Bio string
}

type Auth struct {
	Id uint `json: "id"`
	jwt.StandardClaims
}

func GetLoginUser(c *gin.Context) User {
	contextLoginUser, _ := c.Get("LoginUser")
	loginUser, _ := contextLoginUser.(model.User)
	return loginUser
}

func GithubAuth(code string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", nil)
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var parsed_res GithubOAuthResponse
	if err := json.Unmarshal(body, &parsed_res); err != nil {
		return "", err
	}
	return parsed_res.AccessToken, nil
}

func GetGithubUser(token string) (GithubUserResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/user", nil)
	if err != nil {
		return GithubUserResponse{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return GithubUserResponse{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GithubUserResponse{}, err
	}
	var parsed_res GithubUserResponse
	if err := json.Unmarshal(body, &parsed_res); err != nil {
		return GithubUserResponse{}, err
	}
	return parsed_res, nil
}

func GenerateJwtToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Auth{
		Id: user.ID,
	})
	tokenstr, err := token.SignedString([]byte(os.Getenv("AUTH_SECRET")))
	if err != nil {
		log.Fatalln(err)
	}
	return tokenstr
}

func AuthJwtToken(tokenstr string) (uint, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.InvalidToken
		}
		return []byte(os.Getenv("AUTH_SECRET")), nil
	})
	if claims, ok := token.Claims.(Auth); ok && token.Valid {
		return claims.Id, nil
	}else{
		return 1, err
	}
}
