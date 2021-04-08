package service

import (
	"config-distribute/models/entity"
	"context"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type Userservice struct{}

func (u Userservice) Create(user *entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	coll := mgm.Coll(user)
	err = coll.Create(user)
	if err != nil {
		return err
	}
	_, err = coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

// Find user
func (u Userservice) Find(user *entity.User) (*entity.User, error) {
	ret := &entity.User{}
	coll := mgm.Coll(ret)
	err := coll.First(bson.M{"name": user.Name}, ret)
	if err != nil {
		return nil, err
	}
	return ret, err
}

func (u Userservice) FindByName(name string) (*entity.User, error) {
	user := entity.NewUser(name, "")
	return u.Find(user)
}
