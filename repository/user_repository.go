package repository

import (
	"clean_gin_api/infrastructure"
	"clean_gin_api/lib"

	"gorm.io/gorm"
)

// UserRepository database structure
type UserRepository struct {
	infrastructure.Database
	logger lib.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db infrastructure.Database, logger lib.Logger) UserRepository {
	return UserRepository{db, logger}
}

// WithTrx delegate transaction from user repository
func (r UserRepository) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
