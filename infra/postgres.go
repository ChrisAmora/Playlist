package infra

import (
	"context"
	"errors"
	"fmt"

	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/jmoiron/sqlx"
)

type postgresMusicRepository struct {
	Conn *sqlx.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(Conn *sqlx.DB) data.MusicRepository {
	return &postgresMusicRepository{Conn}
}

func (pm *postgresMusicRepository) Add(c context.Context, id int64) (domain.Music, error) {
	return domain.Music{}, errors.New("")
}

func (pm *postgresMusicRepository) GetById(c context.Context, id int64) (domain.Music, error) {
	return domain.Music{}, errors.New("")
}

func (pm *postgresMusicRepository) GetAll(c context.Context) ([]domain.Music, error) {
	musics := []domain.Music{}
	err := pm.Conn.Select(&musics, "SELECT * FROM music")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(musics)
	return musics, nil
}
