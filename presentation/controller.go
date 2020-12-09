package presentation

import (
	"encoding/json"
	"net/http"

	"github.com/betopompolo/project_playlist_server/domain"
	"github.com/gorilla/mux"
)

type MusicHandler struct {
	domain.MusicUsecase
}

func NewMusicHandler(r *mux.Router, dm domain.MusicUsecase) {
	handler := &MusicHandler{
		MusicUsecase: dm,
	}
	r.HandleFunc("/musics", handler.FetchMusic).Methods("Get")

}

func (mh *MusicHandler) FetchMusic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mh.MusicUsecase.GetAllMusics(ctx)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("betinho")
}
