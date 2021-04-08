package middlewares

import (
	"config-distribute/models/service"
	"config-distribute/utils"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// auth middleware
func AuthErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authentication")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": "Authentication header is missing.",
			})
			return
		}

		temp := strings.Split(authHeader, "Bearer")
		if len(temp) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "Wrong token."})
			return
		}

		tokenString := strings.TrimSpace(temp[1])
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			secretKey := utils.EnvVar("TOKEN_KEY", "")
			return []byte(secretKey), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			name := claims["name"].(string)
			userservice := service.Userservice{}
			user, err := userservice.FindByName(name)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User not found."})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token invalid."})
		}
	}
}

// func GlobalErrorHandler(c *gin.Context) {
// 	c.Next()
// 	if len(c.Errors) > 0 {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": c.Errors})
// 	}
// }
