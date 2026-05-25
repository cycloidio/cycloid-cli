package e2e_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestEnvs(t *testing.T) {
	os.Setenv("CY_API_URL", config.APIUrl)
	os.Setenv("CY_API_KEY", config.APIKey)
	os.Setenv("CY_ORG", config.Org)

	var (
		project = *config.Project.Canonical
		envName = "Test E2E env"
		env     = randomCanonical("e2e")
	)

	t.Run("Create", func(t *testing.T) {
		args := []string{
			"env", "create",
			"--env", env,
			"--name", envName,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create env %q: %v", env, err)
		}
	})

	t.Run("Link", func(t *testing.T) {
		args := []string{
			"env", "link",
			"--project", project,
			"--env", env,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to link env %q: %v", env, err)
		}
	})

	t.Run("CreateWithNameOnly", func(t *testing.T) {
		var (
			nameOnly = "Environment Name " + randomCanonical("e2e")
			inferred = middleware.ToCanonical(nameOnly)
		)

		args := []string{
			"--output", "json",
			"env", "create",
			"--name", nameOnly,
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create env with name-only %q: %v", nameOnly, err)
		}

		defer t.Run("DeleteNameOnlyEnv", func(t *testing.T) {
			args := []string{
				"env", "delete",
				"--env", inferred,
			}
			_, err := executeCommand(args)
			if err != nil {
				t.Errorf("failed to delete inferred env %q: %v", inferred, err)
			}
		})

		var createdEnv models.Environment
		err = json.Unmarshal([]byte(out), &createdEnv)
		if err != nil {
			t.Errorf("failed to parse json output for inferred env creation: %v\noutput: %s", err, out)
		}

		assert.Equal(t, inferred, *createdEnv.Canonical)
		assert.Equal(t, nameOnly, createdEnv.Name)
	})

	defer t.Run("Delete", func(t *testing.T) {
		unlinkArgs := []string{
			"env", "unlink",
			"-p", project,
			"-e", env,
		}
		if _, err := executeCommand(unlinkArgs); err != nil {
			t.Errorf("failed to unlink env %q: %v", env, err)
		}

		args := []string{
			"env", "delete",
			"-e", env,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to delete env %q: %v", env, err)
		}
	})

	t.Run("Get", func(t *testing.T) {
		args := []string{
			"env", "get",
			"--env", env,
			"--output", "json",
		}
		cmdOut, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to get env %q: %v", env, err)
		}

		var got models.Environment
		err = json.Unmarshal([]byte(cmdOut), &got)
		if err != nil {
			t.Errorf("failed to marshal json cli output: %s\nerr: %s", cmdOut, err)
		}

		assert.EqualValues(t, env, *got.Canonical)
	})

	t.Run("CreateWithUpdate", func(t *testing.T) {
		createUpdateName := "helloUpdate"
		args := []string{
			"-o", "json",
			"env", "create",
			"--env", env,
			"--name", createUpdateName,
			"--update",
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create env %q: %v", env, err)
		}

		var envResult models.Environment
		err = json.Unmarshal([]byte(out), &envResult)
		if err != nil {
			t.Errorf("failed to parse json output from the CLI on create --update: %v\noutput: %s", err, out)
		}

		assert.Equal(t, createUpdateName, envResult.Name)
	})

	t.Run("CreateWithUpdateNew", func(t *testing.T) {
		newEnv := randomCanonical("e2e-env")
		args := []string{
			"-o", "json",
			"env", "create",
			"--env", newEnv,
			"--update",
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create env %q: %v", newEnv, err)
		}

		defer t.Run("DeleteCreateUpdate", func(t *testing.T) {
			args := []string{
				"env", "delete",
				"--env", newEnv,
			}
			_, err := executeCommand(args)
			if err != nil {
				t.Errorf("failed to delete env %q: %v", newEnv, err)
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		newName := "NewName"
		args := []string{
			"-o", "json",
			"env", "update",
			"--name", newName,
			"--env", env,
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to update env %q: %v", env, err)
		}

		var envResult models.Environment
		err = json.Unmarshal([]byte(out), &envResult)
		if err != nil {
			t.Errorf("failed to parse json output from the CLI on update: %v", err)
		}

		assert.Equal(t, newName, envResult.Name)
	})
}
