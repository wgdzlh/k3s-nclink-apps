package middlewares

import (
	"k3s-nclink-apps/data-source/service"
	"k3s-nclink-apps/model-manage-backend/rest"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// auth middleware
func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			rest.Unauthorized(c, "Authorization header is missing.")
			return
		}

		temp := strings.SplitN(authHeader, "Bearer", 2)
		if len(temp) < 2 {
			rest.Unauthorized(c, "Wrong token.")
			return
		}

		tokenString := strings.TrimSpace(temp[1])
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return service.UserServ.TokenKey, nil
		})
		if err != nil {
			rest.Unauthorized(c, err.Error())
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			name := claims["name"].(string)
			user, err := service.UserServ.FindByName(name)
			if err != nil {
				rest.Forbidden(c, "User not found.")
				return
			}
			if user.Access != service.UserServ.AccessType {
				rest.Forbidden(c, "User access limited.")
				return
			}
			// c.Set("user", user)
			c.Next()
		} else {
			rest.Forbidden(c, "Token invalid.")
		}
	}
}
