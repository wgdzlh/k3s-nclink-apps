package entity

import (
	_ "k3s-nclink-apps/data-source/db"

	"github.com/kamva/mgm/v3"
)

type Adapter struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	DevId            string `json:"dev_id" bson:"dev_id"`
	ModelName        string `json:"model_name" bson:"model_name"`
}

func NewAdapter(name, devId, modelName string) *Adapter {
	return &Adapter{
		Name:      name,
		DevId:     devId,
		ModelName: modelName,
	}
}
