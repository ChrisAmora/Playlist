package infra

import (
	"github.com/betopompolo/project_playlist_server/presentation"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, c presentation.AppController) *mux.Router {
	r.HandleFunc("/musics", c.Music.FetchMusics).Methods("Get")
	return r
}
