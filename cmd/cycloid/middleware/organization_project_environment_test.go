package middleware_test

import (
	"log"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
	"github.com/stretchr/testify/assert"
)

func TestEnvCrud(t *testing.T) {
	// setup
	m := config.Middleware

	var (
		projectName      = "Test CRUD ENV"
		project          = testcfg.RandomCanonical("test-crud-env")
		description      = "Testing env"
		configRepository = *config.ConfigRepo.Canonical
		owner            = ""
		team             = ""
		color            = "blue"
		icon             = "planet"
	)

	defer func() {
		err := m.DeleteProject(config.Org, project)
		if err != nil {
			log.Fatalf("Failed to decomission project '%s' from CRUD tests: %v", project, err)
			return
		}
	}()

	createdProject, err := m.CreateProject(config.Org, projectName, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project, create project CRUD tests: %v", err)
	}

	// end setup
	var (
		env      = "test-crud"
		envName  = "TestCRUD"
		envColor = "red"
	)

	// create
	createdEnv, err := m.CreateEnv(config.Org, *createdProject.Canonical, env, envName, envColor)
	if err != nil {
		t.Fatalf("Failed to create env '%s': %v", env, err)
	}

	// delete
	defer func() {
		err := m.DeleteEnv(config.Org, *createdProject.Canonical, *createdEnv.Canonical)
		if err != nil {
			log.Fatalf("Failed to delete env '%s': %v", env, err)
			return
		}
	}()

	// update
	var (
		newEnvName  = "New"
		newEnvColor = "blue"
	)

	updatedEnv, err := m.UpdateEnv(config.Org, *createdProject.Canonical, *createdEnv.Canonical, newEnvName, newEnvColor)
	if err != nil {
		t.Fatalf("Failed to update env '%s':\n%v", env, err)
	}

	assert.Equal(t, newEnvName, updatedEnv.Name)
	assert.Equal(t, newEnvColor, *updatedEnv.Color)
}
