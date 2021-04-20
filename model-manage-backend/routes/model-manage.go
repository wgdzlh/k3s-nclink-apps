package routes

import (
	"context"
	pb "k3s-nclink-apps/configmodel"

	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
)

// model manage server
type modelManageServer struct {
	pb.UnimplementedModelManageServer
}

func (s *modelManageServer) SaveModel(ctx context.Context, model *pb.Model) (*pb.OpRet, error) {
	log.Infoln("Save model:", model)
	return nil, nil
}

func (s *modelManageServer) DeleteModel(context.Context, *pb.Model) (*pb.OpRet, error) {
	return nil, nil

}

func (s *modelManageServer) UpdateModel(context.Context, *pb.Model) (*pb.OpRet, error) {
	return nil, nil
}

func (s *modelManageServer) FindModels(*pb.Filter, pb.ModelManage_FindModelsServer) error {
	return nil
}

func RegisterServices(server *grpc.Server) {
	pb.RegisterModelManageServer(server, &modelManageServer{})
}
