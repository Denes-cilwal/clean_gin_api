package routes

import (
	controllers "clean_gin_api/api/controllers"
	"clean_gin_api/infrastructure"
	"clean_gin_api/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	userController controllers.UserController
}

func NewUserRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	userController controllers.UserController,

) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Group("/api")
	{
		api.POST("/user", s.userController.SaveUser)

	}
}
