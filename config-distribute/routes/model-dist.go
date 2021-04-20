package routes

import (
	"context"
	"k3s-nclink-apps/config-distribute/controllers"
	pb "k3s-nclink-apps/configmodel"
	com "k3s-nclink-apps/grpc-common/controllers"
	"k3s-nclink-apps/utils/conv"

	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
)

// auth server
type authServer struct {
	pb.UnimplementedAuthenticationServer
	authController com.AuthController
}

func (s *authServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	name := in.GetName()
	pass := in.GetPassword()
	log.Infoln("Received login request from:", name)
	token, err := s.authController.Login(name, pass)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
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
	log.Infoln("Received model fetch request from:", hostname)
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
