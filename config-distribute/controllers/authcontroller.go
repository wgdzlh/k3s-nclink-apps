package controllers

import (
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"

	"golang.org/x/crypto/bcrypt"
)

// AuthController is for auth logic
type AuthController struct {
	userservice service.UserService
}

type WrongAccessError struct{}

func (e WrongAccessError) Error() string {
	return "wrong user access"
}

func (a AuthController) Login(loginInfo *entity.User) (token string, err error) {
	user, err := a.userservice.Find(loginInfo)
	if err != nil {
		return
	}

	if user.Access != "ro" {
		err = WrongAccessError{}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		return
	}

	token, err = a.userservice.GetJwtToken(user)
	return
}
