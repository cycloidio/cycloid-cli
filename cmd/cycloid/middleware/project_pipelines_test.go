package middleware_test

import (
	"slices"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/sanity-io/litter"
)

func TestGetProjectPipelines(t *testing.T) {
	t.Parallel()
	config, err := getTestConfig()
	if err != nil {
		t.Errorf("failed to get test config: %s", err)
		return
	}

	project, projectDelete, err := config.setupTestProject(t.Name())
	defer projectDelete()
	if err != nil {
		t.Errorf("failed to create project for test '%s': %s", t.Name(), err)
	}

	env, envDelete, err := config.setupTestEnv(t.Name(), *project.Canonical)
	defer envDelete()
	if err != nil {
		t.Errorf("failed to setup env for test '%s': %s", t.Name(), err)
		return
	}

	component, componentDelete, err := config.setupTestComponent(
		*project.Canonical, *env.Canonical, t.Name(), pipelineTestStackRef, pipelineTestStackUseCase, &pipelineTestDefaultVars,
	)
	defer componentDelete()
	if err != nil {
		t.Errorf("failed to setup base component for test '%s': %s", t.Name(), err)
	}

	got, err := m.GetProjectPipelines(config.Org, *project.Canonical)
	if err != nil {
		t.Errorf("middleware.GetProjectPipelines() error = %v", err)
		return
	}

	index := slices.IndexFunc(got, func(p *models.Pipeline) bool {
		return *p.Component.Canonical == *component.Canonical
	})

	if index == -1 {
		t.Fatalf("failed to find created component in the list: %v\nexpected component: %v", litter.Sdump(got), litter.Sdump(component))
	}
}
