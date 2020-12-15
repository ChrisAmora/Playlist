package core

import (
	"context"
	"errors"
	"time"

	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	secret string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWTService(secret string) domain.JWTUsecase {
	return &JWTService{secret: secret}
}

func (js *JWTService) Sign(c context.Context, username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
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

func (js *JWTService) Verify(c context.Context, token string) (*jwt.Token, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(js.secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return &jwt.Token{}, errors.New("bad request")
		}
	}
	if !tkn.Valid {
		return &jwt.Token{}, errors.New("unouth")
	}

	return tkn, nil
}
