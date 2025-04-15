package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error initializing Postgres: %v", err)
	}
	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	logger = NewLogger()
	return logger
}
