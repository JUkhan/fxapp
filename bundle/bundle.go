package bundle

import (
	"context"
	"fxapp/config"
	"fxapp/db"
	"fxapp/logger"
	"fxapp/router"
	"fxapp/server"

	"go.uber.org/fx"
)

func registerHooks(
	lifecycle fx.Lifecycle,
	cfg *config.Config, server *server.Server,
) {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				//go http.ListenAndServe(cfg.ApplicationConfig.Address, server.)

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
	fx.Invoke(registerHooks),
)
