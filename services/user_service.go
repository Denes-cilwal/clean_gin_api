package services

import (
	"clean_gin_api/models"
)

type UserService struct {
	repository models.UserRepo
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewUserService(userrepo models.UserRepo) UserService {
	return UserService{
		repository: userrepo,
	}
}

func (us UserService) CreateUser(user *models.User) (*models.User, error) {

	return us.repository.CreateUser(user)
}
