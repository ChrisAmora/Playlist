package infra

import (
	"context"
	"database/sql"
	"errors"

	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/domain"
)

type postgresMusicRepository struct {
	Conn *sql.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewMysqlArticleRepository(Conn *sql.DB) data.MusicRepository {
	return &postgresMusicRepository{Conn}
}

func (pm *postgresMusicRepository) Add(c context.Context, id int64) (domain.Music, error) {
	return domain.Music{}, errors.New("")
}

func (pm *postgresMusicRepository) GetById(c context.Context, id int64) (domain.Music, error) {
	return domain.Music{}, errors.New("")
}
