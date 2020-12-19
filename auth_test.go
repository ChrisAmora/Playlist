package main

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/betopompolo/project_playlist_server/config"
	"github.com/betopompolo/project_playlist_server/presentation/models"
	"github.com/stretchr/testify/require"
)

func TestAuth(t *testing.T) {
	c := client.New(config.AppInstance.Server)

	t.Run("Should return a valid token and email", func(t *testing.T) {
		var resp struct {
			Login models.AuthResponse
		}
		c.Post(`mutation {Login(input: {email: "a@gmail.com", password: "a"}) { user { email } token }}`, &resp)
		require.Equal(t, "a@gmail.com", resp.Login.User.Email)
		require.NotEmpty(t, resp.Login.Token)
	})

	t.Run("Fail if invalid email format", func(t *testing.T) {
		var resp struct{}
		err := c.Post(`mutation {Login(input: {email: "a", password: "a"}) { user { email } token }}`, &resp)
		require.EqualError(t, err, `[{"message":"Unprocessable Entity","path":["Login"],"extensions":{"customMessage":"email: must be a valid email address.","httpStatusCode":"422"}}]`)
	})

	t.Run("Fail if invalid credentials", func(t *testing.T) {
		var resp struct{}
		err := c.Post(`mutation {Login(input: {email: "a@gmail.com", password: "afefe"}) { user { email } token }}`, &resp)
		require.EqualError(t, err, `[{"message":"Unauthorized","path":["Login"],"extensions":{"customMessage":"Please provide a valid email or password","httpStatusCode":"401"}}]`)
		err = c.Post(`mutation {Login(input: {email: "afwfw@gmail.com", password: "afefe"}) { user { email } token }}`, &resp)
		require.EqualError(t, err, `[{"message":"Unauthorized","path":["Login"],"extensions":{"customMessage":"Please provide a valid email or password","httpStatusCode":"401"}}]`)
	})

}
