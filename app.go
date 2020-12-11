package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/betopompolo/project_playlist_server/app/generated"
	"github.com/betopompolo/project_playlist_server/app/interfaces"
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/registry"
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

	a.initializeGraphql()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) RunGraphql(port string) {
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (a *App) initializeRoutes() {
	r := registry.NewRegistry(a.DB)
	infra.NewRouter(a.Router, r.NewAppController())
}

func (a *App) initializeGraphql() {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		MusicService: data.NewMusicUsecase(infra.NewPostgresMusicRepository(a.DB)),
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

}