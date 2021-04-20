package controllers

import (
	"k3s-nclink-apps/data-source/service"

	"golang.org/x/crypto/bcrypt"
)

// AuthController is for auth logic
type AuthController struct{}

type WrongAccessError struct{}

func (e WrongAccessError) Error() string {
	return "wrong user access"
}

func (a AuthController) Login(name, pass string) (token string, err error) {
	user, err := service.UserServ.FindByName(name)
	if err != nil {
		return
	}

	if user.Access != service.UserServ.AccessType {
		err = WrongAccessError{}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return
	}

	token, err = service.UserServ.GetJwtToken(user)
	return
}
