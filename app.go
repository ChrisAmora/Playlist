package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/presentation"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS music (
	id SERIAL NOT NULL PRIMARY KEY,
	title text,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`

type App struct {
	Router *mux.Router
	DB     *sqlx.DB
}

func (a *App) Initialize(user string, password string, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.DB.MustExec(schema)

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) initializeRoutes() {
	mr := infra.NewMysqlArticleRepository(a.DB)
	mu := data.NewMusicUsecase(mr)
	presentation.NewMusicHandler(a.Router, mu)
}
