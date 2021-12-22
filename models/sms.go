package models

import (
	"clean_gin_api/lib"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Message model
type SMS struct {
	Base
	Message   string    `json:"message"`
	To        string    `json:"to"`
	From      string    `json:"From"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func (n *SMS) BeforeCreate(db *gorm.DB) error {
	id, err := uuid.NewRandom()
	n.ID = lib.BinaryUUID(id)
	return err
}
