package registry

import (
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/rest"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db        *gorm.DB
	jwtSecret string
}

type Registry interface {
	NewAppController() rest.AppController
	NewMusicUseCase() domain.MusicUsecase
	NewAuthUseCase() domain.AuthUsecase
}

func NewRegistry(db *gorm.DB, jwtSecret string) Registry {
	return &registry{db, jwtSecret}
}

func (r *registry) NewAppController() rest.AppController {
	return rest.AppController{
		Music: r.NewMusicController(),
	}
}

func (r *registry) NewMusicController() rest.MusicController {
	return rest.NewMusicController(r.NewMusicUseCase())
}

func (r *registry) NewMusicRepository() data.MusicRepository {
	return infra.NewPostgresMusicRepository(r.db)
}

func (r *registry) NewAuthRepository() data.AuthRepository {
	return infra.NewPostgresAuthRepository(r.db)
}

func (r *registry) NewAJWTRepository() data.JWTRepository {
	return infra.NewJWTService(r.jwtSecret)
}

func (r *registry) NewMusicUseCase() domain.MusicUsecase {
	return data.NewMusicUsecase(r.NewMusicRepository())
}

func (r *registry) NewAuthUseCase() domain.AuthUsecase {
	return data.NewAuthUsecase(r.NewAuthRepository(), r.NewAJWTRepository())
}
