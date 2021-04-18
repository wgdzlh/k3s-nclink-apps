package routes

import (
	"context"
	"k3s-nclink-apps/config-distribute/controllers"
	"k3s-nclink-apps/config-distribute/models/entity"
	pb "k3s-nclink-apps/configmodel"
	"k3s-nclink-apps/utils/conv"

	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
)

// auth server
type authServer struct {
	pb.UnimplementedAuthenticationServer
	authController controllers.AuthController
}

func (s *authServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	name := in.GetName()
	pass := in.GetPassword()
	log.Infof("Received login request from: %v", name)
	loginInfo := &entity.User{Name: name, Password: pass}
	token, err := s.authController.Login(loginInfo)
	return &pb.LoginReply{Token: token}, err
}

func (s *authServer) Ping(context.Context, *emptypb.Empty) (*pb.Pong, error) {
	return &pb.Pong{Message: "pong"}, nil
}

// model server
type modelDistServer struct {
	pb.UnimplementedModelDistServer
	modelcontroller controllers.DistController
}

func (s *modelDistServer) GetModel(ctx context.Context, in *pb.ModelRequest) (*pb.ModelReply, error) {
	hostname := in.GetHostname()
	log.Infof("Received model fetch request from: %v", hostname)
	model, devId, err := s.modelcontroller.Fetch(hostname)
	if err != nil {
		return nil, err
	}
	outModel, err := conv.DbModelToWireModel(model)
	if err != nil {
		return nil, err
	}
	return &pb.ModelReply{Model: outModel, DevId: devId}, nil
}

func RegisterServices(server *grpc.Server) {
	pb.RegisterAuthenticationServer(server, &authServer{})
	pb.RegisterModelDistServer(server, &modelDistServer{})
}
