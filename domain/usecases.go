package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type MusicUsecase interface {
	GetById(c context.Context, id int64) (Music, error)
	Add(c context.Context, id int64) (Music, error)
	GetAllMusics(c context.Context) ([]Music, error)
}

type AuthUsecase interface {
	Signup(c context.Context, email, password string) (User, error)
}

type JWTUsecase interface {
	Sign(c context.Context, username string) (string, error)
	Verify(c context.Context, token string) (*jwt.Token, error)
}
