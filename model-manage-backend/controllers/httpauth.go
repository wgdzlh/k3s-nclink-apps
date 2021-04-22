package controllers

import (
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AuthController is for auth logic
type AuthController struct{}

func (a AuthController) Login(c *gin.Context) {
	var loginInfo entity.User
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}

	user, err := service.UserServ.FindByName(loginInfo.Name)
	if err != nil {
		rest.Unauthorized(c, "User not found.")
		return
	}

	if user.Access != service.UserServ.AccessType {
		rest.Forbidden(c, "User access limited.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		rest.Forbidden(c, "Password invalid.")
		return
	}

	token, err := service.UserServ.GetJwtToken(user)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}

	rest.RetKV(c, "token", token)
}
