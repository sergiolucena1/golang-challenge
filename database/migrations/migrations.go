package migrations

import (
	"golang-challenge/models"
	"gorm.io/gorm"
)

func RunMigrations (db *gorm.DB){
	db.AutoMigrate(models.Product{})
}
