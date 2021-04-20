package controllers

import (
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
)

// DistController is for model distribution logic
type DistController struct{}

func (a DistController) Fetch(hostname string) (model *entity.Model, devId string, err error) {
	adapter, err := service.AdapterServ.FindByName(hostname)
	if err != nil {
		return
	}
	devId = adapter.DevId
	model, err = service.ModelServ.FindByName(adapter.ModelName)
	return
}
