package services

import (
	"clean_gin_api/models"
	"clean_gin_api/respository"
)

type UserService struct {
	repository respository.UserRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(userrepo respository.UserRepository) UserService {
	return UserService{
		repository: userrepo,
	}
}

func (us UserService) CreateUser(user models.User) error {
	err := us.repository.SaveUser(user)
	return err
}
