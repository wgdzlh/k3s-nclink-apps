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

type DummyService struct {
	coll *mgm.Collection
}

func (d *DummyService) setColl(model mgm.Model) {
	d.coll = mgm.Coll(model)
}

func (d *DummyService) Create(model mgm.Model) error {
	if err := d.create(model); err != nil {
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

func (d *DummyService) findById(id string, ret mgm.Model) error {
	if err := d.coll.First(bson.M{"id": id}, ret); err != nil {
		return err
	}
	return nil
}

func (d *DummyService) findAll(ret interface{}) error {
	if err := d.coll.SimpleFind(ret, bson.M{}); err != nil {
		return err
	}
	return nil
}

func (d *DummyService) findPartial(filters map[string]string, ret interface{}) (interface{}, error) {
	fil, opts := transFilters(filters)
	if err := d.coll.SimpleFind(ret, fil, opts); err != nil {
		return nil, err
	}
	return fil, nil
}

func (d *DummyService) findWithFilter(filters map[string]string, ret interface{}) (int64, error) {
	fil, err := d.findPartial(filters, ret)
	if err != nil {
		return 0, err
	}
	return d.coll.CountDocuments(context.Background(), fil)
}

func (d *DummyService) create(model mgm.Model) error {
	return d.coll.Create(model)
}

func (d *DummyService) delete(model mgm.Model) error {
	return d.coll.Delete(model)
}

func (d *DummyService) update(model mgm.Model) error {
	return d.coll.Update(model)
}

func transFilters(filters map[string]string) (bson.M, *options.FindOptions) {
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
