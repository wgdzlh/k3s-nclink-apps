package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type modelService struct {
	sync.Mutex
	DummyService
}

var ModelServ = &modelService{}

func init() {
	ModelServ.SetColl(&entity.Model{})
}

func (m *modelService) FindById(id string) (*entity.Model, error) {
	ret, err := m.DummyService.FindById(id)
	return ret.(*entity.Model), err
}

func (m *modelService) FindAll() ([]entity.Model, error) {
	ret := []entity.Model{}
	if err := m.coll.SimpleFind(&ret, bson.M{}); err != nil {
		return nil, err
	}
	return ret, nil
}

func (m *modelService) FindWithFilter(filters map[string]string) ([]entity.Model, int64, error) {
	ret := []entity.Model{}
	fil, opts := TransFilters(filters)
	err := m.coll.SimpleFind(&ret, fil, opts)
	if err != nil {
		return nil, 0, err
	}
	num, err := m.coll.CountDocuments(context.Background(), fil)
	return ret, num, err
}

func (m *modelService) delete(model *entity.Model) error {
	err := m.coll.Delete(model)
	if err != nil {
		return err
	}
	if model.Id != "" && model.Used > 0 {
		return AdapterServ.RenameModelFrom(model.Id, "")
	}
	return nil
}

func (m *modelService) DeleteById(id string) error {
	model, err := m.FindById(id)
	if err != nil {
		return err
	}
	return m.delete(model)
}

func (m *modelService) update(model *entity.Model) error {
	return m.coll.Update(model)
}

func (m *modelService) updateDef(model *entity.Model, def string) (changed bool, err error) {
	if model.Def != def {
		model.Def = def
		err = m.coll.Update(model)
		changed = err == nil
		if !changed || model.Used == 0 {
			return
		}
		adapters, err := AdapterServ.FindByModelId(model.Id)
		if err != nil {
			return changed, err
		}
		AdapterServ.ResetModel(adapters...)
	}
	return
}

func (m *modelService) UpdateById(id string, in *entity.Model) (changed bool, err error) {
	model, err := m.FindById(id)
	if err != nil {
		return
	}
	return m.updateDef(model, in.Def)
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
	err = m.update(model)
	if err == nil && model.Used > 0 {
		return AdapterServ.RenameModelFrom(id, newId)
	}
	return err
}
