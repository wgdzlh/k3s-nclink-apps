package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func RetKV(c *gin.Context, key string, value interface{}) {
	c.JSON(http.StatusOK, gin.H{
		key: value,
	})
}

func RetRaw(c *gin.Context, value interface{}) {
	c.JSON(http.StatusOK, value)
}

func Created(c *gin.Context, msg string) {
	c.JSON(http.StatusCreated, gin.H{
		"msg": msg,
	})
}

func CreatedRaw(c *gin.Context, value interface{}) {
	c.JSON(http.StatusCreated, value)
}

func BadRequest(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": msg})
}

func InternalError(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": msg})
}

func Unauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": msg})
}

func Forbidden(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": msg})
}
