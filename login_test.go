package main

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/betopompolo/project_playlist_server/config"
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/presentation/models"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
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
		name         string
		email        string
		password     string
		errorMessage string
	}{
		{
			name:         "Should return a valid token and email",
			email:        email,
			password:     password,
			errorMessage: "",
		},
		{
			name:         "Fail if invalid email format",
			email:        "a",
			password:     "a",
			errorMessage: `[{"message":"Unprocessable Entity","path":["Login"],"extensions":{"customMessage":"email: must be a valid email address.","httpStatusCode":"422"}}]`,
		},
		{
			name:         "Fail if invalid credentials",
			email:        email,
			password:     "aaaaaffefeffe",
			errorMessage: `[{"message":"Unauthorized","path":["Login"],"extensions":{"customMessage":"Please provide a valid email or password","httpStatusCode":"401"}}]`,
		},
	}

	for _, tc := range tcs {
		var resp struct {
			Login models.AuthResponse
		}
		t.Run(tc.name, func(t *testing.T) {
			request := fmt.Sprintf(`mutation {Login(input: {email: "%s", password: "%s"}) { user { email } token }}`, tc.email, tc.password)
			err := c.Post(request, &resp)
			if err != nil {
				if err.Error() != tc.errorMessage {
					t.Errorf("Want '%s', got '%s'", tc.errorMessage, err.Error())
				}
			} else {
				if resp.Login.User.Email != tc.email {
					t.Errorf("Want '%s', got '%s'", tc.email, resp.Login.User.Email)
				}
			}
		})
	}

}
