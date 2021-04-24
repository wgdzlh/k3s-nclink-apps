package service

import (
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type modelService struct {
	sync.Mutex
	mongoService
}

var ModelServ = &modelService{}

func init() {
	ModelServ.setColl(&Model{})
}

// Model hooks
func (d *Model) Creating() error {
	// Call the DefaultModel Creating hook
	if err := d.DefaultModel.Creating(); err != nil {
		return err
	}
	d.Used = 0
	return nil
}

func (d *Model) Deleted(result *mongo.DeleteResult) error {
	log.Printf("deleted model %s.\n", d.Id)
	if d.Id != "" && d.Used > 0 {
		return AdapterServ.RenameModelFrom(d.Id, "")
	}
	return nil
}

func (m *modelService) New() interface{} {
	return &Model{}
}

func (m *modelService) IdOf(in interface{}) string {
	return in.(*Model).Id
}

func (m *modelService) Dup(id string, in interface{}) interface{} {
	model := in.(*Model)
	return &Model{Id: id, Def: model.Def}
}

func (m *modelService) Slice() interface{} {
	return &[]Model{}
}

func (m *modelService) LenOf(in interface{}) int64 {
	return int64(len(*in.(*[]Model)))
}

func (m *modelService) DeleteById(id string) error {
	model := &Model{}
	if err := m.FindById(id, model); err != nil {
		return err
	}
	return m.delete(model)
}

func (m *modelService) UpdateById(id string, in interface{}) (changed bool, err error) {
	def := in.(*Model).Def
	model := &Model{}
	if err = m.FindById(id, model); err != nil || model.Def == def {
		return
	}
	model.Def = def
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
	model := &Model{}
	if err := m.FindById(id, model); err != nil {
		return err
	}
	model.Id = newId
	if err := m.update(model); err != nil {
		return err
	}
	if model.Used > 0 {
		return AdapterServ.RenameModelFrom(id, newId)
	}
	return nil
}
