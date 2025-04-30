package middleware_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

/*
1. Start the backend / mysql / vault
2. Init the backend
3. Create first org / admin / apiKey
4. Do the testing
5. Drop the tables
6. Repeat from 3
*/

var (
	backendUrl = "http://127.0.0.1:3001"
	username   = "cycloid"
	password   = "cycloidio"
	email      = "cycloid@example.com"
	defaultOrg = "cycloid"
)

func runMain(ctx context.Context, main *testing.M) (int, error) {
	_ = ctx
	return main.Run(), nil
}

func TestMain(main *testing.M) {
	// init first user/org
	api := common.NewAPI(common.WithURL(backendUrl))
	m := middleware.NewMiddleware(api)
	err := m.UserSignup(username, email, password, "cycloid", "cycloid")
	if err != nil {
		fmt.Println(err)
	}

	session, err := m.UserLogin(&email, &username, nil, password)
	if err != nil {
		log.Fatal(err)
	}
	api.Config.Token = *session.Token

	_, err = m.CreateOrganization(defaultOrg)
	var apiErr *middleware.ApiError
	if errors.As(err, &apiErr) && apiErr.HttpCode != "409" {
		log.Fatal(err)
	}

	refreshedToken, err := m.RefreshToken(&defaultOrg, &defaultOrg, *session.Token)
	if err != nil {
		log.Fatal(err)
	}
	api.Config.Token = *refreshedToken.Token

	licence, ok := os.LookupEnv("API_LICENCE_KEY")
	if !ok {
		log.Fatalf("Licence not set in API_LICENCE_KEY")
	}

	err = m.ActivateLicence(defaultOrg, licence)
	if err != nil {
		log.Fatal(err)
	}

	refreshedToken, err = m.RefreshToken(&defaultOrg, &defaultOrg, *session.Token)
	if err != nil {
		log.Fatal(err)
	}
	api.Config.Token = *refreshedToken.Token

	apiKey, err := m.CreateAPIKey(defaultOrg, "admin", "", username, &[]string{"hello"}[0],
		[]*models.NewRule{{
			Action:    &[]string{"organization:**"}[0],
			Effect:    &[]string{"allow"}[0],
			Resources: []string{},
		}})
	if err != nil {
		log.Fatal(err)
	}
	api.Config.Token = apiKey.Token

	file, err := os.Create("api_key")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(apiKey.Token)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(main.Run())
}

func TestBackend(t *testing.T) {
	fmt.Println("cool")
}
