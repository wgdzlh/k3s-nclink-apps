package controllers

import (
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dummyController struct {
	serv service.Service
}

func (d *dummyController) Fetch(c *gin.Context) {
	ret := d.serv.Slice()
	var num int64
	queries := c.Request.URL.Query()
	if len(queries) == 0 {
		err := d.serv.FindAll(ret)
		if err != nil {
			rest.InternalError(c, err.Error())
			return
		}
		num = d.serv.LenOf(ret)
	} else {
		filter := make(map[string]string, len(queries))
		for key, value := range queries {
			filter[key] = value[0]
		}
		_num, err := d.serv.FindWithFilter(filter, ret)
		if err != nil {
			rest.InternalError(c, err.Error())
			return
		}
		num = _num
	}
	c.Header("X-Total-Count", strconv.FormatInt(num, 10))
	// c.Header("Access-Control-Expose-Headers", "X-Total-Count")
	rest.RetRaw(c, ret)
}

func (d *dummyController) One(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.BadRequest(c, "param 'id' not set.")
		return
	}
	model := d.serv.New()
	if err := d.serv.FindById(id, model); err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	rest.RetRaw(c, model)
}

func (d *dummyController) Save(c *gin.Context) {
	model := d.serv.New()
	if err := c.ShouldBindJSON(model); err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	if err := d.serv.Save(model); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	id := d.serv.IdOf(model)
	rest.CreatedRaw(c, gin.H{"id": id, "msg": "created."})
}

func (d *dummyController) Copy(c *gin.Context) {
	id := c.Param("id")
	newId := c.Query("new-id")
	if id == "" || newId == "" {
		rest.BadRequest(c, "param 'id' or query 'new-id' not set.")
		return
	}
	if id == newId {
		rest.BadRequest(c, "duplicate id not alowed.")
		return
	}
	model := d.serv.New()
	if err := d.serv.FindById(id, model); err != nil {
		rest.BadRequest(c, id+" not found.")
		return
	}
	newModel := d.serv.Dup(newId, model)
	if err := d.serv.Save(newModel); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	rest.CreatedRaw(c, gin.H{"id": newId, "msg": id + "duplicated."})
}

func (d *dummyController) Edit(c *gin.Context) {
	model := d.serv.New()
	if err := c.ShouldBindJSON(model); err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	modelId := d.serv.IdOf(model)
	if modelId == "" {
		rest.BadRequest(c, "'id' not set in JSON.")
		return
	}
	id := c.Param("id")
	if modelId != id {
		rest.BadRequest(c, "'id' not match to param.")
		return
	}
	changed, err := d.serv.UpdateById(id, model)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := "unchanged."
	if changed {
		msg = "updated."
	}
	rest.RetRaw(c, gin.H{"id": id, "msg": msg})
}

func (d *dummyController) Rename(c *gin.Context) {
	id := c.Param("id")
	newId := c.Query("new-id")
	if id == "" || newId == "" {
		rest.BadRequest(c, "param 'id' or query 'new-id' not set.")
		return
	}
	if err := d.serv.Rename(id, newId); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	rest.RetRaw(c, gin.H{"id": newId, "msg": id + " renamed."})
}

func (d *dummyController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := d.serv.DeleteById(id)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	rest.RetRaw(c, gin.H{"id": id, "msg": "deleted."})
}
