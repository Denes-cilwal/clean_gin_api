package respository

import (
	"clean_gin_api/infrastructure"
	"clean_gin_api/models"

	"gorm.io/gorm"
)

type SMSRepository struct {
	db infrastructure.Database
}

func NewSMSRepository(db infrastructure.Database) SMSRepository {
	return SMSRepository{
		db: db,
	}
}

// WithTrx delegate transaction from sms repository
func (s SMSRepository) WithTrx(trxHandle *gorm.DB) SMSRepository {
	s.db.DB = trxHandle
	return s
}

func (s SMSRepository) SendSMS(msg models.SMS) error {
	return s.db.Create(msg).Error
}
