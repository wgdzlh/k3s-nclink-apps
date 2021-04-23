package routes

import (
	"k3s-nclink-apps/model-manage-backend/controllers"
	"k3s-nclink-apps/model-manage-backend/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := controllers.AuthController{}
	modelController := controllers.ModelController{}
	adapterController := controllers.AdapterController{}

	router.POST("/login", authController.Login)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.AuthChecker())
	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	authGroup.GET("/models", modelController.Fetch)
	authGroup.GET("/models/:id", modelController.One)
	authGroup.POST("/models", modelController.New)
	authGroup.POST("/models/:id", modelController.Dup)
	authGroup.PUT("/models/:id", modelController.Edit)
	authGroup.PUT("/models/:id/rename", modelController.Rename)
	authGroup.DELETE("/models/:id", modelController.Delete)

	authGroup.GET("/adapters", adapterController.FetchAll)
	authGroup.GET("/adapters/:id", adapterController.One)
	authGroup.POST("/adapters", adapterController.New)
	authGroup.POST("/adapters/:id", adapterController.Dup)
	authGroup.PUT("/adapters/:id", adapterController.Edit)
	authGroup.PUT("/adapters/:id/rename", adapterController.Rename)
	authGroup.DELETE("/adapters/:id", adapterController.Delete)
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setAuthRoute(router)
	return router
}
