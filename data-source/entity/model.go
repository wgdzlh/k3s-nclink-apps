package entity

import (
	_ "k3s-nclink-apps/data-source/db"

	"github.com/kamva/mgm/v3"
)

type Model struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string `json:"id" bson:"id"`
	Used             uint32 `json:"used" bson:"used"`
	Def              string `json:"def" bson:"def"`
}

type Adapter struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string `json:"id" bson:"id"`
	DevId            string `json:"dev_id" bson:"dev_id"`
	ModelId          string `json:"model_id" bson:"model_id"`
}
