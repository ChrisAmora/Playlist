package data

import (
	"context"

	"github.com/betopompolo/project_playlist_server/domain"
)

type MusicRepository interface {
	GetById(c context.Context, id int64) (domain.Music, error)
	Add(c context.Context, id int64) (domain.Music, error)
}

type musicUsecase struct {
	MusicRepository
}

func NewMusicUsecase(mr MusicRepository) domain.MusicUsecase {
	return &musicUsecase{
		MusicRepository: mr,
	}
}

func (mu *musicUsecase) Add(c context.Context, id int64) (domain.Music, error) {
	return mu.MusicRepository.Add(c, id)
}

func (mu *musicUsecase) GetById(c context.Context, id int64) (domain.Music, error) {
	return mu.MusicRepository.GetById(c, id)
}
