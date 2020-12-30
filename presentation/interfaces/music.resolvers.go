package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/betopompolo/project_playlist_server/presentation/generated"
	models1 "github.com/betopompolo/project_playlist_server/presentation/models"
)

func (r *mutationResolver) CreateMusic(ctx context.Context, title string) (*models1.Music, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOneMusic(ctx context.Context, id string) (*models1.Music, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllMusics(ctx context.Context) ([]*models1.Music, error) {
	r.MusicUsecase.GetAllMusics(ctx)
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
