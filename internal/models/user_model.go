package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          string `gorm:"not null; unique"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(100)"`
	Email       string `gorm:"type:varchar(100);not null;unique"`
	Password    string `gorm:"type:varchar(100);not null"`
}
