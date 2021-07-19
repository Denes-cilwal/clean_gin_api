package routes

import (
	controller "clean_gin_api/api/controllers"
	"clean_gin_api/infrastructure"
)

// UserRoutes struct
type UserRoutes struct {
	handler        infrastructure.Router
	userController controller.UserController
}

func NewUserRoutes(
	handler infrastructure.Router,
	userController controller.UserController,

) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	api := s.handler.Group("/api")
	{
		api.POST("/user", s.userController.SaveUser)
	}
}
