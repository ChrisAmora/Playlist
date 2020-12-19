package main

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/betopompolo/project_playlist_server/config"
)

func TestAuth(t *testing.T) {
	c := client.New(config.AppInstance.Server)

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
