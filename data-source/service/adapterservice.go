package service

import (
	"k3s-nclink-apps/model-manage-backend/mqtt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type adapterService struct {
	mongoService
}

var AdapterServ = &adapterService{}

func init() {
	AdapterServ.setColl(&Adapter{})
}

// Adapter hooks
func (d *Adapter) Deleted(result *mongo.DeleteResult) error {
	log.Printf("deleted adapter %s.\n", d.Id)
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model := &Model{}
	ModelServ.FindById(d.ModelId, model)
	if model.Used > 0 {
		model.Used--
		if err := ModelServ.update(model); err != nil {
			return err
		}
	}
	return nil
}

func (a *adapterService) New() interface{} {
	return &Adapter{}
}

func (a *adapterService) Dup(id string, in interface{}) interface{} {
	adapter := in.(*Adapter)
	return &Adapter{Id: id, DevId: adapter.DevId, ModelId: adapter.ModelId}
}

func (a *adapterService) IdOf(in interface{}) string {
	return in.(*Adapter).Id
}

func (a *adapterService) Slice() interface{} {
	return &[]Adapter{}
}

func (a *adapterService) LenOf(in interface{}) int64 {
	return int64(len(*in.(*[]Adapter)))
}

func (a *adapterService) Create(in interface{}) error {
	adapter := in.(*Adapter)
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model := &Model{}
	if err := ModelServ.FindById(adapter.ModelId, model); err != nil {
		return err
	}
	if err := a.create(adapter); err != nil {
		return err
	}
	model.Used++
	if err := ModelServ.update(model); err != nil {
		return err
	}
	return a.IndexId()
}

func (a *adapterService) Save(in interface{}) error {
	adapter := in.(*Adapter)
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model := &Model{}
	if err := ModelServ.FindById(adapter.ModelId, model); err != nil {
		return err
	}
	if err := a.create(adapter); err != nil {
		return err
	}
	model.Used++
	if err := ModelServ.update(model); err != nil {
		return err
	}
	return nil
}

func (a *adapterService) FindByModelId(modelId string) ([]Adapter, error) {
	ret := []Adapter{}
	_, err := a.FindPartial(map[string]string{"model_id": modelId}, &ret)
	return ret, err
}

func (a *adapterService) DeleteById(id string) error {
	adapter := &Adapter{}
	if err := a.FindById(id, adapter); err != nil {
		return err
	}
	return a.delete(adapter)
}

func (a *adapterService) changeModel(adapter *Adapter, modelId string) error {
	if adapter.ModelId == modelId {
		return nil
	}
	ModelServ.Lock()
	defer ModelServ.Unlock()
	newModel, model := &Model{}, &Model{}
	if err := ModelServ.FindById(modelId, newModel); err != nil {
		return err
	}
	ModelServ.FindById(adapter.ModelId, model)
	adapter.ModelId = modelId
	if err := a.update(adapter); err != nil {
		return err
	}
	if model.Used > 0 {
		model.Used--
		if err := ModelServ.update(model); err != nil {
			return err
		}
	}
	newModel.Used++
	return ModelServ.update(newModel)
}

func (a *adapterService) UpdateById(id string, in interface{}) (changed bool, err error) {
	devId, modelId := in.(*Adapter).DevId, in.(*Adapter).ModelId
	adapter := &Adapter{}
	if err = a.FindById(id, adapter); err != nil {
		return
	}
	if adapter.DevId == devId && adapter.ModelId == modelId {
		return
	}
	adapter.DevId = devId
	if err = a.changeModel(adapter, modelId); err != nil {
		return
	}
	mqtt.ResetModel(adapter.Id)
	return true, nil
}

func (a *adapterService) Rename(id, newId string) error {
	if id == newId {
		return nil
	}
	adapter := &Adapter{}
	if err := a.FindById(id, adapter); err != nil {
		return err
	}
	adapter.Id = newId
	return a.update(adapter)
}

func (a *adapterService) RenameModel(newId string, adapters ...Adapter) error {
	for _, adapter := range adapters {
		adapter.ModelId = newId
		if err := a.update(&adapter); err != nil {
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

// MQTT stuff
func (a *adapterService) ResetModel(adapters ...Adapter) {
	for _, adapter := range adapters {
		log.Println("reset model on:", adapter.Id)
		mqtt.ResetModel(adapter.Id)
	}
}
