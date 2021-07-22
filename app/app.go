package app

import (
	controller "clean_gin_api/api/controllers"
	"clean_gin_api/api/middlewares"
	"clean_gin_api/api/routes"
	"clean_gin_api/infrastructure"
	"clean_gin_api/lib"
	"clean_gin_api/respository"
	"clean_gin_api/services"
	"context"

	"go.uber.org/fx"
)

/**

func fx.Options(opts ...fx.Option) fx.Option
Options converts a collection of Options into a single Option.
This allows packages to bundle sophisticated functionality into easy-to-use Fx modules

**/

//  Module exported for initializing application
var Module = fx.Options(
	// Group Providers...
	controller.Module,
	routes.Module,
	services.Module,
	infrastructure.Module,
	middlewares.Module,
	respository.Module,
	lib.Module,
	fx.Invoke(App),
)

func App(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	env lib.Env,
	middlewares middlewares.Middlewares,
	logger lib.Logger,
	database infrastructure.Database,
	migrations infrastructure.Migrations,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting Application")
			logger.Info("-------------------------------------")
			logger.Info("------- gin-apiu📺 -------")
			logger.Info("-------------------------------------")

			go migrations.Migrate()
			go func() {
				middlewares.Setup()
				routes.Setup()
				if env.ServerPort == "" {
					handler.Run()
				} else {
					handler.Run(":" + env.ServerPort)
				}

			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping Application")
			conn, _ := database.DB.DB()
			conn.Close()
			return nil
		},
	})
}
