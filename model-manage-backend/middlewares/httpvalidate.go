package middlewares

import (
	"k3s-nclink-apps/data-source/service"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// auth middleware
func AuthErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Authorization header is missing.",
			})
			return
		}

		temp := strings.SplitN(authHeader, "Bearer", 2)
		if len(temp) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong token."})
			return
		}

		tokenString := strings.TrimSpace(temp[1])
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return service.UserServ.TokenKey, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			name := claims["name"].(string)
			user, err := service.UserServ.FindByName(name)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User not found."})
				return
			}
			if user.Access != service.UserServ.AccessType {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User access limited."})
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token invalid."})
		}
	}
}
