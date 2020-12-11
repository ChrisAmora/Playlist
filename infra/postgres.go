package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/jinzhu/gorm"
)

type postgresMusicRepository struct {
	Conn *gorm.DB
}

func NewPostgresMusicRepository(Conn *gorm.DB) data.MusicRepository {
	return &postgresMusicRepository{Conn}
}

func (pm *postgresMusicRepository) Add(c context.Context, id int64) (domain.Music, error) {
	return domain.Music{}, errors.New("")
}

func (pm *postgresMusicRepository) GetById(c context.Context, id int64) (domain.Music, error) {
	return domain.Music{}, errors.New("")
}

func (pm *postgresMusicRepository) GetAll(c context.Context) ([]domain.Music, error) {
	musics := []data.Music{}
	err := pm.Conn.Find(&musics)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(musics)
	return []domain.Music{}, nil
}
