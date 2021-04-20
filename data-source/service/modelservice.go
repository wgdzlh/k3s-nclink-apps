package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"
	"log"
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

// Find model
func (m *modelService) FindByName(name string) (*entity.Model, error) {
	ret := &entity.Model{}
	err := m.coll.First(bson.M{"name": name}, ret)
	return ret, err
}

func (m *modelService) FindAll() ([]entity.Model, error) {
	ret := []entity.Model{}
	err := m.coll.SimpleFind(&ret, bson.M{})
	return ret, err
}

func (m *modelService) DeleteByName(name string) error {
	model, err := m.FindByName(name)
	if err != nil {
		return err
	}
	return m.coll.Delete(model)
}

func (m *modelService) UpdateByName(name, def string) (changed bool, err error) {
	model, err := m.FindByName(name)
	if err != nil {
		return
	}
	if model.Def != def {
		model.Def = def
		err = m.coll.Update(model)
		changed = err == nil
		if changed && model.Used > 0 {
			adapters, err := AdapterServ.FindByModelName(model.Name)
			if err != nil {
				log.Println("no adapters match model name:", model.Name)
			} else {
				AdapterServ.ResetModel(adapters...)
			}
		}
	}
	return
}

func (m *modelService) Update(model *entity.Model) error {
	return m.coll.Update(model)
}
