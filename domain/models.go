package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Music struct {
	ID        int64
	Title     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type User struct {
	Email    string
	Password string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
