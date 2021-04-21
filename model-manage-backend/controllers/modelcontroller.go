package controllers

import (
	"fmt"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"

	"github.com/gin-gonic/gin"
)

type ModelController struct{}

func (m ModelController) FetchAll(c *gin.Context) {
	ret, num, err := service.ModelServ.FindAll()
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	c.Header("X-Total-Count", fmt.Sprint(num))
	rest.Ret(c, "models", ret)
}

func (m ModelController) New(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	err = service.ModelServ.Save(model.Name, model.Def)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s created.", model.Name)
	rest.Created(c, msg)
}

func (m ModelController) Dup(c *gin.Context) {
	id := c.Param("id")
	newName := c.Param("new-name")
	if id == "" || newName == "" {
		rest.BadRequest(c, "param 'name' or 'dup' not set.")
		return
	}
	model, err := service.ModelServ.FindById(id)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	err = service.ModelServ.Save(newName, model.Def)
	if err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s duplicated as %s.", model.Name, newName)
	rest.Created(c, msg)
}

func (m ModelController) Edit(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		rest.BadRequest(c, err.Error())
		return
	}
	if model.ID.IsZero() {
		rest.BadRequest(c, "'id' not set in JSON.")
		return
	}
	id := c.Param("id")
	if model.ID.Hex() != id {
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
	msg := fmt.Sprintf("model %s %s.", model.Name, res)
	rest.OK(c, msg)
}

func (m ModelController) Rename(c *gin.Context) {
	id := c.Param("id")
	newName := c.Param("new-name")
	if id == "" || newName == "" {
		rest.BadRequest(c, "param 'id' or 'new-name' not set.")
		return
	}
	if err := service.ModelServ.Rename(id, newName); err != nil {
		rest.InternalError(c, err.Error())
		return
	}
	msg := fmt.Sprintf("model %s renamed to %s.", id, newName)
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
