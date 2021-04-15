package routes

import (
	"context"
	"k3s-nclink-apps/config-distribute/controllers"
	"k3s-nclink-apps/config-distribute/models/entity"
	pb "k3s-nclink-apps/configmodel"

	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
)

// func setAuthRoute(router *gin.Engine) {
// 	authController := new(controllers.AuthController)
// 	router.POST("/login", authController.Login)

// 	authGroup := router.Group("/")
// 	authGroup.Use(middlewares.AuthErrorHandler())
// 	authGroup.GET("/ping", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "pong"})
// 	})
// }

// func InitRoute() *gin.Engine {
// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	router.Use(gin.Recovery())

// 	setAuthRoute(router)
// 	return router
// }

type authServer struct {
	pb.UnimplementedAuthenticationServer
	authController controllers.AuthController
}

func (s *authServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Infof("Received login request from: %v", in.GetName())
	loginInfo := &entity.User{Name: in.GetName(), Password: in.GetPassword()}
	token, _ := s.authController.Login(loginInfo)
	return &pb.LoginReply{Token: token}, nil
}

func (s *authServer) Ping(context.Context, *emptypb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Message: "pong"}, nil
}

func RegisterServices(server *grpc.Server) {
	auth := &authServer{}
	pb.RegisterAuthenticationServer(server, auth)
}
