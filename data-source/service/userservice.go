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

type userService struct {
	TokenKey   []byte
	AccessType string
	coll       *mgm.Collection
}

var UserServ = &userService{
	TokenKey:   []byte(utils.EnvVar("TOKEN_KEY", "")), // JWT token key
	AccessType: utils.EnvVar("USER_ACCESS_TYPE", "ro"),
	coll:       mgm.Coll(&entity.User{}),
}

func (u *userService) Create(user *entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	err = u.coll.Create(user)
	if err != nil {
		return err
	}
	ctx := context.Background()
	num, err := u.coll.EstimatedDocumentCount(ctx)
	if num <= 1 {
		_, err = u.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "name", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	}
	return err
}

// Find user
func (u *userService) FindByName(name string) (*entity.User, error) {
	ret := &entity.User{}
	err := u.coll.First(bson.M{"name": name}, ret)
	return ret, err
}

func (u *userService) GetJwtToken(user *entity.User) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
	})
	tokenString, err = token.SignedString(u.TokenKey)
	return
}
