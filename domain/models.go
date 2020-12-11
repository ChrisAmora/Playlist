package domain

import "time"

type Music struct {
	ID        int64
	Title     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Auth struct {
	ID        int64
	Email     string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
}
