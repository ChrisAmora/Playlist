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

type postgresAuthRepository struct {
	Conn *gorm.DB
}

func NewPostgresMusicRepository(Conn *gorm.DB) data.MusicRepository {
	return &postgresMusicRepository{Conn}
}

func NewPostgresAuthRepository(Conn *gorm.DB) data.AuthRepository {
	return &postgresAuthRepository{Conn}
}

func (ar *postgresAuthRepository) GetUser(c context.Context, email string) (data.Auth, error) {
	auth := &data.Auth{Email: email}
	db := ar.Conn.Take(auth)
	return *auth, db.Error
}

func (ar *postgresAuthRepository) CreateUser(c context.Context, email, password string) (*data.Auth, error) {
	auth := &data.Auth{Email: email, Password: password}
	result := ar.Conn.Create(&auth)
	return auth, result.Error
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
