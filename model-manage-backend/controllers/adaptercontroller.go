package controllers

import (
	"fmt"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"

	"github.com/gin-gonic/gin"
)

type AdapterController struct{}

func (a AdapterController) FetchAll(c *gin.Context) {
	ret, num, err := service.AdapterServ.FindAll()
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	c.Header("X-Total-Count", fmt.Sprint(num))
	rest.RetRaw(c, ret)
}

func (a AdapterController) One(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.BadRequest(c, "param 'id' not set.")
		return
	}
	adapter, err := service.AdapterServ.FindById(id)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	rest.RetRaw(c, adapter)
}

func (a AdapterController) New(c *gin.Context) {
	var da entity.Adapter
	err := c.ShouldBindJSON(&da)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	model, err := service.ModelServ.FindById(da.ModelId)
	if err != nil {
		rest.BadRequest(c, "model not found.")
		return
	}
	newAdapter := entity.NewAdapter(da.Id, da.DevId, model.Id)
	err = service.AdapterServ.Save(newAdapter, model)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s created.", newAdapter.Id)
	rest.CreatedRaw(c, gin.H{"id": newAdapter.Id, "msg": msg})
}

func (a AdapterController) Dup(c *gin.Context) {
	id := c.Param("id")
	newId := c.Query("new-id")
	if id == "" || newId == "" {
		rest.BadRequest(c, "param 'id' or 'new-id' not set.")
		return
	}
	adapter, err := service.AdapterServ.FindById(id)
	if err != nil {
		rest.BadRequest(c, "adapter not found.")
		return
	}
	model, err := service.ModelServ.FindById(adapter.ModelId)
	if err != nil {
		rest.BadRequest(c, "model not found.")
		return
	}
	newAdapter := entity.NewAdapter(newId, adapter.DevId, adapter.ModelId)
	err = service.AdapterServ.Save(newAdapter, model)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s duplicated as %s.", adapter.Id, newId)
	rest.CreatedRaw(c, gin.H{"id": newId, "msg": msg})
}

func (a AdapterController) Edit(c *gin.Context) {
	var da entity.Adapter
	err := c.ShouldBindJSON(&da)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	if da.Id == "" {
		rest.BadRequest(c, "'id' not set in JSON.")
		return
	}
	id := c.Param("id")
	if da.Id != id {
		rest.BadRequest(c, "'id' not match to param.")
		return
	}
	model, err := service.ModelServ.FindById(da.ModelId)
	if err != nil {
		rest.BadRequest(c, "model not found.")
		return
	}
	newAdapter := entity.NewAdapter(da.Id, da.DevId, model.Id)
	changed, err := service.AdapterServ.UpdateById(id, newAdapter)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	res := "unchanged"
	if changed {
		res = "updated"
	}
	msg := fmt.Sprintf("adapter %s %s.", da.Id, res)
	rest.RetRaw(c, gin.H{"id": id, "msg": msg})
}

func (a AdapterController) Rename(c *gin.Context) {
	id := c.Param("id")
	newId := c.Query("new-id")
	if id == "" || newId == "" {
		rest.BadRequest(c, "param 'id' or 'new-id' not set.")
		return
	}
	if err := service.AdapterServ.Rename(id, newId); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s renamed to %s.", id, newId)
	rest.OK(c, msg)
}

func (a AdapterController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := service.AdapterServ.DeleteById(id)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s deleted.", id)
	rest.OK(c, msg)
}
