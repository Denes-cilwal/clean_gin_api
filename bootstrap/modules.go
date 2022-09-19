package bootstrap

import (
	"clean_gin_api/api/controllers"
	"clean_gin_api/api/middlewares"
	"clean_gin_api/api/routes"
	"clean_gin_api/infrastructure"
	"clean_gin_api/lib"
	"clean_gin_api/repository"
	"clean_gin_api/seeds"
	"clean_gin_api/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	services.Module,
	repository.Module,
	infrastructure.Module,
	middlewares.Module,
	lib.Module,
	seeds.Module,
)
