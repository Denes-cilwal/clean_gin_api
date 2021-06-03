package infrastructure

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database model
type Database struct {
	*gorm.DB
}

// NewDatabase creates a new database instance
func NewDatabase(env Env) Database {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{Logger: newLogger})
	_ = db.Exec("CREATE DATABASE IF NOT EXISTS " + env.DBName + ";")
	if err != nil {
		Zaplogger.Zap.Info("Url: ", url)
		Zaplogger.Zap.Panic(err)
	}

	Zaplogger.Zap.Info("Database connection established")

	return Database{DB: db}

}
