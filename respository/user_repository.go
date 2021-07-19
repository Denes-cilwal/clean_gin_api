package respository

import (
	"clean_gin_api/infrastructure"
	"clean_gin_api/models"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// create Users

func (u UserRepository) SaveUser(user models.User) error {
	return u.db.DB.Create(user).Error
}
