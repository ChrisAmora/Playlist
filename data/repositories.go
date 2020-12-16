package data

import (
	"context"

	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type MusicRepository interface {
	GetById(c context.Context, id int64) (domain.Music, error)
	Add(c context.Context, id int64) (domain.Music, error)
	GetAll(c context.Context) ([]domain.Music, error)
}

type AuthRepository interface {
	GetUser(c context.Context, email string) (Auth, error)
	CreateUser(c context.Context, email, password string) (*Auth, error)
}

type JWTRepository interface {
	Sign(c context.Context, username string) (string, error)
	Verify(c context.Context, token string) (*Claims, error)
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type authUsecase struct {
	AuthRepository
	JWTRepository
}

type musicUsecase struct {
	MusicRepository
}

func NewMusicUsecase(mr MusicRepository) domain.MusicUsecase {
	return &musicUsecase{
		MusicRepository: mr,
	}
}

func NewAuthUsecase(ar AuthRepository, jr JWTRepository) domain.AuthUsecase {
	return &authUsecase{
		AuthRepository: ar,
		JWTRepository:  jr,
	}
}

func (au *authUsecase) Signup(c context.Context, email, password string) (domain.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return domain.User{}, err
	}
	auth, err := au.AuthRepository.CreateUser(c, email, string(hash))
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{Email: auth.Email}, nil
}

func (au *authUsecase) Login(c context.Context, email, password string) (domain.Auth, error) {
	user, err := au.AuthRepository.GetUser(c, email)
	if err != nil {
		return domain.Auth{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.Auth{}, err
	}
	token, err := au.JWTRepository.Sign(c, email)
	if err != nil {
		return domain.Auth{}, err
	}
	return domain.Auth{User: domain.User{Email: email}, Token: token}, err
}

func (mu *musicUsecase) Add(c context.Context, id int64) (domain.Music, error) {
	return mu.MusicRepository.Add(c, id)
}

func (mu *musicUsecase) GetById(c context.Context, id int64) (domain.Music, error) {
	return mu.MusicRepository.GetById(c, id)
}

func (mu *musicUsecase) GetAllMusics(c context.Context) ([]domain.Music, error) {
	return mu.MusicRepository.GetAll(c)
}
