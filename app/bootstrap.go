package app

import (
	controller "clean_gin_api/api/controllers"
	"clean_gin_api/api/routes"
	"clean_gin_api/models"
	"clean_gin_api/respository"
	"clean_gin_api/services"
	"context"
	"fmt"

	"go.uber.org/fx"
)

// func fx.Options(opts ...fx.Option) fx.Option
// Options converts a collection of Options into a single Option.
// This allows packages to bundle sophisticated functionality into easy-to-use Fx modules

//  Module exported for initializing application
var Module = fx.Options(
	// Group Providers...
	routes.Module,
	controller.Module,
	services.Module,
	respository.Module,
	models.Module,
)

func bootstrap(
	lifecycle fx.Lifecycle,

) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("starting Application.......")
			return nil

		},
		OnStop: func(context.Context) error {
			fmt.Println("closing Application")
			return nil
		},
	})
}
