package registry

import (
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/presentation"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() presentation.AppController
	NewMusicUseCase() domain.MusicUsecase
	NewAuthUseCase() domain.AuthUsecase
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() presentation.AppController {
	return presentation.AppController{
		Music: r.NewMusicController(),
	}
}

func (r *registry) NewMusicController() presentation.MusicController {
	return presentation.NewMusicController(r.NewMusicUseCase())
}

func (r *registry) NewMusicRepository() data.MusicRepository {
	return infra.NewPostgresMusicRepository(r.db)
}

func (r *registry) NewAuthRepository() data.AuthRepository {
	return infra.NewPostgresAuthRepository(r.db)
}

func (r *registry) NewMusicUseCase() domain.MusicUsecase {
	return data.NewMusicUsecase(r.NewMusicRepository())
}

func (r *registry) NewAuthUseCase() domain.AuthUsecase {
	return data.NewAuthUsecase(r.NewAuthRepository())
}
