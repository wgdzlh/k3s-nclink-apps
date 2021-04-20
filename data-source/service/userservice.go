package service

import (
	"context"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (u UserService) Create(user *entity.User) error {
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

// Find user
func (u UserService) Find(user *entity.User) (*entity.User, error) {
	return u.FindByName(user.Name)
}

func (u UserService) FindByName(name string) (*entity.User, error) {
	ret := &entity.User{}
	coll := mgm.Coll(ret)
	err := coll.First(bson.M{"name": name}, ret)
	return ret, err
}

// Get JWT token
var TokenKey = []byte(utils.EnvVar("TOKEN_KEY", ""))

func (u UserService) GetJwtToken(user *entity.User) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
	})
	tokenString, err = token.SignedString(TokenKey)
	return
}
