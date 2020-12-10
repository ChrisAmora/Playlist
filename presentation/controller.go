package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/betopompolo/project_playlist_server/domain"
)

type MusicController interface {
	FetchMusics(w http.ResponseWriter, r *http.Request)
}

type musicController struct {
	domain.MusicUsecase
}

func NewMusicController(dm domain.MusicUsecase) MusicController {
	return &musicController{
		MusicUsecase: dm,
	}
}

func (mh *musicController) FetchMusics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mh.MusicUsecase.GetAllMusics(ctx)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("betinho")
}
