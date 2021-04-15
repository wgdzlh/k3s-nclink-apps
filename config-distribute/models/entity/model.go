package entity

import (
	_ "k3s-nclink-apps/config-distribute/models/db"

	"github.com/kamva/mgm/v3"
)

type Model struct {
	mgm.DefaultModel `bson:",inline"`
	Hostname         string `json:"hostname" bson:"hostname"`
	Model            string `json:"model" bson:"model"`
}

func NewModel(hostname string, model string) *Model {
	return &Model{
		Hostname: hostname,
		Model:    model,
	}
}
