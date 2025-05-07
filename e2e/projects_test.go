// //go:build e2e
// // +build e2e
package e2e

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/stretchr/testify/assert"
)

func TestProjects(t *testing.T) {
	t.Parallel()

	os.Setenv("CY_API_URL", CY_API_URL)
	os.Setenv("CY_API_KEY", CY_TEST_API_KEY)
	os.Setenv("CY_ORG", CY_TEST_ROOT_ORG)

	var (
		projectName = "Test E2E project"
		project     = "test-e2e-project"
		description = "Testing project"
		owner       = ""
		color       = "blue"
		icon        = "planet"
	)

	t.Run("Create", func(t *testing.T) {
		args := []string{
			"project", "create",
			"--project", project,
			"--name", projectName,
			"--description", description,
			"--owner", owner,
			"--icon", icon,
			"--color", color,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to create project '%s': %v", project, err)
		}
	})

	defer t.Run("Delete", func(t *testing.T) {
		args := []string{
			"project", "delete",
			"--project", project,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to delete project '%s': %v", project, err)
		}
	})

	t.Run("CreateWithUpdate", func(t *testing.T) {
		var createUpdateName = "helloUpdate"
		args := []string{
			"-o", "json",
			"project", "create",
			"--project", project,
			"--name", createUpdateName,
			"--description", description,
			"--owner", owner,
			"--icon", icon,
			"--color", color,
			"--update",
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to create project '%s': %v", project, err)
		}

		var projectResult models.Project
		err = json.Unmarshal([]byte(out), &projectResult)
		if err != nil {
			t.Fatalf("failed to parse json output from the CLI on create --update: %v\noutput: %s", err, out)
		}

		assert.Equal(t, createUpdateName, *projectResult.Name)
	})

	t.Run("Update", func(t *testing.T) {
		var updateDesc = "NewDesc"
		args := []string{
			"-o", "json",
			"project", "update",
			"--project", project,
			"--name", projectName,
			"--description", updateDesc,
			"--owner", owner,
			"--icon", icon,
			"--color", color,
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to update project '%s': %v", project, err)
		}

		var projectResult models.Project
		err = json.Unmarshal([]byte(out), &projectResult)
		if err != nil {
			t.Fatalf("failed to parse json output from the CLI on update: %v", err)
		}

		assert.Equal(t, updateDesc, projectResult.Description)
	})
}
