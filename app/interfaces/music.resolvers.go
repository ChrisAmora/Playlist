package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/betopompolo/project_playlist_server/app/generated"
	"github.com/betopompolo/project_playlist_server/app/models"
)

func (r *mutationResolver) CreateMusic(ctx context.Context, title string) (*models.MusicResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOneMusic(ctx context.Context, id string) (*models.MusicResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllMusics(ctx context.Context) (*models.MusicResponse, error) {
	r.MusicService.GetAllMusics(ctx)
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
