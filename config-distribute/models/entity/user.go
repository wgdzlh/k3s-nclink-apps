package entity

import (
	_ "k3s-nclink-apps/config-distribute/models/db"

	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Password         string `json:"password" bson:"password"`
}

func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}
