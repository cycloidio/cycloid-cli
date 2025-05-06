package middleware_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

var (
	m                 middleware.Middleware
	api               *common.APIClient
	config            *TestConfig
	configRepoName    = "CLI test config repo"
	configRepo        = "config"
	configRepoUrl     = "git@github.com:cycloidio/cycloid-cli-test-catalog.git"
	configRepoBranch  = "config"
	isDefault         = false
	gitCred           = "github"
	catalogRepo       = "cli-test-catalog"
	catalogRepoName   = "CLI test catalog"
	catalogRepoUrl    = "git@github.com:cycloidio/cycloid-cli-test-catalog.git"
	catalogRepoBranch = "stacks"
	// gitCredName = "CLI Git Cred"
	// gitCredKey  = ""
)

type TestConfig struct {
	APIKey string
	APIUrl string
	Org    string
	// GitCredential models.Credential
	ConfigRepo  models.ConfigRepository
	CatalogRepo models.ServiceCatalogSource
	Middleware  middleware.Middleware
}

func getTestConfig() (*TestConfig, error) {
	var apiUrl, apiKey, org string
	apiUrl, ok := os.LookupEnv("CY_API_URL")
	if !ok {
		apiUrl = "https://api-cli-test.staging.cycloid.io/"
	}

	org, ok = os.LookupEnv("CY_TEST_ROOT_ORG")
	if !ok {
		org = "cycloid"
	}

	apiKey, ok = os.LookupEnv("CY_TEST_API_KEY")
	if !ok {
		return nil, fmt.Errorf("api key not set in CY_TEST_API_KEY env var.")
	}

	api = common.NewAPI(
		common.WithURL(apiUrl),
		common.WithInsecure(true),
		common.WithToken(apiKey),
	)

	m = middleware.NewMiddleware(api)
	return &TestConfig{
		APIUrl: apiUrl,
		APIKey: apiKey,
		Org:    org,
		ConfigRepo: models.ConfigRepository{
			Name:                &configRepoName,
			Canonical:           &configRepo,
			Default:             &isDefault,
			URL:                 &configRepoUrl,
			Branch:              configRepoBranch,
			CredentialCanonical: gitCred,
		},
		CatalogRepo: models.ServiceCatalogSource{
			Name:                &catalogRepoName,
			Canonical:           &catalogRepo,
			URL:                 &catalogRepoUrl,
			Branch:              catalogRepoBranch,
			CredentialCanonical: gitCred,
		},
		Middleware: m,
	}, nil
}

// Put any preparation code here so that defer() can work
func runMain(ctx context.Context, main *testing.M) (int, error) {
	_ = ctx
	// Initialize global vars
	config, err := getTestConfig()
	if err != nil {
		return 1, fmt.Errorf("Config setup failed: %v", err)
	}

	log.Printf("Starting tests with config:\nurl: %s\norg: %s", config.APIUrl, config.Org)
	return main.Run(), nil
}

func TestMain(main *testing.M) {
	ctx := context.Background()
	code, err := runMain(ctx, main)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

// CreateTestChildOrg is a helper function that create a suborg for a
// specific test and return a function to defer() for its deletion.
func CreateTestChildOrg(m middleware.Middleware, parent, child string) (func(), error) {
	deferFunc := func() {
		err := m.DeleteOrganization(child)
		if err != nil {
			log.Fatalf("Failed to delete org '%s': %v", child, err)
			return
		}
	}

	_, err := m.CreateOrganizationChild(parent, child, nil)
	if err != nil {
		return deferFunc, fmt.Errorf("Failed to create child org '%s': %v", child, err)
	}

	return deferFunc, nil
}
