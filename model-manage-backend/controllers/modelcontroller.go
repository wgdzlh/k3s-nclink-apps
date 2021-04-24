package controllers

import "k3s-nclink-apps/data-source/service"

type ModelController struct {
	dummyController
}

func NewModelController() *ModelController {
	mc := &ModelController{}
	mc.serv = service.ModelServ
	return mc
}
