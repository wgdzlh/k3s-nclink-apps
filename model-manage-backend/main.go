package main

import (
	"k3s-nclink-apps/grpc-common/middlewares"
	"k3s-nclink-apps/model-manage-backend/routes"
	"k3s-nclink-apps/utils"
	"net"

	"google.golang.org/grpc"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func grpcServ() {
	grpcHost := utils.EnvVar("GRPC_SERVER_HOST", "localhost")
	grpcPort := utils.EnvVar("GRPC_SERVER_PORT", "9000")

	stage := utils.GetEnvOrExit("DEV_STAGE")

	grpcAddr := grpcHost + ":" + grpcPort
	lis, err := net.Listen("tcp4", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(middlewares.EnsureValid),
	}

	server := grpc.NewServer(opts...)
	routes.RegisterServices(server)
	if stage == "debug" {
		reflection.Register(server)
	}
	log.Infoln("GRPC start serving on:", grpcAddr)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("GRPC failed to serve: %v", err)
	}
}

func main() {
	// userservice := service.UserService{}
	// user := entity.NewUser("admin", "123456")
	// err := userservice.Create(user)
	// if err != nil {
	// 	log.Println("Error creating mongodb doc:", err)
	// }
	go grpcServ()

	host := utils.EnvVar("SERVER_HOST", "localhost")
	port := utils.EnvVar("SERVER_PORT", "8000")
	addr := host + ":" + port
	log.Infoln("GIN start serving on:", addr)
	router := routes.InitRoute()
	router.Run(addr)
}
