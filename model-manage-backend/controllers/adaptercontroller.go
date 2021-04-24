package controllers

import "k3s-nclink-apps/data-source/service"

type AdapterController struct {
	dummyController
}

func NewAdapterController() *AdapterController {
	ac := &AdapterController{}
	ac.serv = service.AdapterServ
	return ac
}
