package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/betopompolo/project_playlist_server/graphql/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, name string, password string) (*models.UserResponse, error) {
	user, err := r.UserService.Signup(ctx, name, password)
	return &models.UserResponse{Message: "efewhfuwe", Status: 900, Data: &models.User{ID: "few", Email: user.Email}}, err
}
