package respository

import "go.uber.org/fx"

var Module = fx.Options(
	// list all repo providers
	fx.Provide(NewUserRepository),
)
