package config

import (
	"context"
	pb "k3s-nclink-apps/configmodel"
	"k3s-nclink-apps/utils"
	"log"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type Model struct {
	token *oauth2.Token
	creds credentials.TransportCredentials
}

var (
	configUser = utils.GetEnvOrExit("CONFIG_USER")
	configPass = utils.GetEnvOrExit("CONFIG_PASS")
	configHost = utils.GetEnvOrExit("CONFIG_HOST")
	configPort = utils.GetEnvOrExit("CONFIG_PORT")
	caCert     = utils.GetEnvOrExit("CA_CRT")
	configAddr = configHost + ":" + configPort
)

func NewModel() *Model {
	model := Model{}
	creds, err := credentials.NewClientTLSFromFile(utils.Path(caCert), "")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	model.creds = creds
	return &model
}

func (m *Model) getToken() *oauth2.Token {
	if m.token != nil {
		return m.token
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(m.creds),
		grpc.WithBlock(),
	}

	conn, err := grpc.Dial(configAddr, opts...)
	if err != nil {
		log.Fatalf("getToken could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthenticationClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := client.Login(ctx, &pb.LoginRequest{Name: configUser, Password: configPass})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	} else {
		log.Printf("token: %s\n", resp.Token)
	}
	m.token = &oauth2.Token{AccessToken: resp.Token}
	return m.token
}

func (m *Model) Fetch(hostname string) string {
	// Set up the credentials for the connection.
	perRPC := oauth.NewOauthAccess(m.getToken())
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithTransportCredentials(m.creds),
		grpc.WithBlock(),
	}

	conn, err := grpc.Dial(configAddr, opts...)
	if err != nil {
		log.Fatalf("Fetch could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewModelDistClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	resp, err := client.GetModel(ctx, &pb.ModelRequest{Hostname: hostname})
	if err != nil {
		log.Fatalf("GetModel failed: %v", err)
	}
	return resp.ModelJson
}
