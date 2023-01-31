package bundle

import (
	"context"
	"fxapp/config"
	"fxapp/db"
	"fxapp/logger"
	"fxapp/proto"
	"fxapp/router"
	"fxapp/rpc"
	"fxapp/server"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	cfg *config.Config,
	server *server.Server,
	addService rpc.AddService,
	logger *zap.SugaredLogger,
) {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				//go http.ListenAndServe(cfg.ApplicationConfig.Address, server.)
				// rpc server
				lis, err := net.Listen("tcp", cfg.ApplicationConfig.RpcAddress)
				if err != nil {
					logger.Fatalf("failed to listen: %v", err)
				}
				var opts []grpc.ServerOption
				grpcServer := grpc.NewServer(opts...)
				proto.RegisterAddServiceServer(grpcServer, addService)
				go func() {
					if e := grpcServer.Serve(lis); e != nil {
						panic(err)
					}
				}()

				go func() {
					server.Logger.Fatal(server.App.Listen(cfg.ApplicationConfig.Address))
				}()
				server.Logger.Info("Server is listening on port" + cfg.ApplicationConfig.Address)

				return nil
			},
			OnStop: func(context.Context) error {
				return server.Logger.Sync()
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	config.Module,
	logger.Module,
	db.Module,
	server.Module,
	router.BookModule,
	rpc.AddServiceModule,
	rpc.ConnectorModule,
	router.AddServiceModule,
	fx.Invoke(registerHooks),
)
