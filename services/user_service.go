package services

import (
	"clean_gin_api/models"
)

type UserService struct {
	repository models.User
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(userrepo models.User) UserService {
	return UserService{
		repository: userrepo,
	}
}
