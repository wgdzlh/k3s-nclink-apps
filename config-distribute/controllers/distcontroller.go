package controllers

import (
	"k3s-nclink-apps/data-source/service"
)

// DistController is for model distribution logic
type DistController struct{}

func (a DistController) Fetch(hostname string) (model *service.Model, devId string, err error) {
	adapter := &service.Adapter{}
	if err = service.AdapterServ.FindById(hostname, adapter); err != nil {
		return
	}
	devId = adapter.DevId
	model = &service.Model{}
	err = service.ModelServ.FindById(adapter.ModelId, model)
	return
}
