package e2e

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/stretchr/testify/assert"
)

func TestEnvs(t *testing.T) {
	t.Parallel()

	os.Setenv("CY_API_URL", TestAPIURL)
	os.Setenv("CY_API_KEY", TestAPIKey)
	os.Setenv("CY_ORG", TestRootOrg)

	// setup
	api := common.NewAPI(
		common.WithInsecure(true),
		common.WithURL(TestAPIURL),
		common.WithToken(TestAPIKey),
	)
	m := middleware.NewMiddleware(api)

	var (
		projectName        = "Test E2E env"
		project            = randomCanonical("test-e2e-env")
		projectDescription = "Testing envs"
		configRepository   = CyTestConfigRepo
		projectColor       = "blue"
		projectIcon        = "planet"
	)

	defer func() {
		err := m.DeleteProject(TestRootOrg, project)
		if err != nil {
			t.Fatalf("Failed to cleanup project '%s' for test '%s': %v", project, t.Name(), err)
		}
	}()

	_, err := m.CreateProject(TestRootOrg, projectName, project, projectDescription, configRepository, "", "", projectColor, projectIcon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project '%s' for test '%s': %v", project, t.Name(), err)
	}
	// env setup

	var (
		envName = "Test E2E env"
		env     = randomCanonical("e2e")
		color   = "blue"
	)

	t.Run("Create", func(t *testing.T) {
		args := []string{
			"env", "create",
			"--project", project,
			"--env", env,
			"--name", envName,
			"--color", color,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to create env '%s': %v", env, err)
		}
	})

	defer t.Run("Delete", func(t *testing.T) {
		args := []string{
			"env", "delete",
			"-p", project,
			"-e", env,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to delete env '%s': %v", env, err)
		}
	})

	t.Run("CreateWithUpdate", func(t *testing.T) {
		var createUpdateName = "helloUpdate"
		args := []string{
			"-o", "json",
			"env", "create",
			"--env", env,
			"--project", project,
			"--name", createUpdateName,
			"--color", projectColor,
			"--update",
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to create env '%s': %v", env, err)
		}

		var envResult models.Environment
		err = json.Unmarshal([]byte(out), &envResult)
		if err != nil {
			t.Fatalf("failed to parse json output from the CLI on create --update: %v\noutput: %s", err, out)
		}

		assert.Equal(t, createUpdateName, envResult.Name)
	})

	t.Run("Update", func(t *testing.T) {
		var newName = "NewName"
		args := []string{
			"-o", "json",
			"env", "update",
			"--name", newName,
			"--project", project,
			"--env", env,
			"--color", color,
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to update env '%s': %v", env, err)
		}

		var envResult models.Environment
		err = json.Unmarshal([]byte(out), &envResult)
		if err != nil {
			t.Fatalf("failed to parse json output from the CLI on update: %v", err)
		}

		assert.Equal(t, newName, envResult.Name)
	})
}
