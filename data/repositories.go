package data

import (
	"context"

	"github.com/betopompolo/project_playlist_server/domain"
)

type MusicRepository interface {
	GetById(c context.Context, id int64) (domain.Music, error)
	Add(c context.Context, id int64) (domain.Music, error)
	GetAll(c context.Context) ([]domain.Music, error)
}

type AuthRepository interface {
	Signup(c context.Context, email, password string) (Auth, error)
}

type authUsecase struct {
	AuthRepository
}

type musicUsecase struct {
	MusicRepository
}

func NewMusicUsecase(mr MusicRepository) domain.MusicUsecase {
	return &musicUsecase{
		MusicRepository: mr,
	}
}

func NewAuthUsecase(ar AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		AuthRepository: ar,
	}
}

func (au *authUsecase) Signup(c context.Context, email, password string) (bool, error) {
	au.AuthRepository.Signup(c, email, password)
	return true, nil
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
