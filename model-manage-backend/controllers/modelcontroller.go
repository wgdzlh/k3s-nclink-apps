package controllers

import (
	"fmt"
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModelController struct{}

func (m ModelController) FetchAll(c *gin.Context) {
	ret, err := service.ModelServ.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"models": ret,
	})
}

func (m ModelController) New(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newModel := entity.NewModel(model.Name, model.Def)
	err = service.ModelServ.Create(newModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	retMsg := fmt.Sprintf("model %s created.", model.Name)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    retMsg,
	})
}

func (m ModelController) Dup(c *gin.Context) {
	name := c.Query("name")
	newName := c.Query("new-name")
	if name == "" || newName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "query 'name' or 'new-name' not set."})
		return
	}
	model, err := service.ModelServ.FindByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newModel := entity.NewModel(newName, model.Def)
	err = service.ModelServ.Create(newModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	retMsg := fmt.Sprintf("model %s duplicated as %s.", name, newName)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    retMsg,
	})
}

func (m ModelController) Edit(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	changed, err := service.ModelServ.UpdateByName(model.Name, model.Def)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	opRet := "unchanged"
	if changed {
		opRet = "updated"
		
	}
	retMsg := fmt.Sprintf("model %s %s.", model.Name, opRet)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    retMsg,
	})
}

func (m ModelController) Remove(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "query 'name' not set."})
		return
	}
	err := service.ModelServ.DeleteByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	retMsg := fmt.Sprintf("model %s deleted.", name)
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    retMsg,
	})
}
