package service

import (
	"context"
	"k3s-nclink-apps/config-distribute/models/entity"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdapterService struct{}

func (m AdapterService) Create(adapter *entity.Adapter) error {
	coll := mgm.Coll(adapter)
	err := coll.Create(adapter)
	if err != nil {
		return err
	}
	_, err = coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

// Find adapter
func (u AdapterService) FindByName(name string) (*entity.Adapter, error) {
	ret := &entity.Adapter{}
	coll := mgm.Coll(ret)
	err := coll.First(bson.M{"name": name}, ret)
	return ret, err
}
