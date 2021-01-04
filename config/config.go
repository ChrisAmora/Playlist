package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/infra"
	"github.com/betopompolo/project_playlist_server/presentation/generated"
	"github.com/betopompolo/project_playlist_server/presentation/handlers"
	"github.com/betopompolo/project_playlist_server/presentation/interfaces"
	"github.com/betopompolo/project_playlist_server/registry"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var AppInstance *App

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Config *Config
	Server *handler.Server
}

type Config struct {
	Debug  bool
	Server struct {
		Address string
	}
	Jwt struct {
		Secret string
	}
	Context struct {
		Timeout int64
	}
	Database struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

func GetConf() *Config {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	conf := &Config{}
	err = viper.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func Setup() *App {
	a := App{}
	a.Config = GetConf()
	a.InitializeDB()
	a.InitializeRouter()
	a.InitializeGraphql()
	return &a
}

func (a *App) InitializeDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", a.Config.Database.Host, a.Config.Database.User, a.Config.Database.Pass, a.Config.Database.Name, a.Config.Database.Port)
	var err error
	a.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := Automigrate(a.DB); err != nil {
		panic(err)
	}

}

func (a *App) InitializeRouter() {
	a.Router = mux.NewRouter()
	a.Router.Use(handlers.Auth(infra.NewJWTService(a.Config.Jwt.Secret)))
}

func (a *App) RunRest() {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) RunGraphql() {
	a.Router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	a.Router.Handle("/query", a.Server)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", a.Config.Server.Address)
	log.Fatal(http.ListenAndServe(":"+a.Config.Server.Address, a.Router))
}

func (a *App) initializeRoutesRest() {
	c := GetConf()

	r := registry.NewRegistry(a.DB, c.Jwt.Secret)
	infra.NewRouter(a.Router, r.NewAppController())
}

func (a *App) InitializeGraphql() {
	r := registry.NewRegistry(a.DB, a.Config.Jwt.Secret)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		MusicUsecase: r.NewMusicUseCase(),
		UserUsecase:  r.NewAuthUseCase(),
		TrackUsecase: r.NewTrackUsecase(),
	}}))
	a.Server = srv
}

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(&data.Music{}, &data.Auth{}, &data.PlayList{}, &data.Track{})
}
