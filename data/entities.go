package data

import "time"

type Music struct {
	ID        int64
	Title     string
	UpdatedAt time.Time
	CreatedAt time.Time
}
