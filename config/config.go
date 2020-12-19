package config

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
	Config *Config
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
		Port int64
		User string
		Pass string
		Name string
	}
}

func GetConf() *Config {
	conf := &Config{}
	// err := viper.Unmarshal(&conf)
	// if err != nil {
	// 	panic(err)
	// }
	conf.Database.Host = "localhost"
	conf.Database.Port = 5432
	conf.Database.User = "postgres"
	conf.Database.Pass = "postgres"
	conf.Database.Name = "postgres"
	conf.Context.Timeout = 2
	conf.Debug = true
	conf.Jwt.Secret = "betin"
	conf.Server.Address = "8080"
	return conf
}

func Setup() *App {
	// viper.SetConfigFile(`config.json`)
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	panic(err)
	// }
	a := App{}
	// if err != nil {
	// 	panic(err)
	// }
	a.Initialize()
	a.initializeGraphql()
	a.RunGraphql()
	return &a
}

func (a *App) Initialize() {
	a.Config = GetConf()
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", a.Config.Database.User, a.Config.Database.Pass, a.Config.Database.Name)

	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := Automigrate(a.DB); err != nil {
		panic(err)
	}

	a.Router = mux.NewRouter()
	a.Router.Use(handlers.Auth(infra.NewJWTService(a.Config.Jwt.Secret)))

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

func (a *App) RunGraphql() {
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", a.Config.Server.Address)
	log.Fatal(http.ListenAndServe(":"+a.Config.Server.Address, a.Router))
}

func (a *App) initializeRoutes() {
	c := GetConf()

	r := registry.NewRegistry(a.DB, c.Jwt.Secret)
	infra.NewRouter(a.Router, r.NewAppController())
}

func (a *App) initializeGraphql() {
	r := registry.NewRegistry(a.DB, a.Config.Jwt.Secret)
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
