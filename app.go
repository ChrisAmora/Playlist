package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/graphql/generated"
	"github.com/betopompolo/project_playlist_server/graphql/interfaces"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/registry"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(user string, password string, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := Automigrate(a.DB); err != nil {
		panic(err)
	}
	a.DB.Create(&data.Music{Title: "lasanha"})

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
	r := registry.NewRegistry(a.DB)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		MusicService: r.NewMusicUseCase(),
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&data.Music{}).Error
}
