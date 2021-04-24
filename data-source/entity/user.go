package entity

import (
	_ "k3s-nclink-apps/data-source/db"

	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Access           string `json:"access" bson:"access"`
	Password         string `json:"password" bson:"password"`
}
