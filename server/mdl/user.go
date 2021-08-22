package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"os"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"server/excp"
)

var (
	github_client_id = os.Getenv("GITHUB_CLIENT_ID")
	github_client_secret = os.Getenv("GITHUB_CLIENT_SECRET")
)

type User struct {
	gorm.Model
	ScreenName string `gorm:"index"`
	Name string
	Description string
	GithubUserId uint `gorm:"index:uniqueIndex" json:"-"`
	GithubAccessToken string `json:"-"`
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
	AccessToken string `json:"access_token"`
}

type GithubOAuthErrorResponse struct {
	Error string `json:"error"`
}

type GithubUserResponse struct {
	Id uint `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Bio string `json:"bio"`
}

type Auth struct {
	Id uint `json: "id"`
	jwt.StandardClaims
}

func GetLoginUser(c *gin.Context) User {
	contextLoginUser, _ := c.Get("LoginUser")
	loginUser, _ := contextLoginUser.(User)
	return loginUser
}

func GithubAuth(code string) (string, error) {
	payload := url.Values{}
	payload.Add("client_id", github_client_id)
	payload.Add("client_secret", github_client_secret)
	payload.Add("code", code)
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		strings.NewReader(payload.Encode()),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		var parsed_res GithubOAuthErrorResponse
		if err := json.Unmarshal(body, &parsed_res); err != nil {
			return "", err
		}
		return "", errors.New(parsed_res.Error)
	}else{
		var parsed_res GithubOAuthResponse
		if err := json.Unmarshal(body, &parsed_res); err != nil {
			return "", err
		}
		return parsed_res.AccessToken, nil
	}
}

func GetGithubUser(token string) (GithubUserResponse, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", "token " + token)
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
