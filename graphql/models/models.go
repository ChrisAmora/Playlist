// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type Music struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MusicResponse struct {
	Message  string   `json:"message"`
	Status   int      `json:"status"`
	Data     *Music   `json:"data"`
	DataList []*Music `json:"dataList"`
}

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    *User  `json:"data"`
}
