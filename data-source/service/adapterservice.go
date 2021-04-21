package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/model-manage-backend/mqtt"
	"log"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type adapterService struct {
	coll *mgm.Collection
}

var AdapterServ = &adapterService{
	coll: mgm.Coll(&entity.Adapter{}),
}

func (a *adapterService) Create(adapter *entity.Adapter) error {
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model, err := ModelServ.FindByName(adapter.ModelName)
	if err != nil {
		return err
	}
	if err = a.coll.Create(adapter); err != nil {
		return err
	}
	model.Used++
	if err = ModelServ.update(model); err != nil {
		return err
	}
	ctx := context.Background()
	num, err := a.coll.EstimatedDocumentCount(ctx)
	if num <= 1 {
		_, err = a.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}
	return err
}

// Find adapter
func (a *adapterService) FindByName(name string) (*entity.Adapter, error) {
	ret := &entity.Adapter{}
	err := a.coll.First(bson.M{"name": name}, ret)
	return ret, err
}

func (a *adapterService) FindByModelName(modelName string) ([]entity.Adapter, error) {
	ret := []entity.Adapter{}
	err := a.coll.SimpleFind(&ret, bson.M{"model_name": modelName})
	return ret, err
}

func (a *adapterService) DeleteByName(name string) error {
	adapter, err := a.FindByName(name)
	if err != nil {
		return err
	}
	if err = a.coll.Delete(adapter); err != nil {
		return err
	}
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model, err := ModelServ.FindByName(adapter.ModelName)
	if model.Used > 0 {
		model.Used--
		err = ModelServ.update(model)
	}
	return err
}

func (a *adapterService) Update(adapter *entity.Adapter) error {
	return a.coll.Update(adapter)
}

func (a *adapterService) ChangeModel(name string, modelName string) error {
	adapter, err := a.FindByName(name)
	if err != nil {
		return err
	}
	if adapter.ModelName == modelName {
		return nil
	}
	ModelServ.Lock()
	defer ModelServ.Unlock()
	newModel, err := ModelServ.FindByName(modelName)
	if err != nil {
		return err
	}
	model, _ := ModelServ.FindByName(adapter.ModelName)
	adapter.ModelName = modelName
	if err = a.Update(adapter); err != nil {
		return err
	}
	defer a.ResetModel(*adapter)
	if model.Used > 0 {
		model.Used--
		if err = ModelServ.update(model); err != nil {
			return err
		}
	}
	newModel.Used++
	return ModelServ.update(newModel)
}

func (a *adapterService) ResetModel(adapters ...entity.Adapter) {
	for _, adapter := range adapters {
		log.Println("reset model on:", adapter.Name)
		mqtt.ResetModel(adapter.Name)
	}
}

func (a *adapterService) RenameModel(newName string, adapters ...entity.Adapter) error {
	for _, adapter := range adapters {
		adapter.ModelName = newName
		if err := a.coll.Update(&adapter); err != nil {
			return err
		}
	}
	return nil
}

func (a *adapterService) RenameModelFrom(oldName, newName string) error {
	adapters, err := a.FindByModelName(oldName)
	if err != nil {
		return err
	}
	return a.RenameModel(newName, adapters...)
}
