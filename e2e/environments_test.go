package e2e_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

func TestEnvs(t *testing.T) {
	os.Setenv("CY_API_URL", config.APIUrl)
	os.Setenv("CY_API_KEY", config.APIKey)
	os.Setenv("CY_ORG", config.Org)

	var (
		project = *config.Project.Canonical
		envName = "Test E2E env"
		env     = randomCanonical("e2e")
		color   = "demo"
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
			t.Errorf("failed to create env '%s': %v", env, err)
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
			t.Errorf("failed to delete env '%s': %v", env, err)
		}
	})

	t.Run("Get", func(t *testing.T) {
		args := []string{
			"env", "get",
			"--project", project,
			"--env", env,
			"--output", "json",
		}
		cmdOut, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to get env '%s': %v", env, err)
		}

		var got models.Environment
		err = json.Unmarshal([]byte(cmdOut), &got)
		if err != nil {
			t.Errorf("failed to marshal json cli output: %s\nerr: %s", cmdOut, err)
		}

		assert.EqualValues(t, env, *got.Canonical)
	})

	t.Run("CreateWithUpdate", func(t *testing.T) {
		var (
			createUpdateName = "helloUpdate"
			newColor         = cyargs.PickRandomColor(&env)
		)
		args := []string{
			"-o", "json",
			"env", "create",
			"--env", env,
			"--project", project,
			"--name", createUpdateName,
			"--color", newColor,
			"--update",
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create env '%s': %v", env, err)
		}

		var envResult models.Environment
		err = json.Unmarshal([]byte(out), &envResult)
		if err != nil {
			t.Errorf("failed to parse json output from the CLI on create --update: %v\noutput: %s", err, out)
		}

		assert.Equal(t, createUpdateName, envResult.Name)
		assert.Equal(t, newColor, *envResult.Color)
	})

	t.Run("CreateWithUpdateNew", func(t *testing.T) {
		var newEnv = randomCanonical("e2e-env")
		args := []string{
			"-o", "json",
			"env", "create",
			"--project", project,
			"--env", newEnv,
			"--color", color,
			"--update",
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create env '%s': %v", newEnv, err)
		}

		defer t.Run("DeleteCreateUpdate", func(t *testing.T) {
			args := []string{
				"env", "delete",
				"--project", project,
				"--env", newEnv,
			}
			_, err := executeCommand(args)
			if err != nil {
				t.Errorf("failed to delete env '%s': %v", newEnv, err)
			}
		})
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
			t.Errorf("failed to update env '%s': %v", env, err)
		}

		var envResult models.Environment
		err = json.Unmarshal([]byte(out), &envResult)
		if err != nil {
			t.Errorf("failed to parse json output from the CLI on update: %v", err)
		}

		assert.Equal(t, newName, envResult.Name)
	})
}
