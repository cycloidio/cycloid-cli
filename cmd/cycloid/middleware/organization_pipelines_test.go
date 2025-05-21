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

	m := config.Middleware

	pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *config.Component.Canonical)
	got, err := m.GetOrgPipelines(config.Org, &pipelineName, config.Project.Canonical, config.Environment.Canonical, []string{})
	if err != nil {
		t.Errorf("middleware.GetOrgPipelines() error = %v", err)
		return
	}

	index := slices.IndexFunc(got, func(p *models.Pipeline) bool {
		return *p.Component.Canonical == *config.Component.Canonical
	})

	if index == -1 {
		t.Fatalf("failed to find created component in the list: %v\nexpected component: %v", litter.Sdump(got), litter.Sdump(config.Component))
	}
}
