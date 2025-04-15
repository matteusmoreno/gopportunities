package config

import (
	"github.com/matteusmoreno/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ()

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger()

	dsn := "host=localhost user=postgres password=postgres dbname=gopportunities_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Postgres opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Postgres automigration error: %v", err)
	}

	return db, nil
}
