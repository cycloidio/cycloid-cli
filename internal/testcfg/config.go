package testcfg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

func NewConfig(testName string) (*Config, error) {
	if len(testName) < 1 {
		return nil, fmt.Errorf("testName argument must not be empty")
	}

	var (
		localGitSSHKey = strings.Join([]string{
			"-----BEGIN OPENSSH PRIVATE KEY-----",
			"b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW",
			"QyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+AAAAJjCF9jzwhfY",
			"8wAAAAtzc2gtZWQyNTUxOQAAACD8O9mhkl5CAiD0NLeQcoPf1duYHImQlTjOXcCOgHmC+A",
			"AAAEC0ryBZ1uJQ2drmjsO+WpsC2E/5SWheJD/r8+Q4LghWxfw72aGSXkICIPQ0t5Byg9/V",
			"25gciZCVOM5dwI6AeYL4AAAAE2N5Y2xvaWRAZXhhbXBsZS5jb20BAg==",
			"-----END OPENSSH PRIVATE KEY-----",
		}, "\n")
		localGitCredential    = "local-git"
		dockerIPAM            = EnvDefault("DOCKER_IPAM", "192.168.10")
		configRepoName        = "cli-test-config"
		configRepository      = "cli-test-config"
		configRepoURL         = fmt.Sprintf("git@%s.7:/git-server/repos/backend-test-config-repo.git", dockerIPAM)
		configRepoBranch      = "master"
		isDefault             = true
		gitCred               = "github"
		catalogRepo           = "cli-test-stacks"
		catalogRepoName       = "CLI test catalog"
		catalogRepoURL        = "https://github.com/cycloidio/cycloid-cli-test-catalog.git"
		catalogRepoBranch     = "stacks"
		defaultStackCanonical = "stack-e2e-stackforms"
		defaultStackUseCase   = "default"
	)

	var config = &Config{}
	config.ConfigRepo = &models.ConfigRepository{
		Name:                &configRepoName,
		Canonical:           &configRepository,
		URL:                 &configRepoURL,
		Branch:              configRepoBranch,
		CredentialCanonical: gitCred,
		Default:             &isDefault,
	}
	config.CatalogRepo = &models.ServiceCatalogSource{
		Name:                &catalogRepoName,
		Canonical:           &catalogRepo,
		URL:                 &catalogRepoURL,
		Branch:              catalogRepoBranch,
		CredentialCanonical: gitCred,
	}

	provisionAPI, _ := strconv.ParseBool(EnvDefault("CY_TEST_PROVISION_API", "1"))
	config.APIUrl = EnvDefault("CY_TEST_API_URL", "http://"+dockerIPAM+".10:3001")
	config.Org = EnvDefault("CY_TEST_ROOT_ORG", "cycloid")
	licence, ok := os.LookupEnv("API_LICENCE_KEY")
	if !ok && provisionAPI {
		return config, fmt.Errorf("licence required for provisionning, set it with API_LICENCE_KEY")
	}

	// If we provision the api, we will try to login first
	if !provisionAPI {
		apiKey, ok := os.LookupEnv("CY_TEST_API_KEY")
		if !ok {
			return config, fmt.Errorf("api key not set in CY_TEST_API_KEY env var")
		}
		config.APIKey = apiKey
	}

	api := common.NewAPI(
		common.WithURL(config.APIUrl),
		common.WithInsecure(true),
		common.WithToken(config.APIKey),
	)
	m := middleware.NewMiddleware(api)
	config.Middleware = m

	var (
		userName        = "administrator"
		email           = "admin@cycloid.io"
		password        = "cycloidadmin"
		apiKeyCanonical = "admin-" + testName
	)

	if provisionAPI {
		// try to login, is successful, console is initialized
		init, err := m.InitFirstOrg(config.Org, userName, userName, userName, email, password, licence, &apiKeyCanonical)
		if err != nil {
			return nil, fmt.Errorf("failed to init console: %w", err)
		}

		config.APIKey = *init.APIKey
		api.Config.Token = *init.APIKey

		// Write the API for the User, we'll look up a better way later
		root, err := FindRepoRoot()
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(root+"/.api_key", []byte("CY_API_KEY="+config.APIKey), 0666)
		if err != nil {
			return nil, err
		}
	}

	_, err := m.CreateCredential(config.Org, localGitCredential, "ssh",
		&models.CredentialRaw{SSHKey: localGitSSHKey}, "", localGitCredential, "",
	)
	var apiErr *middleware.APIError
	if errors.As(err, &apiErr) {
		if apiErr.HTTPCode != "409" {
			return config, fmt.Errorf("failed to init config repo credential: %w", err)
		}
	}

	currentConfigRepo, err := m.CreateConfigRepository(config.Org,
		*config.ConfigRepo.Canonical, *config.ConfigRepo.Canonical, *config.ConfigRepo.URL,
		config.ConfigRepo.Branch, localGitCredential, *config.ConfigRepo.Default,
	)
	if errors.As(err, &apiErr) {
		var getErr error
		currentConfigRepo, getErr = m.GetConfigRepository(config.Org, configRepository)
		if apiErr.HTTPCode != "409" || getErr != nil {
			return config, fmt.Errorf("failed to setup config repo: %w%w", err, getErr)
		}
	}

	config.ConfigRepo = currentConfigRepo

	_, err = m.CreateCatalogRepository(config.Org, *config.CatalogRepo.Canonical,
		*config.CatalogRepo.URL, config.CatalogRepo.Branch, "", "local", "",
	)
	if errors.As(err, &apiErr) {
		if apiErr.HTTPCode != "409" {
			return config, fmt.Errorf("failed to setup catalog repo: %w", err)
		}
	}

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

	stackRef := config.Org + ":" + defaultStackCanonical
	component, err := config.NewTestComponent(
		*project.Canonical, *environment.Canonical, "common", stackRef, defaultStackUseCase, nil,
	)
	if err != nil {
		return config, err
	}
	config.Component = component

	stackConfig, err := m.GetComponentStackConfig(config.Org, *project.Canonical, *environment.Canonical, *component.Canonical, defaultStackUseCase)
	if err != nil {
		return config, err
	}

	vars, err := common.FormUseCaseToFormVars(stackConfig, defaultStackUseCase)
	if err != nil {
		return config, err
	}

	// Add a random value in forms to avoid git conflict
	if vars != nil {
		common.UpdateMapField("types.tests.string", RandomCanonical("common"), vars)
	} else {
		vars = models.FormVariables{
			"types": {"tests": {"string": RandomCanonical("common")}},
		}
	}

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
		return nil, fmt.Errorf("failed to setup test project: %w", err)
	}

	config.AppendCleanup(func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			log.Fatalf("cannot cleanup projet %q for test %q: %v", project, identifier, err)
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
		return nil, fmt.Errorf("failed to setup test environment: %w", err)
	}

	config.AppendCleanup(func() {
		err := m.DeleteEnv(config.Org, project, env)
		if err != nil {
			log.Fatalf("cannot cleanup env %q for test %q: %v", env, identifier, err)
			return
		}
	})

	return out, nil
}

// setupTestProject will create an component with a random canonical derived from identifier
// and return the component, the function to defer for its deletion and error.
// The func will always be returned so even if err != nil, defer the func.
func (config *Config) NewTestComponent(project, env, identifier, stackRef, useCase string, inputs models.FormVariables) (*models.Component, error) {
	m := config.Middleware
	component := RandomCanonical(identifier)

	outComponent, err := m.CreateAndConfigureComponent(
		config.Org, project, env, component, "", &component, stackRef, useCase, "", inputs,
	)
	if err != nil {
		return nil, err
	}

	config.AppendCleanup(func() {
		if err := m.DeleteComponent(config.Org, project, env, component); err != nil {
			log.Printf("failed to cleanup component for test %q: %v", identifier, err)
		}
	})

	return outComponent, nil
}

func (config *Config) AppendCleanup(f ...func()) {
	config.cleanupFuncs = append(config.cleanupFuncs, f...)
}

func (config *Config) Cleanup() {
	if config == nil {
		return
	}

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
			log.Fatalf("Failed to delete org %q: %v", child, err)
			return
		}
	}

	_, err := m.CreateOrganizationChild(parent, child, nil)
	if err != nil {
		return deferFunc, fmt.Errorf("failed to create child org %q: %v", child, err)
	}

	return deferFunc, nil
}
