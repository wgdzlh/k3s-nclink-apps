package service

import (
	"context"
	"k3s-nclink-apps/config-distribute/models/entity"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModelService struct{}

func (m ModelService) Create(model *entity.Model) error {
	coll := mgm.Coll(model)
	err := coll.Create(model)
	if err != nil {
		return err
	}
	_, err = coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "hostname", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

// Find model
func (u ModelService) FindByHost(hostname string) (*entity.Model, error) {
	ret := &entity.Model{}
	coll := mgm.Coll(ret)
	err := coll.First(bson.M{"hostname": hostname}, ret)
	return ret, err
}
