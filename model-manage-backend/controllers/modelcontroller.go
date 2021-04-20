package controllers

import (
	"k3s-nclink-apps/data-source/entity"
	"k3s-nclink-apps/data-source/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ModelController struct {
	modelService service.ModelService
}

func (m ModelController) FetchAll(c *gin.Context) {
	ret, err := m.modelService.FindAll()
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
	err = m.modelService.Create(newModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "model created.",
	})
}

func (m ModelController) Dup(c *gin.Context) {
	name := c.Query("name")
	newName := c.Query("new-name")
	if name == "" || newName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "query 'name' or 'new-name' not set."})
		return
	}
	model, err := m.modelService.FindByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newModel := entity.NewModel(newName, model.Def)
	err = m.modelService.Create(newModel)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "model duplicated.",
	})
}

func (m ModelController) Edit(c *gin.Context) {
	var model entity.Model
	err := c.ShouldBindJSON(&model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = m.modelService.UpdateByName(model.Name, model.Def)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "model updated.",
	})
}

func (m ModelController) Remove(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "query 'name' not set."})
		return
	}
	err := m.modelService.DeleteByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "model deleted.",
	})
}
