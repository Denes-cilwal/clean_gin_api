package models

import "time"

// user model

type User struct {
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Age       int        `json:"age"`
	Birthday  *time.Time `json:"time"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
}

// TableName gives table name of model
func (u User) TableName() string {
	return "users"
}
