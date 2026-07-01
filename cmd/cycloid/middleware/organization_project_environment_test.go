package middleware_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/pkg/testcfg"
)

func TestEnvCrud(t *testing.T) {
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
		_, err := m.DeleteProject(config.Org, project, middleware.DeleteOptions{})
		if err != nil {
			log.Fatalf("Failed to decommission project '%s' from CRUD tests: %v", project, err)
		}
	}()

	createdProject, _, err := m.CreateProject(config.Org, projectName, project, description, configRepository, owner, team, color, icon)
	if err != nil {
		t.Fatalf("Failed to create pre-requisite project: %v", err)
	}

	var (
		env     = testcfg.RandomCanonical("test-crud")
		envName = "TestCRUD"
	)

	createBody := &models.NewEnvironment{
		Canonical: env,
		Name:      ptr.Ptr(envName),
		Type:      "production",
	}
	createdEnv, _, err := m.CreateOrgEnv(config.Org, createBody)
	if err != nil {
		t.Fatalf("Failed to create env %q: %v", env, err)
	}

	defer func() {
		if createdEnv == nil {
			return
		}
		// Org-level envs cannot be deleted while still linked to a project
		// (backend returns 409). Unlink first so cleanup is best-effort even
		// when the test panics mid-flight.
		if _, err := m.UnlinkEnvFromProject(config.Org, *createdProject.Canonical, *createdEnv.Canonical, middleware.DeleteOptions{}); err != nil {
			log.Printf("Warning: failed to unlink env %q from project %q during cleanup: %v", env, *createdProject.Canonical, err)
		}
		if _, err := m.DeleteOrgEnv(config.Org, *createdEnv.Canonical); err != nil {
			log.Fatalf("Failed to delete org env %q: %v", env, err)
		}
	}()

	_, err = m.LinkEnvToProject(config.Org, *createdProject.Canonical, env)
	if err != nil {
		t.Fatalf("Failed to link env %q to project: %v", env, err)
	}

	newEnvName := "New"
	updateBody := &models.UpdateEnvironment{
		Name: ptr.Ptr(newEnvName),
		Type: ptr.Ptr("production"),
	}
	updatedEnv, _, err := m.UpdateOrgEnv(config.Org, *createdEnv.Canonical, updateBody)
	if err != nil {
		t.Fatalf("Failed to update env %q: %v", env, err)
	}

	assert.Equal(t, newEnvName, updatedEnv.Name)
}
