package interfaces

import "github.com/betopompolo/project_playlist_server/domain"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MusicService domain.MusicUsecase
	UserService  domain.AuthUsecase
}
