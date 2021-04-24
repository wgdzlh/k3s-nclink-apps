package service

import (
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type modelService struct {
	sync.Mutex
	DummyService
}

var ModelServ = &modelService{}

func init() {
	ModelServ.setColl(&Model{})
}

// Model hooks
func (d *Model) Deleted(result *mongo.DeleteResult) error {
	log.Printf("deleted model %s.\n", d.Id)
	if d.Id != "" && d.Used > 0 {
		return AdapterServ.RenameModelFrom(d.Id, "")
	}
	return nil
}

// Model CRUDs
func (m *modelService) Save(model *Model) error {
	return m.create(model)
}

func (m *modelService) FindById(id string) (*Model, error) {
	ret := &Model{}
	if err := m.findById(id, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *modelService) FindAll() ([]Model, error) {
	ret := []Model{}
	if err := m.findAll(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *modelService) FindWithFilter(filters map[string]string) ([]Model, int64, error) {
	ret := []Model{}
	num, err := m.findWithFilter(filters, &ret)
	if err != nil {
		return nil, 0, err
	}
	return ret, num, nil
}

func (m *modelService) DeleteById(id string) error {
	model, err := m.FindById(id)
	if err != nil {
		return err
	}
	return m.delete(model)
}

func (m *modelService) UpdateById(id string, in *Model) (changed bool, err error) {
	model, err := m.FindById(id)
	if err != nil || model.Def == in.Def {
		return
	}
	model.Def = in.Def
	err = m.update(model)
	changed = err == nil
	if !changed || model.Used == 0 {
		return
	}
	adapters, err := AdapterServ.FindByModelId(model.Id)
	if err != nil {
		return changed, err
	}
	AdapterServ.ResetModel(adapters...)
	return
}

func (m *modelService) Rename(id, newId string) error {
	if id == newId {
		return nil
	}
	model, err := m.FindById(id)
	if err != nil {
		return err
	}
	model.Id = newId
	if err = m.update(model); err == nil && model.Used > 0 {
		return AdapterServ.RenameModelFrom(id, newId)
	}
	return err
}
