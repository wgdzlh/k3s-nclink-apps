package service

import (
	"context"
	"reflect"
	"strconv"
	"strings"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Service interface {
// 	Create(model interface{}) error
// 	IndexId() error
// 	Save(model interface{}) error
// 	FindById(id string) (interface{}, error)
// 	UpdateById(id string, in interface{}) (changed bool, err error)
// 	Rename(id, newId string) error
// 	DeleteById(id string) error
// }

type DummyService struct {
	mt   reflect.Type
	coll *mgm.Collection
}

func (d *DummyService) SetColl(model mgm.Model) {
	d.mt = reflect.TypeOf(model).Elem()
	d.coll = mgm.Coll(model)
}

func (d *DummyService) Create(model mgm.Model) error {
	err := d.coll.Create(model)
	if err != nil {
		return err
	}
	return d.IndexId()
}

func (d *DummyService) IndexId() error {
	ctx := context.Background()
	_, err := d.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func (d *DummyService) Save(model mgm.Model) error {
	return d.coll.Create(model)
}

func (d *DummyService) FindById(id string) (interface{}, error) {
	ret := reflect.New(d.mt).Interface().(mgm.Model)
	if err := d.coll.First(bson.M{"id": id}, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

func TransFilters(filters map[string]string) (bson.M, *options.FindOptions) {
	fil := bson.M{}
	opts := &options.FindOptions{}
	var skip, limit int64
	if _sort, ok := filters["_sort"]; ok {
		order := 1
		if _order, ok := filters["_order"]; ok && strings.ToUpper(_order) == "DESC" {
			order = -1
		}
		opts.SetSort(bson.D{{Key: _sort, Value: order}})
	}
	if _start, ok := filters["_start"]; ok && _start != "0" {
		if start, err := strconv.ParseInt(_start, 10, 64); err == nil {
			skip = start
		}
	}
	if _end, ok := filters["_end"]; ok && _end != "0" {
		if end, err := strconv.ParseInt(_end, 10, 64); err == nil && end > skip {
			limit = end - skip
		}
	}
	opts.SetSkip(skip)
	opts.SetLimit(limit)
	for k, v := range filters {
		if k[0] != '_' {
			fil[k] = v
		}
	}
	return fil, opts
}
