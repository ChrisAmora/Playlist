package domain

import (
	"context"
)

type MusicUsecase interface {
	GetById(c context.Context, id int64) (Music, error)
	Add(c context.Context, id int64) (Music, error)
	GetAllMusics(c context.Context) ([]Music, error)
}

type TrackUsecase interface {
	SaveTrack(c context.Context, playListID int, title, album, artist string) (Track, error)
}

type AuthUsecase interface {
	Signup(c context.Context, email, password string) (User, error)
	Login(c context.Context, email, password string) (Auth, error)
}
