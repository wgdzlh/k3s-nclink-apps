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
