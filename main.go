package main

import (
	"clean_gin_api/app"

	"go.uber.org/fx"
)

func main() {
	// create a Fx application
	fx.New(app.Module).Run()
}
