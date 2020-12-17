package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/presentation/generated"
	"github.com/betopompolo/project_playlist_server/presentation/handlers"
	"github.com/betopompolo/project_playlist_server/presentation/interfaces"
	"github.com/betopompolo/project_playlist_server/registry"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	c := GetConf()
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.Database.User, c.Database.Pass, c.Database.Name)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := Automigrate(a.DB); err != nil {
		panic(err)
	}

	a.Router = mux.NewRouter()
	a.Router.Use(handlers.Auth(infra.NewJWTService(c.Jwt.Secret)))

	a.initializeGraphql()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) RunGraphql() {
	c := GetConf()
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", c.Server.Address)
	log.Fatal(http.ListenAndServe(":"+c.Server.Address, a.Router))
}

func (a *App) initializeRoutes() {
	c := GetConf()

	r := registry.NewRegistry(a.DB, c.Jwt.Secret)
	infra.NewRouter(a.Router, r.NewAppController())
}

func (a *App) initializeGraphql() {
	c := GetConf()
	r := registry.NewRegistry(a.DB, c.Jwt.Secret)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		MusicService: r.NewMusicUseCase(),
		UserService:  r.NewAuthUseCase(),
	}}))

	srv.Use(extension.Introspection{})
	a.Router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	a.Router.Handle("/query", srv)

}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&data.Music{}, &data.Auth{}).Error
}
