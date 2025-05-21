package middleware_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/sanity-io/litter"
)

var (
	pipelineTestStackRef     = "cycloid:stack-test-pipeline"
	pipelineTestStackUseCase = "default"
	pipelineTestDefaultVars  = models.FormVariables{"section": {"group": {"var": "hello"}}}
)

func TestGetOrgPipelines(t *testing.T) {
	t.Parallel()
	testName := "getOrgPipelines"
	config, err := GetTestConfig()
	if err != nil {
		t.Errorf("failed to get test config: %s", err)
		return
	}

	project, projectDelete, err := config.SetupTestProject(testName)
	defer projectDelete()
	if err != nil {
		t.Errorf("failed to create project for test '%s': %s", testName, err)
	}

	env, envDelete, err := config.SetupTestEnv(testName, *project.Canonical)
	defer envDelete()
	if err != nil {
		t.Errorf("failed to setup env for test '%s': %s", testName, err)
		return
	}

	component, componentDelete, err := config.SetupTestComponent(
		*project.Canonical, *env.Canonical, testName, pipelineTestStackRef, pipelineTestStackUseCase, &pipelineTestDefaultVars,
	)
	defer componentDelete()
	if err != nil {
		t.Errorf("failed to setup base component for test '%s': %s", testName, err)
	}

	pipelineName := fmt.Sprintf("%s-%s-%s", *project.Canonical, *env.Canonical, *component.Canonical)
	got, err := m.GetOrgPipelines(config.Org, &pipelineName, project.Canonical, env.Canonical, []string{})
	if err != nil {
		t.Errorf("middleware.GetOrgPipelines() error = %v", err)
		return
	}

	index := slices.IndexFunc(got, func(p *models.Pipeline) bool {
		return *p.Component.Canonical == *component.Canonical
	})

	if index == -1 {
		t.Fatalf("failed to find created component in the list: %v\nexpected component: %v", litter.Sdump(got), litter.Sdump(component))
	}
}
