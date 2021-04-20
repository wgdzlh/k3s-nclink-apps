package main

import (
	"crypto/tls"
	"k3s-nclink-apps/config-distribute/routes"
	"k3s-nclink-apps/grpc-common/middlewares"
	"k3s-nclink-apps/utils"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	log "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {
	host := utils.EnvVar("SERVER_HOST", "localhost")
	port := utils.GetEnvOrExit("SERVER_PORT")
	// router := routes.InitRoute()
	// router.Run(host + ":" + port)
	serverCert := utils.GetEnvOrExit("SERVER_CRT")
	serverkey := utils.GetEnvOrExit("SERVER_KEY")

	stage := utils.GetEnvOrExit("DEV_STAGE")

	cert, err := tls.LoadX509KeyPair(utils.Path(serverCert), utils.Path(serverkey))
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	addr := host + ":" + port
	lis, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		// The following grpc.ServerOption adds an interceptor for all unary
		// RPCs. To configure an interceptor for streaming RPCs, see:
		// https://godoc.org/google.golang.org/grpc#StreamInterceptor
		grpc.UnaryInterceptor(middlewares.EnsureValid),
		// Enable TLS for all incoming connections.
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	server := grpc.NewServer(opts...)
	routes.RegisterServices(server)
	if stage == "debug" {
		reflection.Register(server)
	}

	log.Infoln("start serving on:", addr)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
