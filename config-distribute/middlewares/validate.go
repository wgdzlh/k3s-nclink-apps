package middlewares

import (
	"context"
	"k3s-nclink-apps/config-distribute/models/service"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// auth middleware
var (
	userservice        = service.UserService{}
	errMissingMetadata = status.Error(codes.InvalidArgument, "missing metadata")
	errMissingToken    = status.Error(codes.InvalidArgument, "missing token")
	errInvalidToken    = status.Error(codes.Unauthenticated, "invalid token")
)

// ensureValid ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func EnsureValid(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	if info.FullMethod != "/configmodel.Authentication/Login" {
		if err := authorize(ctx); err != nil {
			return nil, err
		}
	}

	h, err := handler(ctx, req)

	log.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

func authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errMissingMetadata
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return errMissingToken
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	if !validateToken(tokenString) {
		return errInvalidToken
	}
	return nil
}

func validateToken(tokenString string) bool {

	// log.Infoln(tokenKey)

	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return service.TokenKey, nil
	})
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		name := claims["name"].(string)
		if _, err = userservice.FindByName(name); err == nil {
			return true
		}
	}
	return false
}

// func AuthErrorHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.Request.Header.Get("Authorization")
// 		if len(authHeader) == 0 {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"error": "Authorization header is missing.",
// 			})
// 			return
// 		}

// 		temp := strings.Split(authHeader, "Bearer")
// 		if len(temp) < 2 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Wrong token."})
// 			return
// 		}

// 		tokenString := strings.TrimSpace(temp[1])
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			secretKey := utils.EnvVar("TOKEN_KEY", "")
// 			return []byte(secretKey), nil
// 		})
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			name := claims["name"].(string)
// 			userservice := service.UserService{}
// 			user, err := userservice.FindByName(name)
// 			if err != nil {
// 				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "User not found."})
// 				return
// 			}
// 			c.Set("user", user)
// 			c.Next()
// 		} else {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token invalid."})
// 		}
// 	}
// }
