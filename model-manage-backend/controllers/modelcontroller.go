package controllers

import (
	"fmt"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"

	"github.com/gin-gonic/gin"
)

type ModelController struct{}

func (m ModelController) Fetch(c *gin.Context) {
	var ret []entity.Model
	var num int64
	queries := c.Request.URL.Query()
	if len(queries) == 0 {
		all, err := service.ModelServ.FindAll()
		if err != nil {
			rest.InternalError(c, err.Error())
			return
		}
		ret = all
		num = int64(len(ret))
	} else {
		filter := make(map[string]string, len(queries))
		for key, value := range queries {
			filter[key] = value[0]
		}
		patial, _num, err := service.ModelServ.FindWithFilter(filter)
		if err != nil {
			rest.InternalError(c, err.Error())
			return
		}
		ret = patial
		num = _num
	}
	c.Header("X-Total-Count", fmt.Sprint(num))
	// c.Header("Access-Control-Expose-Headers", "X-Total-Count")
	rest.RetRaw(c, ret)
}

func (m ModelController) One(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		rest.BadRequest(c, "param 'id' not set.")
		return
	}
	model, err := service.ModelServ.FindById(id)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	rest.RetRaw(c, model)
}

func (m ModelController) New(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	model.Used = 0
	err = service.ModelServ.Save(&model)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s created.", model.Id)
	rest.CreatedRaw(c, gin.H{"id": model.Id, "msg": msg})
}

func (m ModelController) Dup(c *gin.Context) {
	id := c.Param("id")
	newId := c.Query("new-id")
	if id == "" || newId == "" {
		rest.BadRequest(c, "param 'id' or 'new-id' not set.")
		return
	}
	model, err := service.ModelServ.FindById(id)
	if err != nil {
		rest.BadRequest(c, "model not found.")
		return
	}
	newModel := entity.NewModel(newId, model.Def)
	err = service.ModelServ.Save(newModel)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s duplicated as %s.", model.Id, newId)
	rest.CreatedRaw(c, gin.H{"id": newId, "msg": msg})
}

func (m ModelController) Edit(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	if model.Id == "" {
		rest.BadRequest(c, "'id' not set in JSON.")
		return
	}
	id := c.Param("id")
	if model.Id != id {
		rest.BadRequest(c, "'id' not match to param.")
		return
	}
	changed, err := service.ModelServ.UpdateById(id, &model)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	res := "unchanged"
	if changed {
		res = "updated"
	}
	msg := fmt.Sprintf("model %s %s.", model.Id, res)
	rest.RetRaw(c, gin.H{"id": id, "msg": msg})
}

func (m ModelController) Rename(c *gin.Context) {
	id := c.Param("id")
	newId := c.Query("new-id")
	if id == "" || newId == "" {
		rest.BadRequest(c, "param 'id' or 'new-id' not set.")
		return
	}
	if err := service.ModelServ.Rename(id, newId); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s renamed to %s.", id, newId)
	rest.OK(c, msg)
}

func (m ModelController) Delete(c *gin.Context) {
	id := c.Param("id")
	err := service.ModelServ.DeleteById(id)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s deleted.", id)
	rest.OK(c, msg)
}
