package controllers

import "go.uber.org/fx"

//Module exports controller dependency
var Module = fx.Options(
	// list all controller providers
	fx.Provide(NewUserController),
	fx.Provide(NewSMSController),
)
