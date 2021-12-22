package services

import "go.uber.org/fx"

var Module = fx.Options(
	// list all services providers
	fx.Provide(NewUserService),
	fx.Provide(NewSMSService),
)
