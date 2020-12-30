package data

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Title string
}

type Auth struct {
	gorm.Model
	Email    string `gorm:"primaryKey;unique;not null"`
	Password string
}

type PlayList struct {
	gorm.Model
	Name   string
	Tracks []Track
}

type Track struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index"`
	PlayListID uint         `gorm:"not null"`
	Title      string       `gorm:"primaryKey"`
	Album      string       `gorm:"primaryKey"`
	Artist     string       `gorm:"primaryKey"`
}
