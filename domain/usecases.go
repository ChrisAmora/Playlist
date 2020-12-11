package domain

import "context"

type MusicUsecase interface {
	GetById(c context.Context, id int64) (Music, error)
	Add(c context.Context, id int64) (Music, error)
	GetAllMusics(c context.Context) ([]Music, error)
}