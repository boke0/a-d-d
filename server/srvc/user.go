package service

import (
	"os"
	"server/mdl"
	"server/excp"
)

var (
	github_client_id = os.Getenv("GITHUB_CLIENT_ID")
	github_client_secret = os.Getenv("GITHUB_CLIENT_SECRET")
)

func UserCreate(params model.UserCreateParam) (model.User, error) {
	user := model.User {
		ScreenName: params.ScreenName,
		Name: params.Name,
		Icon: params.Icon,
		Description: params.Description,
		GithubUserId: params.GithubUserId,
		GithubAccessToken: params.GithubAccessToken,
	}
	result := Db.Create(user)
	return user, result.Error
}

func UserRead(id uint) (model.User, error) {
	var user model.User
	result := Db.First(&user, id)
	return user, result.Error
}

func UserUpdate(loginUser model.User, params model.UpdateUserParam) (model.User, error) {
	var user model.User
	user.ScreenName = params.ScreenName
	user.Name = params.Name
	user.Description = params.Description
	user.Icon = params.Icon
	result := Db.Save(&user)
	return user, result.Error
}

func UserDelete(loginUser model.User) (model.User, error) {
	result := Db.Delete(&loginUser)
	return loginUser, result.Error
}

func UserList() ([]model.User, error) {
	var users []model.User
	result := Db.Find(&users)
	return users, result.Error
}

func Authorize(auth string) (model.User, error) {
	var user model.User
	id, err := model.AuthJwtToken(auth)
	if err != nil {
		return model.User{}, nil
	}
	result := Db.Find(&user, id)
	return user, result.Error
}

func Login(params model.LoginParams) (string, error) {
	var user model.User
	token, err := model.GithubAuth(params.Code)
	if err != nil {
		return "", exception.Unauthorized
	}
	githubUser, err := model.GetGithubUser(token)
	if err != nil {
		return "", err
	}
	result := Db.Where("GithubUserId", githubUser.Id).First(&user)
	if result.Error != nil {
		user, err := UserCreate(model.UserCreateParam{
			Name: githubUser.Name,
			ScreenName: githubUser.Login,
			Description: githubUser.Bio,
			Icon: githubUser.AvatarUrl,
			GithubAccessToken: token,
			GithubUserId: githubUser.Id,
		})
		if err != nil {
			return "", err
		}
		return model.GenerateJwtToken(user), nil
	}else{
		return model.GenerateJwtToken(user), nil
	}
}
