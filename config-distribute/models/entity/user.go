package entity

import (
	_ "k3s-nclink-apps/config-distribute/models/db"
	"k3s-nclink-apps/utils"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Password         string `json:"password" bson:"password"`
	VerifiedAt       *time.Time
}

func NewUser(name string, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}

func (user *User) GetJwtToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
	})
	secretKey := utils.EnvVar("TOKEN_KEY", "")
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
