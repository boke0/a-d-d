package service

import (
	. "server/mdl"
)

func UserCreate(params CreateUserParam) (User, error) {
	user := User {
		ScreenName: params.ScreenName,
		Name: params.Name,
		Icon: params.Icon,
		Description: params.Description,
		GithubUserId: params.GithubUserId,
		GithubAccessToken: params.GithubAccessToken,
	}
	result := Db.Create(&user)
	return user, result.Error
}

func UserRead(id uint) (User, error) {
	var user User
	result := Db.Preload("Works").Preload("Works.Drinks").First(&user, id)
	return user, result.Error
}

func UserReadByScreenName(screenName string) (User, error) {
	var user User
	result := Db.Preload("Works").Preload("Works.Drinks").First(&user, &User{ ScreenName: screenName })
	return user, result.Error
}

func UserUpdate(loginUser User, params UpdateUserParam) (User, error) {
	var user User
	user.ScreenName = params.ScreenName
	user.Name = params.Name
	user.Description = params.Description
	user.Icon = params.Icon
	result := Db.Save(&user)
	return user, result.Error
}

func UserUpdateGithubAccessToken(loginUser User, params UpdateUserAccessTokenParam) (User, error) {
	var user User
	user.GithubAccessToken = params.GithubAccessToken
	result := Db.Save(&user)
	return user, result.Error
}

func UserDelete(loginUser User) (User, error) {
	result := Db.Delete(&loginUser)
	return loginUser, result.Error
}

func UserList() ([]User, error) {
	var users []User
	result := Db.Find(&users)
	return users, result.Error
}

func Authorize(auth string) (User, error) {
	var user User
	id, err := AuthJwtToken(auth)
	if err != nil {
		return User{}, nil
	}
	result := Db.First(&user, id)
	return user, result.Error
}

func Login(params LoginParams) (string, error) {
	var user User
	token, err := GithubAuth(params.Code)
	if err != nil {
		return "", err
	}
	githubUser, err := GetGithubUser(token)
	if err != nil {
		return "", err
	}
	result := Db.Where(&User{ GithubUserId: githubUser.Id }).First(&user)
	if result.Error != nil {
		user, err := UserCreate(CreateUserParam{
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
		return GenerateJwtToken(user), nil
	}else{
		user, err := UserUpdateGithubAccessToken(user, UpdateUserAccessTokenParam {
			GithubAccessToken: token,
		})
		return GenerateJwtToken(user), err
	}
}
