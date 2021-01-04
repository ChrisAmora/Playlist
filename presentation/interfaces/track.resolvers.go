package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/betopompolo/project_playlist_server/presentation/models"
)

func (r *mutationResolver) CreateTrack(ctx context.Context, input models.TrackInput) (*models.Track, error) {
	track, err := r.TrackUsecase.SaveTrack(ctx, input.PlayListID, input.Title, input.Album, input.Artist)
	trackModel := &models.Track{}
	if err != nil {
		return trackModel, err
	}
	trackModel.Album = track.Album
	trackModel.Artist = track.Artist
	trackModel.CreatedAt = track.CreatedAt
	trackModel.PlayListID = int(track.PlayListID)
	trackModel.Title = track.Title
	return trackModel, err
}
