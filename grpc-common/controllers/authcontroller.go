package controllers

import (
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

func (a AuthController) Login(name, pass string) (token string, err error) {
	user, err := a.userservice.FindByName(name)
	if err != nil {
		return
	}

	if user.Access != service.UserAccessType {
		err = WrongAccessError{}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return
	}

	token, err = a.userservice.GetJwtToken(user)
	return
}
