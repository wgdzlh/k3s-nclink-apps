package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/model-manage-backend/mqtt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type adapterService struct {
	DummyService
}

var AdapterServ = &adapterService{}

func init() {
	AdapterServ.SetColl(&entity.Model{})
}

func (a *adapterService) Create(adapter *entity.Adapter) error {
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model, err := ModelServ.FindById(adapter.ModelId)
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
	return a.IndexId()
}

func (a *adapterService) Save(adapter *entity.Adapter, model *entity.Model) error {
	ModelServ.Lock()
	defer ModelServ.Unlock()
	if err := a.coll.Create(adapter); err != nil {
		return err
	}
	model.Used++
	if err := ModelServ.update(model); err != nil {
		return err
	}
	return nil
}

// Find adapter
func (a *adapterService) FindById(id string) (*entity.Adapter, error) {
	ret, err := a.DummyService.FindById(id)
	return ret.(*entity.Adapter), err
}

func (a *adapterService) FindAll() ([]entity.Adapter, int64, error) {
	ret := []entity.Adapter{}
	err := a.coll.SimpleFind(&ret, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	num, err := a.coll.EstimatedDocumentCount(context.Background())
	return ret, num, err
}

func (a *adapterService) FindByModelId(modelId string) ([]entity.Adapter, error) {
	ret := []entity.Adapter{}
	err := a.coll.SimpleFind(&ret, bson.M{"model_id": modelId})
	return ret, err
}

func (a *adapterService) delete(adapter *entity.Adapter) error {
	if err := a.coll.Delete(adapter); err != nil {
		return err
	}
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model, _ := ModelServ.FindById(adapter.ModelId)
	if model.Used > 0 {
		model.Used--
		if err := ModelServ.update(model); err != nil {
			return err
		}
	}
	return nil
}

func (a *adapterService) DeleteById(id string) error {
	adapter, err := a.FindById(id)
	if err != nil {
		return err
	}
	return a.delete(adapter)
}

func (a *adapterService) update(adapter *entity.Adapter) error {
	return a.coll.Update(adapter)
}

func (a *adapterService) changeModel(adapter *entity.Adapter, modelId string) error {
	if adapter.ModelId == modelId {
		return nil
	}
	ModelServ.Lock()
	defer ModelServ.Unlock()
	newModel, err := ModelServ.FindById(modelId)
	if err != nil {
		return err
	}
	model, _ := ModelServ.FindById(adapter.ModelId)
	adapter.ModelId = modelId
	if err = a.update(adapter); err != nil {
		return err
	}
	if model.Used > 0 {
		model.Used--
		if err = ModelServ.update(model); err != nil {
			return err
		}
	}
	newModel.Used++
	return ModelServ.update(newModel)
}

func (a *adapterService) UpdateById(id string, in *entity.Adapter) (changed bool, err error) {
	adapter, err := a.FindById(id)
	if err != nil {
		return
	}
	if adapter.DevId == in.DevId && adapter.ModelId == in.ModelId {
		return
	}
	adapter.DevId = in.DevId
	if err = a.changeModel(adapter, in.ModelId); err != nil {
		return
	}
	mqtt.ResetModel(adapter.Id)
	return true, nil
}

func (a *adapterService) Rename(id, newId string) error {
	if id == newId {
		return nil
	}
	adapter, err := a.FindById(id)
	if err != nil {
		return err
	}
	adapter.Id = newId
	return a.update(adapter)
}

func (a *adapterService) ResetModel(adapters ...entity.Adapter) {
	for _, adapter := range adapters {
		log.Println("reset model on:", adapter.Id)
		mqtt.ResetModel(adapter.Id)
	}
}

func (a *adapterService) RenameModel(newId string, adapters ...entity.Adapter) error {
	for _, adapter := range adapters {
		adapter.ModelId = newId
		if err := a.coll.Update(&adapter); err != nil {
			return err
		}
	}
	return nil
}

func (a *adapterService) RenameModelFrom(oldId, newId string) error {
	adapters, err := a.FindByModelId(oldId)
	if err != nil {
		return err
	}
	return a.RenameModel(newId, adapters...)
}
