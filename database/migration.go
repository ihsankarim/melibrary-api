package database

import (
	"github.com/ihsankarim/melibrary-api/config"
	"github.com/ihsankarim/melibrary-api/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Category{})
}
