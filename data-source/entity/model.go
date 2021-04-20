package entity

import (
	_ "k3s-nclink-apps/data-source/db"

	"github.com/kamva/mgm/v3"
)

type Model struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Used             uint32 `json:"used" bson:"used"`
	Def              string `json:"def" bson:"def"`
}

func NewModel(name, def string) *Model {
	return &Model{
		Name: name,
		Def:  def,
	}
}
