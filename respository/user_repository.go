package respository

import (
	"clean_gin_api/infrastructure"
	"clean_gin_api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

// WithTrx delegate transaction from user repository
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	r.db.DB = trxHandle
	return r
}

// create Users

func (u UserRepository) SaveUser(user models.User) error {
	return u.db.DB.Create(user).Error
}

// Get single user
func (r UserRepository) GetOne(id uint) (user models.User, err error) {
	return user, r.db.DB.Where("id = ?", id).First(&user).Error
}
