package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AdapterService struct{}

func (a AdapterService) Create(adapter *entity.Adapter) error {
	coll := mgm.Coll(adapter)
	err := coll.Create(adapter)
	if err != nil {
		return err
	}
	ctx := context.Background()
	num, err := coll.EstimatedDocumentCount(ctx)
	if num <= 1 {
		_, err = coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}
	return err
}

// Find adapter
func (a AdapterService) FindByName(name string) (*entity.Adapter, error) {
	ret := &entity.Adapter{}
	coll := mgm.Coll(ret)
	err := coll.First(bson.M{"name": name}, ret)
	return ret, err
}
