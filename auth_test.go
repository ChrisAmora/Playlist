package main

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/betopompolo/project_playlist_server/config"
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/presentation/models"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuth(t *testing.T) {
	c := client.New(config.AppInstance.Server)
	email := "betindaora22@gmail.com"
	password := "lasanhaboa"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	auth := data.Auth{Email: email, Password: string(hash)}
	config.AppInstance.DB.Create(&auth)

	t.Cleanup(func() {
		config.AppInstance.DB.Unscoped().Delete(&auth)
	})

	tcs := []struct {
		name     string
		email    string
		password string
		request  string
	}{
		{
			name:     "Should return a valid token and email",
			email:    email,
			password: password,
			request:  fmt.Sprintf(`mutation {Login(input: {email: "%s", password: "%s"}) { user { email } token }}`, email, password),
		},
	}

	for _, tc := range tcs {
		var resp struct {
			Login models.AuthResponse
		}
		t.Run(tc.name, func(t *testing.T) {
			c.Post(tc.request, &resp)
			if resp.Login.User.Email != tc.email {
				t.Errorf("Want '%s', got '%s'", tc.email, resp.Login.User.Email)
			}
		})
	}

	t.Run("Should return a valid token and email", func(t *testing.T) {
		var resp struct {
			Login models.AuthResponse
		}

		c.Post(fmt.Sprintf(`mutation {Login(input: {email: "%s", password: "%s"}) { user { email } token }}`, email, password), &resp)
		require.Equal(t, email, resp.Login.User.Email)
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
