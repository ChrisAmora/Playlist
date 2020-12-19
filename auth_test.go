package main

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/betopompolo/project_playlist_server/config"
	"github.com/betopompolo/project_playlist_server/presentation/generated"
	"github.com/betopompolo/project_playlist_server/presentation/interfaces"
	"github.com/betopompolo/project_playlist_server/registry"
)

func TestAuth(t *testing.T) {
	a := config.App{}
	a.Initialize()
	r := registry.NewRegistry(a.DB, a.Config.Jwt.Secret)
	aa := generated.Config{Resolvers: &interfaces.Resolver{
		MusicService: r.NewMusicUseCase(),
		UserService:  r.NewAuthUseCase(),
	}}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(aa)))

	t.Run("Auth", func(t *testing.T) {
		var resp struct {
			Login struct {
				token string
				user  struct{ email string }
			}
		}
		c.MustPost(`mutation {Login(input: {email: "a@gmail.com", password: "a"}) { user { email } token }}`, &resp)
		fmt.Println(resp)
	})

}
