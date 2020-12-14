package infra

import (
	"github.com/betopompolo/project_playlist_server/rest"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, c rest.AppController) *mux.Router {
	r.HandleFunc("/musics", c.Music.FetchMusics).Methods("Get")
	return r
}
