package testcfg

import (
	"fmt"
	"log"
	"os"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

type Config struct {
	APIKey string
	APIUrl string
	Org    string
	// GitCredential models.Credential
	ConfigRepo  *models.ConfigRepository
	CatalogRepo *models.ServiceCatalogSource
	Middleware  middleware.Middleware
	// Common project to use for tests that require one
	Project *models.Project
	// Common environment to use for tests that require one
	Environment *models.Environment
	// Common component to use for tests that require one
	Component *models.Component

	// Slice containing all functions to exec for cleaning common test resources
	cleanupFuncs []func()
}

func NewConfig() (*Config, error) {
	var (
		configRepoName        = "stacks-test-config"
		configRepository      = "stacks-test-config"
		configRepoURL         = "git@github.com:cycloidio/cycloid-stacks-test.git"
		configRepoBranch      = "config"
		isDefault             = false
		gitCred               = "github"
		catalogRepo           = "cli-test-stacks"
		catalogRepoName       = "CLI test catalog"
		catalogRepoURL        = "git@github.com:cycloidio/cycloid-cli-test-catalog.git"
		catalogRepoBranch     = "stacks"
		defaultStackCanonical = "stack-e2e-stackforms"
		defaultStackUseCase   = "default"
		// gitCredName = "CLI Git Cred"
		// gitCredKey  = ""
	)

	var apiURL, apiKey, org string
	var config = &Config{}
	config.ConfigRepo = &models.ConfigRepository{
		Name:                &configRepoName,
		Canonical:           &configRepository,
		Default:             &isDefault,
		URL:                 &configRepoURL,
		Branch:              configRepoBranch,
		CredentialCanonical: gitCred,
	}
	config.CatalogRepo = &models.ServiceCatalogSource{
		Name:                &catalogRepoName,
		Canonical:           &catalogRepo,
		URL:                 &catalogRepoURL,
		Branch:              catalogRepoBranch,
		CredentialCanonical: gitCred,
	}

	apiURL, ok := os.LookupEnv("CY_API_URL")
	if !ok {
		apiURL = "https://api-cli-test.staging.cycloid.io/"
	}
	config.APIUrl = apiURL

	org, ok = os.LookupEnv("CY_TEST_ROOT_ORG")
	if !ok {
		org = "cycloid"
	}
	config.Org = org

	apiKey, ok = os.LookupEnv("CY_TEST_API_KEY")
	if !ok {
		return config, fmt.Errorf("api key not set in CY_TEST_API_KEY env var")
	}
	config.APIKey = apiKey

	api := common.NewAPI(
		common.WithURL(apiURL),
		common.WithInsecure(true),
		common.WithToken(apiKey),
	)
	m := middleware.NewMiddleware(api)
	config.Middleware = m

	project, err := config.NewTestProject("common")
	if err != nil {
		return config, err
	}
	config.Project = project

	environment, err := config.NewTestEnv("common", *project.Canonical)
	if err != nil {
		return config, err
	}
	config.Environment = environment

	stackRef := org + ":" + defaultStackCanonical
	stackConfig, err := m.GetStackConfig(org, stackRef)
	if err != nil {
		return config, err
	}

	vars, err := common.FormUseCaseToFormVars(stackConfig, defaultStackUseCase)
	if err != nil {
		return config, err
	}

	// Add a random value in forms to avoid git conflict
	if vars != nil {
		common.UpdateMapField("types.tests.string", RandomCanonical("common"), *vars)
	} else {
		vars = &models.FormVariables{
			"types": {"tests": {"string": RandomCanonical("common")}},
		}
	}

	component, err := config.NewTestComponent(
		*project.Canonical, *environment.Canonical, "common", stackRef, defaultStackUseCase, vars,
	)
	if err != nil {
		return config, err
	}
	config.Component = component

	return config, nil
}

// NewTestProject will create a project with a random canonical derived from identifier
// and return the project, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *Config) NewTestProject(identifier string) (*models.Project, error) {
	var (
		project          = RandomCanonical(identifier)
		description      = "Testing project " + identifier
		configRepository = *config.ConfigRepo.Canonical
		owner            = ""
		team             = ""
		color            = "default"
		icon             = "inventory"
	)

	m := config.Middleware

	out, err := m.CreateProject(config.Org, project, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		return nil, fmt.Errorf("failed to setup test project: %s", err)
	}

	config.AppendCleanup(func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			log.Fatalf("cannot cleanup projet '%s' for test '%s': %s", project, identifier, err)
			return
		}
	})

	return out, nil
}

// setupTestProject will create an env with a random canonical derived from identifier
// and return the env, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *Config) NewTestEnv(identifier, project string) (*models.Environment, error) {
	var (
		env   = RandomCanonical(identifier)
		color = "default"
	)

	m := config.Middleware

	out, err := m.CreateEnv(config.Org, project, env, env, color)
	if err != nil {
		return nil, fmt.Errorf("failed to setup test environment: %s", err)
	}
	config.AppendCleanup(func() {
		err := m.DeleteEnv(config.Org, project, env)
		if err != nil {
			log.Fatalf("cannot cleanup env '%s' for test '%s': %s", env, identifier, err)
			return
		}
	})

	return out, nil
}

// setupTestProject will create an component with a random canonical derived from identifier
// and return the component, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *Config) NewTestComponent(project, env, identifier, stackRef, useCase string, inputs *models.FormVariables) (*models.Component, error) {
	m := config.Middleware
	component := RandomCanonical(identifier)

	createdComponent, err := m.CreateComponent(
		config.Org, project, env, component, "", &component, &stackRef, &useCase, nil, inputs,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to setup component '%s' for test '%s':\n%v", component, identifier, err)
	}
	config.AppendCleanup(func() {
		if err := m.DeleteComponent(config.Org, project, env, component); err != nil {
			log.Fatalf("failed to cleanup component for test '%s': %s", identifier, err)
		}
	})

	return createdComponent, nil
}

func (config *Config) AppendCleanup(f ...func()) {
	config.cleanupFuncs = append(config.cleanupFuncs, f...)
}

func (config *Config) Cleanup() {
	for _, f := range config.cleanupFuncs {
		defer f()
	}
}

// NewTestChildOrg is a helper function that create a suborg for a
// specific test and return a function to defer() for its deletion.
func (config *Config) NewTestChildOrg(parent, child string) (func(), error) {
	m := config.Middleware
	deferFunc := func() {
		err := m.DeleteOrganization(child)
		if err != nil {
			log.Fatalf("Failed to delete org '%s': %v", child, err)
			return
		}
	}

	_, err := m.CreateOrganizationChild(parent, child, nil)
	if err != nil {
		return deferFunc, fmt.Errorf("failed to create child org '%s': %v", child, err)
	}

	return deferFunc, nil
}
