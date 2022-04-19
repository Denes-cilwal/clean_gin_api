package controllers

import (
	"clean_gin_api/services"
)

type UserController struct {
	Userservice services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		Userservice: userService,
	}
}
