package auth

import (
	"strconv"

	"github.com/znbang/eventmap/pkg/userservice"
)

type UserInfo interface {
	ToUser() userservice.User
}

type GoogleUserInfo struct {
	Sub   string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *GoogleUserInfo) ToUser() userservice.User {
	return userservice.User{
		Provider: userservice.GOOGLE,
		Uid:      u.Sub,
		Email:    u.Email,
		Name:     u.Name,
	}
}

type FacebookUserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *FacebookUserInfo) ToUser() userservice.User {
	return userservice.User{
		Provider: userservice.FACEBOOK,
		Uid:      u.ID,
		Email:    u.Email,
		Name:     u.Name,
	}
}

type GithubUserInfo struct {
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func (u *GithubUserInfo) ToUser() userservice.User {
	return userservice.User{
		Provider: userservice.GITHUB,
		Uid:      strconv.FormatInt(u.ID, 10),
		Email:    u.Login,
		Name:     u.Name,
	}
}

type LineUserInfo struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
}

func (u *LineUserInfo) ToUser() userservice.User {
	return userservice.User{
		Provider: userservice.LINE,
		Uid:      u.UserID,
		Email:    u.DisplayName,
		Name:     u.DisplayName,
	}
}
