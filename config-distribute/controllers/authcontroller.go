package controllers

import (
	"k3s-nclink-apps/config-distribute/models/entity"
	"k3s-nclink-apps/config-distribute/models/service"

	"golang.org/x/crypto/bcrypt"
)

// AuthController is for auth logic
type AuthController struct {
	userservice service.UserService
}

func (a AuthController) Login(loginInfo *entity.User) (token string, err error) {
	user, err := a.userservice.Find(loginInfo)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		return
	}

	token, err = a.userservice.GetJwtToken(user)
	return
}
