package controllers

import (
	"k3s-nclink-apps/config-distribute/models/entity"
	"k3s-nclink-apps/config-distribute/models/service"
)

// DistController is for model distribution logic
type DistController struct {
	adapterservice service.AdapterService
	modelservice   service.ModelService
}

func (a DistController) Fetch(hostname string) (model *entity.Model, devId string, err error) {
	adapter, err := a.adapterservice.FindByName(hostname)
	if err != nil {
		return
	}
	devId = adapter.DevId
	model, err = a.modelservice.FindByName(adapter.ModelName)
	return
}
