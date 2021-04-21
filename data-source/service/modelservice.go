package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"
	"sync"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type modelService struct {
	sync.Mutex
	coll *mgm.Collection
}

var ModelServ = &modelService{
	coll: mgm.Coll(&entity.Model{}),
}

func (m *modelService) Create(model *entity.Model) error {
	err := m.coll.Create(model)
	if err != nil {
		return err
	}
	ctx := context.Background()
	num, err := m.coll.EstimatedDocumentCount(ctx)
	if num <= 1 {
		_, err = m.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}
	return err
}

func (m *modelService) Save(name, def string) error {
	return m.coll.Create(entity.NewModel(name, def))
}

// Find model
func (m *modelService) FindById(id string) (*entity.Model, error) {
	ret := &entity.Model{}
	err := m.coll.FindByID(id, ret)
	return ret, err
}

func (m *modelService) FindByName(name string) (*entity.Model, error) {
	ret := &entity.Model{}
	err := m.coll.First(bson.M{"name": name}, ret)
	return ret, err
}

func (m *modelService) FindAll() ([]entity.Model, int64, error) {
	ret := []entity.Model{}
	err := m.coll.SimpleFind(&ret, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	num, err := m.coll.EstimatedDocumentCount(context.Background())
	return ret, num, err
}

func (m *modelService) delete(model *entity.Model) error {
	err := m.coll.Delete(model)
	if err != nil {
		return err
	}
	if model.Name != "" && model.Used > 0 {
		return AdapterServ.RenameModelFrom(model.Name, "")
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

func (m *modelService) DeleteByName(name string) error {
	model, err := m.FindByName(name)
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
		adapters, err := AdapterServ.FindByModelName(model.Name)
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

func (m *modelService) UpdateByName(name string, in *entity.Model) (changed bool, err error) {
	model, err := m.FindByName(name)
	if err != nil {
		return
	}
	return m.updateDef(model, in.Def)
}

func (m *modelService) Rename(id, newName string) error {
	model, err := m.FindById(id)
	if err != nil {
		return err
	}
	oldName := model.Name
	if oldName != newName {
		model.Name = newName
		err = m.update(model)
		if err == nil && model.Used > 0 {
			err = AdapterServ.RenameModelFrom(oldName, newName)
		}
	}
	return err
}
