package data

import (
	"github.com/jinzhu/gorm"
)

type Music struct {
	gorm.Model
	Title string
}

type Auth struct {
	gorm.Model
	Email    string `gorm:"primaryKey"`
	Password string
}
