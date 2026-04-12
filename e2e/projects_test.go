package e2e_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

func TestProjects(t *testing.T) {

	os.Setenv("CY_API_URL", config.APIUrl)
	os.Setenv("CY_API_KEY", config.APIKey)
	os.Setenv("CY_ORG", config.Org)

	var (
		projectName = "Test E2E project"
		project     = randomCanonical("test-e2e-project")
		description = "Testing project"
		owner       = ""
		color       = "demo"
		icon        = "public"
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
			t.Errorf("failed to create project '%s': %v", project, err)
		}
	})

	t.Run("CreateWithNameOnlyAndUpdate", func(t *testing.T) {
		var (
			nameOnly    = "Project Name " + randomCanonical("e2e")
			canonical   = middleware.ToCanonical(nameOnly)
			updatedDesc = "Updated inferred project description"
			newIcon     = cyargs.PickRandomIcon(nil)
			newColor    = cyargs.PickRandomColor(nil)
		)

		createArgs := []string{
			"--output", "json",
			"project", "create",
			"--name", nameOnly,
			"--description", description,
			"--owner", owner,
			"--icon", icon,
			"--color", color,
		}
		createOut, err := executeCommand(createArgs)
		if err != nil {
			t.Errorf("failed to create project with name-only '%s': %v", nameOnly, err)
		}

		defer t.Run("DeleteNameOnlyProject", func(t *testing.T) {
			args := []string{
				"project", "delete",
				"--project", canonical,
			}
			_, err := executeCommand(args)
			if err != nil {
				t.Errorf("failed to delete project '%s': %v", canonical, err)
			}
		})

		var createdProject models.Project
		err = json.Unmarshal([]byte(createOut), &createdProject)
		if err != nil {
			t.Errorf("failed to parse create output for inferred project: %v\noutput: %s", err, createOut)
		}
		assert.Equal(t, nameOnly, *createdProject.Name)
		assert.Equal(t, canonical, *createdProject.Canonical)

		updateArgs := []string{
			"--output", "json",
			"project", "create",
			"--name", nameOnly,
			"--description", updatedDesc,
			"--owner", owner,
			"--icon", newIcon,
			"--color", newColor,
			"--update",
		}
		updateOut, err := executeCommand(updateArgs)
		if err != nil {
			t.Errorf("failed to create --update project with name-only '%s': %v", nameOnly, err)
		}

		var updatedProject models.Project
		err = json.Unmarshal([]byte(updateOut), &updatedProject)
		if err != nil {
			t.Errorf("failed to parse update output for inferred project: %v\noutput: %s", err, updateOut)
		}
		assert.Equal(t, canonical, *updatedProject.Canonical)
		assert.Equal(t, nameOnly, *updatedProject.Name)
		assert.Equal(t, updatedDesc, updatedProject.Description)
		assert.Equal(t, newColor, *updatedProject.Color)
		assert.Equal(t, newIcon, *updatedProject.Icon)
	})

	defer t.Run("Delete", func(t *testing.T) {
		args := []string{
			"project", "delete",
			"--project", project,
		}
		_, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to delete project '%s': %v", project, err)
		}
	})

	t.Run("CreateWithUpdateExisting", func(t *testing.T) {
		var (
			createUpdateName = "helloUpdate"
			newIcon          = cyargs.PickRandomIcon(nil)
			newColor         = cyargs.PickRandomColor(nil)
		)
		args := []string{
			"-o", "json",
			"project", "create",
			"--project", project,
			"--name", createUpdateName,
			"--description", description,
			"--owner", owner,
			"--icon", newIcon,
			"--color", newColor,
			"--update",
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to create project '%s': %v", project, err)
		}

		var projectResult models.Project
		err = json.Unmarshal([]byte(out), &projectResult)
		if err != nil {
			t.Errorf("failed to parse json output from the CLI on create --update: %v\noutput: %s", err, out)
		}

		assert.Equal(t, createUpdateName, *projectResult.Name)
		assert.Equal(t, newColor, *projectResult.Color)
		assert.Equal(t, newIcon, *projectResult.Icon)
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
			t.Errorf("failed to create project '%s': %v", newProject, err)
		}

		defer t.Run("DeleteCreateUpdate", func(t *testing.T) {
			args := []string{
				"project", "delete",
				"--project", newProject,
			}
			_, err := executeCommand(args)
			if err != nil {
				t.Errorf("failed to delete project '%s': %v", newProject, err)
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
			t.Errorf("failed to update project '%s': %v", project, err)
		}

		var projectResult models.Project
		err = json.Unmarshal([]byte(jsonOut), &projectResult)
		if err != nil {
			t.Errorf("failed to parse json output from the CLI on update: %v\noutput: %s", err, jsonOut)
		}

		assert.Equal(t, updateDesc, projectResult.Description)
	})

	t.Run("ListEnv", func(t *testing.T) {
		args := []string{
			"--output", "json",
			"project", "list-env",
			"--project", project,
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to list project environments for '%s': %v", project, err)
		}

		var envList []*models.Environment
		if jsonErr := json.Unmarshal([]byte(out), &envList); jsonErr != nil {
			t.Errorf("failed to parse json output of list-env: %v\noutput: %s", jsonErr, out)
		}
	})

	// Positional-arg tests — verify the new interface works alongside deprecated --project flag

	t.Run("GetWithPositionalArg", func(t *testing.T) {
		args := []string{
			"--output", "json",
			"project", "get",
			project, // positional, no --project flag
		}
		out, err := executeCommand(args)
		if err != nil {
			t.Errorf("failed to get project '%s' via positional arg: %v", project, err)
		}

		var got models.Project
		if jsonErr := json.Unmarshal([]byte(out), &got); jsonErr != nil {
			t.Errorf("failed to parse get output: %v\noutput: %s", jsonErr, out)
		}
		assert.Equal(t, project, *got.Canonical)
	})

	t.Run("DeleteWithPositionalArg", func(t *testing.T) {
		// Create a dedicated project to delete via positional arg
		tmpProject := randomCanonical("e2e-del-pos")
		_, createErr := executeCommand([]string{
			"--output", "json",
			"project", "create",
			"--project", tmpProject,
			"--name", tmpProject,
			"--icon", icon,
			"--color", color,
		})
		if createErr != nil {
			t.Errorf("setup: failed to create project '%s': %v", tmpProject, createErr)
		}

		_, deleteErr := executeCommand([]string{
			"--output", "json",
			"project", "delete",
			tmpProject, // positional, no --project flag
		})
		assert.NoError(t, deleteErr, "delete via positional arg should succeed")
	})

	t.Run("DeleteDefaultTableOutput", func(t *testing.T) {
		// Verify delete uses table mode by default and shows the deleted canonical.
		tmpProject := randomCanonical("e2e-del-tbl")
		_, createErr := executeCommand([]string{
			"project", "create",
			"--project", tmpProject,
			"--name", tmpProject,
			"--icon", icon,
			"--color", color,
		})
		if createErr != nil {
			t.Fatalf("setup: failed to create project '%s': %v", tmpProject, createErr)
		}

		out, deleteErr := executeCommand([]string{
			"project", "delete",
			tmpProject,
		})
		assert.NoError(t, deleteErr, "delete should succeed")
		assert.Contains(t, out, tmpProject, "default table output should contain the deleted canonical")
		assert.NotContains(t, out, "[", "default output should not be JSON array")
	})

	t.Run("DeleteMultiplePositionalArgs", func(t *testing.T) {
		// Verify deleting multiple projects at once outputs a table with all canonicals.
		proj1 := randomCanonical("e2e-del-m1")
		proj2 := randomCanonical("e2e-del-m2")
		for _, p := range []string{proj1, proj2} {
			_, createErr := executeCommand([]string{
				"project", "create",
				"--project", p,
				"--name", p,
				"--icon", icon,
				"--color", color,
			})
			if createErr != nil {
				t.Fatalf("setup: failed to create project '%s': %v", p, createErr)
			}
		}

		out, deleteErr := executeCommand([]string{
			"project", "delete",
			proj1, proj2,
		})
		assert.NoError(t, deleteErr, "bulk delete should succeed")
		assert.Contains(t, out, proj1, "output should contain first deleted canonical")
		assert.Contains(t, out, proj2, "output should contain second deleted canonical")
		assert.NotContains(t, out, "[", "default output should not be JSON array")
	})
}
