package core

import (
	"context"
	"time"

	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secret string
}

func NewJWTService(secret string) domain.JWTUsecase {
	return &JWTService{secret}
}

func (js *JWTService) Sign(c context.Context, username string) (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &domain.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(js.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func (js *JWTService) Verify(c context.Context, token string) (*domain.Claims, error) {
	claims := &domain.Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.secret), nil
	})

	return claims, err
}
