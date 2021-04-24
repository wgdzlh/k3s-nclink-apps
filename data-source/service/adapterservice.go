package service

import (
	"k3s-nclink-apps/model-manage-backend/mqtt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type adapterService struct {
	DummyService
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
	model, _ := ModelServ.FindById(d.ModelId)
	if model != nil && model.Used > 0 {
		model.Used--
		if err := ModelServ.update(model); err != nil {
			return err
		}
	}
	return nil
}

// Adapter CRUDs
func (a *adapterService) Create(adapter *Adapter) error {
	ModelServ.Lock()
	defer ModelServ.Unlock()
	model, err := ModelServ.FindById(adapter.ModelId)
	if err != nil {
		return err
	}
	if err = a.create(adapter); err != nil {
		return err
	}
	model.Used++
	if err = ModelServ.update(model); err != nil {
		return err
	}
	return a.IndexId()
}

func (a *adapterService) Save(adapter *Adapter, model *Model) error {
	ModelServ.Lock()
	defer ModelServ.Unlock()
	if err := a.create(adapter); err != nil {
		return err
	}
	model.Used++
	if err := ModelServ.update(model); err != nil {
		return err
	}
	return nil
}

func (a *adapterService) FindById(id string) (*Adapter, error) {
	ret := &Adapter{}
	if err := a.findById(id, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (a *adapterService) FindAll() ([]Adapter, error) {
	ret := []Adapter{}
	if err := a.findAll(&ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func (a *adapterService) FindWithFilter(filters map[string]string) ([]Adapter, int64, error) {
	ret := []Adapter{}
	num, err := a.findWithFilter(filters, &ret)
	if err != nil {
		return nil, 0, err
	}
	return ret, num, nil
}

func (a *adapterService) FindByModelId(modelId string) ([]Adapter, error) {
	ret := []Adapter{}
	_, err := a.findPartial(map[string]string{"model_id": modelId}, &ret)
	return ret, err
}

func (a *adapterService) DeleteById(id string) error {
	adapter, err := a.FindById(id)
	if err != nil {
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
	newModel, err := ModelServ.FindById(modelId)
	if err != nil {
		return err
	}
	model, _ := ModelServ.FindById(adapter.ModelId)
	adapter.ModelId = modelId
	if err = a.update(adapter); err != nil {
		return err
	}
	if model != nil && model.Used > 0 {
		model.Used--
		if err = ModelServ.update(model); err != nil {
			return err
		}
	}
	newModel.Used++
	return ModelServ.update(newModel)
}

func (a *adapterService) UpdateById(id string, in *Adapter) (changed bool, err error) {
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
