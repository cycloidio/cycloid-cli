package middleware_test

import (
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
	"github.com/stretchr/testify/assert"
)

func TestProjectCrud(t *testing.T) {
	t.Parallel()
	m := config.Middleware

	var (
		projectName = "Test CRUD Projects"
		project     = testcfg.RandomCanonical("test-crud-project")
		description = "My cool project description !\nWith a nexline"
		owner       = ""
		team        = ""
		color       = "default"
		icon        = "world"
	)

	createProjet, err := m.CreateProject(config.Org, projectName, project, description, *config.ConfigRepo.Canonical, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create project '%s': %v", project, err)
	}

	_, err = m.GetProject(config.Org, *createProjet.Canonical)
	if err != nil {
		t.Fatalf("Did not found create project '%s' with get request: %v", *createProjet.Canonical, err)
	}

	defer func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			t.Fatalf("Failed to delete project '%s': %v", project, err)
		}
	}()

	var (
		newIcon        = "world"
		newName        = "My cool new name"
		newDescription = "Updated description"
		newColor       = "red"
	)

	updatedProject, err := m.UpdateProject(
		config.Org, newName, project, newDescription, *config.ConfigRepo.Canonical,
		owner, team, newColor, newIcon, "aws", createProjet.UpdatedAt,
	)
	if err != nil {
		t.Fatalf("Failed to update project '%s': %v", project, err)
	}

	assert.Equal(t, newName, *updatedProject.Name)
	assert.Equal(t, newDescription, updatedProject.Description)
	assert.Equal(t, newIcon, *updatedProject.Icon)
	assert.Equal(t, newColor, *updatedProject.Color)
}
