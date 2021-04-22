package controllers

import (
	"fmt"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"

	"github.com/gin-gonic/gin"
)

type AdapterController struct{}

type dummyAdapter struct {
	Name    string `json:"name"`
	DevId   string `json:"dev_id"`
	ModelId string `json:"model_id"`
}

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
	var da dummyAdapter
	err := c.ShouldBindJSON(&da)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	model, err := service.ModelServ.FindById(da.ModelId)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	realone := entity.NewAdapter(da.Name, da.DevId, model.Name)
	err = service.AdapterServ.Save(realone, model)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s created.", realone.Name)
	rest.CreatedRaw(c, gin.H{"id": realone.ID.Hex(), "msg": msg})
}

func (a AdapterController) Dup(c *gin.Context) {
	id := c.Param("id")
	newName := c.Param("new-name")
	if id == "" || newName == "" {
		rest.BadRequest(c, "param 'name' or 'dup' not set.")
		return
	}
	adapter, err := service.AdapterServ.FindById(id)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	newAdapter := entity.NewAdapter(newName, adapter.DevId, adapter.ModelName)
	err = service.AdapterServ.Create(newAdapter)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s duplicated as %s.", adapter.Name, newName)
	rest.Created(c, msg)
}

func (a AdapterController) Edit(c *gin.Context) {
	var adapter entity.Adapter
	err := c.ShouldBindJSON(&adapter)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	if adapter.ID.IsZero() {
		rest.BadRequest(c, "'id' not set in JSON.")
		return
	}
	id := c.Param("id")
	if adapter.ID.Hex() != id {
		rest.BadRequest(c, "'id' not match to param.")
		return
	}
	changed, err := service.AdapterServ.UpdateById(id, &adapter)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	res := "unchanged"
	if changed {
		res = "updated"
	}
	msg := fmt.Sprintf("adapter %s %s.", adapter.Name, res)
	rest.RetRaw(c, gin.H{"id": id, "msg": msg})
}

func (a AdapterController) Rename(c *gin.Context) {
	id := c.Param("id")
	newName := c.Query("new-name")
	if id == "" || newName == "" {
		rest.BadRequest(c, "param 'id' or 'new-name' not set.")
		return
	}
	if err := service.AdapterServ.Rename(id, newName); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("adapter %s renamed to %s.", id, newName)
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
