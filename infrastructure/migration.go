package infrastructure

import (
	"clean_gin_api/lib"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//Migrations -> Migration Struct
type Migrations struct {
	logger lib.Logger
	env    lib.Env
}

//NewMigrations -> return new Migrations struct
func NewMigrations(logger lib.Logger, env lib.Env) Migrations {
	return Migrations{
		logger: logger,
		env:    env,
	}
}

//Migrate -> migrates all table
// Automatic migrations -> while running Go server
func (m Migrations) Migrate() {
	m.logger.Info("Migrating schemas...")

	USER := m.env.DBUsername
	PASS := m.env.DBPassword
	HOST := m.env.DBHost
	PORT := m.env.DBPort
	DBNAME := m.env.DBName

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	/*
				m, err := migrate.NewWithDatabaseInstance( "file:///migrations",  "mysql", driver)
		       OR
			     m, err := migrate.New("file://migrations/", dbType+"://"+dbURL)

	*/

	/*
		m.Steps(2) //
	*/

	migrations, err := migrate.New("file://migration/", "mysql://"+dbConfig)
	if err != nil {
		m.logger.Error("error in migration", err.Error())
	}

	m.logger.Info("--- Running Migration ---")
	err = migrations.Steps(1000)
	if err != nil {
		m.logger.Error("Error in migration: ", err.Error())
	}
}
