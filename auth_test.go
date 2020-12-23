package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/betopompolo/project_playlist_server/config"
	"github.com/betopompolo/project_playlist_server/data"
	"github.com/betopompolo/project_playlist_server/presentation/models"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestCreateUser(t *testing.T) {
	c := client.New(config.AppInstance.Server)

	tcs := []struct {
		name     string
		email    string
		password string
	}{
		{
			name:     "Should create a user",
			email:    "efewfuhhuef@gmail.com",
			password: "fwefuehfwheuf",
		},
	}
	tcsE := []struct {
		name      string
		email     string
		password  string
		errorType gqlerror.Error
	}{
		{
			name:      "Should fail if invalid email",
			email:     "a",
			password:  "fwefuehfwheuf",
			errorType: gqlerror.Error{Message: "email: must be a valid email address."},
		},
	}

	for _, tc := range tcs {
		var resp struct {
			CreateUser models.User
		}
		t.Run(tc.name, func(t *testing.T) {
			request := fmt.Sprintf(`mutation { CreateUser(input: {email: "%s", password: "%s"}) { email } }`, tc.email, tc.password)
			c.Post(request, &resp)
			auth := data.Auth{}
			defer func() {
				config.AppInstance.DB.Unscoped().Delete(&auth)
			}()
			config.AppInstance.DB.Where("email = ?", tc.email).First(&auth)
			if resp.CreateUser.Email != tc.email || resp.CreateUser.Email != auth.Email {
				t.Errorf("Want '%s', got '%s'", tc.email, resp.CreateUser.Email)
			}
		})
	}

	for _, tc := range tcsE {
		var resp struct {
		}
		t.Run(tc.name, func(t *testing.T) {
			var gqlErrors []gqlerror.Error

			request := fmt.Sprintf(`mutation { CreateUser(input: {email: "%s", password: "%s"}) { email } }`, tc.email, tc.password)
			err := c.Post(request, &resp)
			if err := json.Unmarshal([]byte(err.Error()), &gqlErrors); err != nil {
				t.Error()
			}
			if gqlErrors[0].Message != tc.errorType.Message {
				t.Error()
			}
		})
	}

	t.Run("Fail if user already exists", func(t *testing.T) {
		var resp struct {
			CreateUser models.User
		}
		var resp2 struct{}
		email := "betin@gmail.com"
		password := "passworddobetin"
		request := fmt.Sprintf(`mutation { CreateUser(input: {email: "%s", password: "%s"}) { email } }`, email, password)
		c.Post(request, &resp)
		auth := data.Auth{}
		defer func() {
			config.AppInstance.DB.Unscoped().Delete(&auth)
		}()
		err := c.Post(request, &resp2)
		if err.Error() != `[{"message":"pq: duplicate key value violates unique constraint \"auths_email_key\"","path":["CreateUser"]}]` {
			t.Error()
		}

	})

}
