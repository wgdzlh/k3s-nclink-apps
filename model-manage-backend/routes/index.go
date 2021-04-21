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

	router.POST("/login", authController.Login)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.AuthErrorHandler())
	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	authGroup.GET("/models", modelController.FetchAll)
	authGroup.POST("/models", modelController.New)
	authGroup.POST("/models/:id/:new-name", modelController.Dup)
	authGroup.PUT("/models/:id", modelController.Edit)
	authGroup.PUT("/models/:id/:new-name", modelController.Rename)
	authGroup.DELETE("/models/:id", modelController.Delete)
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setAuthRoute(router)
	return router
}
