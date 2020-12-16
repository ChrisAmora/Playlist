package domain

import (
	"time"
)

type Music struct {
	ID        int64
	Title     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Auth struct {
	User
	Token string
}

type User struct {
	Email string
}
