package handler

import (
	"github.com/matteusmoreno/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger()
	db = config.GetPostgres()
}
