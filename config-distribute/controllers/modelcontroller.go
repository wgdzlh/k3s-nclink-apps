package controllers

import (
	"k3s-nclink-apps/config-distribute/models/entity"
	"k3s-nclink-apps/config-distribute/models/service"
)

// ModelController is for auth logic
type ModelController struct {
	modelservice service.ModelService
}

func (a ModelController) Fetch(hostname string) (model *entity.Model, err error) {
	model, err = a.modelservice.FindByHost(hostname)
	return
}
