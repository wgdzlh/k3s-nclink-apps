package service

import (
	"context"
	"strconv"
	"strings"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoService struct {
	coll *mgm.Collection
}

type myModel = mgm.Model

func (d *mongoService) Create(model interface{}) error {
	if err := d.create(model.(myModel)); err != nil {
		return err
	}
	return d.IndexId()
}

func (d *mongoService) Save(model interface{}) error {
	return d.create(model.(myModel))
}

func (d *mongoService) IndexId() error {
	ctx := context.Background()
	_, err := d.coll.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func (d *mongoService) FindById(id string, ret interface{}) error {
	return d.coll.First(bson.M{"id": id}, ret.(myModel))
}

func (d *mongoService) FindAll(ret interface{}) error {
	return d.coll.SimpleFind(ret, bson.M{})
}

func (d *mongoService) FindPartial(filters map[string]string, ret interface{}) (interface{}, error) {
	fil, opts := d.transFilters(filters)
	if err := d.coll.SimpleFind(ret, fil, opts); err != nil {
		return nil, err
	}
	return fil, nil
}

func (d *mongoService) FindWithFilter(filters map[string]string, ret interface{}) (int64, error) {
	fil, err := d.FindPartial(filters, ret)
	if err != nil {
		return 0, err
	}
	return d.coll.CountDocuments(context.Background(), fil)
}

func (d *mongoService) Delete(model interface{}) error {
	return d.delete(model.(myModel))
}

func (d *mongoService) Update(model interface{}) error {
	return d.update(model.(myModel))
}

func (d *mongoService) setColl(model myModel) {
	d.coll = mgm.Coll(model)
}

func (d *mongoService) create(model myModel) error {
	return d.coll.Create(model)
}

func (d *mongoService) delete(model myModel) error {
	return d.coll.Delete(model)
}

func (d *mongoService) update(model myModel) error {
	return d.coll.Update(model)
}

func (d *mongoService) transFilters(filters map[string]string) (bson.M, *options.FindOptions) {
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
