package domain

import (
	"fmt"
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

type RequestError struct {
	Err error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("Request Error %v", r.Err)
}
