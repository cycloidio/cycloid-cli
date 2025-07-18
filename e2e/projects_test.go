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

	os.Setenv("CY_API_URL", TestAPIURL)
	os.Setenv("CY_API_KEY", TestAPIKey)
	os.Setenv("CY_ORG", TestRootOrg)

	var (
		projectName = "Test E2E project"
		project     = randomCanonical("test-e2e-project")
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

	t.Run("CreateWithUpdateExisting", func(t *testing.T) {
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

	t.Run("CreateWithUpdateNew", func(t *testing.T) {
		var newProject = randomCanonical("e2e-new")
		args := []string{
			"-o", "json",
			"project", "create",
			"--project", newProject,
			"--description", description,
			"--owner", owner,
			"--icon", icon,
			"--color", color,
			"--update",
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to create project '%s': %v", newProject, err)
		}

		defer t.Run("DeleteCreateUpdate", func(t *testing.T) {
			args := []string{
				"project", "delete",
				"--project", newProject,
			}
			_, err := executeCommand(args)
			if err != nil {
				t.Fatalf("failed to delete project '%s': %v", newProject, err)
			}
		})
	})

	t.Run("Update", func(t *testing.T) {
		var updateDesc = "NewDesc"
		args := []string{
			"project", "update",
			"--output", "json",
			"--project", project,
			"--name", projectName,
			"--description", updateDesc,
			"--owner", owner,
			"--icon", icon,
			"--color", color,
		}
		jsonOut, err := executeCommand(args)
		if err != nil {
			t.Fatalf("failed to update project '%s': %v", project, err)
		}

		var projectResult models.Project
		err = json.Unmarshal([]byte(jsonOut), &projectResult)
		if err != nil {
			t.Fatalf("failed to parse json output from the CLI on update: %v\noutput: %s", err, jsonOut)
		}

		assert.Equal(t, updateDesc, projectResult.Description)
	})
}
