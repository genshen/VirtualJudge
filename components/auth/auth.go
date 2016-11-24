package auth

import (
	"errors"
)

const LenAccessTokenName = 13 //access_token=
const (
	UserStatusUnAuth = iota
	UserStatusAuthOK
	UserStatusAlreadyAuth
)

//auth types
const (
	Self = iota
	Github
)

type User struct {
	Id          int    `json:"id"`
	Status      int    `json:"status"`
	ThirdPartId string `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	Email       string `json:"-"`
	Avatar      string `json:"avatar"`
}

type AuthInterface interface {
	GetAccessToken(code string) string
	GetAuthUserInfo(access_token string) (*User, error)
}

func StartAuth(a AuthInterface, code string) (*User, error) {
	if access_token := a.GetAccessToken(code); access_token == "" {
		return &User{}, errors.New("can't fetch access_token")
	} else {
		return a.GetAuthUserInfo(access_token)
	}
}