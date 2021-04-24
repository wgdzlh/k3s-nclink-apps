package service

import "k3s-nclink-apps/data-source/entity"

type User entity.User

func NewUser(name, access, password string) *User {
	return &User{
		Name:     name,
		Access:   access,
		Password: password,
	}
}

type Model entity.Model

func NewModel(id, def string) *Model {
	return &Model{
		Id:  id,
		Def: def,
	}
}

type Adapter entity.Adapter

func NewAdapter(id, devId, modelId string) *Adapter {
	return &Adapter{
		Id:      id,
		DevId:   devId,
		ModelId: modelId,
	}
}
