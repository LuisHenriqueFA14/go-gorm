package database

import (
	"gorm.io/gorm"
	"github.com/LuisHenriqueFA14/go-gorm/internal/models"
)

var Db *gorm.DB

func CreateDatabase(db *gorm.DB) {
	Db = db

	Db.AutoMigrate(&models.User{})
}

