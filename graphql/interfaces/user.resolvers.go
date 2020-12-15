package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/betopompolo/project_playlist_server/graphql/models"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	err := validation.ValidateStruct(&input, validation.Field(&input.Email, is.Email))
	if err != nil {
		return &models.User{}, err
	}

	user, err := r.UserService.Signup(ctx, input.Email, input.Password)

	if err != nil {
		return &models.User{}, err
	}

	return &models.User{Email: user.Email}, nil
}
