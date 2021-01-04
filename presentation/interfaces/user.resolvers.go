package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/betopompolo/project_playlist_server/domain"
	gqlerrors "github.com/betopompolo/project_playlist_server/presentation/gql-errors"
	"github.com/betopompolo/project_playlist_server/presentation/models"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	err := validation.ValidateStruct(&input, validation.Field(&input.Email, is.Email))
	if err != nil {
		return &models.User{}, err
	}

	user, err := r.UserUsecase.Signup(ctx, input.Email, input.Password)

	if err != nil {
		if _, ok := err.(*domain.RequestError); ok {
			return &models.User{}, gqlerrors.GraphqlUserAlreadyExist(ctx)
		}
		return &models.User{}, err
	}

	return &models.User{Email: user.Email}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models.UserInput) (*models.Auth, error) {
	err := validation.ValidateStruct(&input, validation.Field(&input.Email, is.Email))
	if err != nil {
		return &models.Auth{}, gqlerrors.GraphqlInvalidInput(ctx, err.Error())
	}
	auth, err := r.UserUsecase.Login(ctx, input.Email, input.Password)
	if err != nil {
		return &models.Auth{}, gqlerrors.GraphqlUnauthorized(ctx, "Please provide a valid email or password")
	}

	return &models.Auth{User: &models.User{Email: auth.User.Email}, Token: auth.Token}, err
}
