package rpc

import (
	"fxapp/config"
	"log"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Connector struct {
	Conn *grpc.ClientConn
}

func newConnection(cfg *config.Config) *Connector {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	return &Connector{conn}
}

// Module provided to fx
var ConnectorModule = fx.Options(
	fx.Provide(newConnection),
)
