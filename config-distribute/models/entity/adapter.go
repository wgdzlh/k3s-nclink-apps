package entity

import (
	_ "k3s-nclink-apps/config-distribute/models/db"

	"github.com/kamva/mgm/v3"
)

type Adapter struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	ModelName        string `json:"model_name" bson:"model_name"`
}

func NewAdapter(name, modelName string) *Adapter {
	return &Adapter{
		Name:      name,
		ModelName: modelName,
	}
}
