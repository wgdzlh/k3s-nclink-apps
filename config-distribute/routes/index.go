package routes

import (
	"k3s-nclink-apps/config-distribute/controllers"
	"k3s-nclink-apps/config-distribute/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.AuthErrorHandler())
	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setAuthRoute(router)
	return router
}
