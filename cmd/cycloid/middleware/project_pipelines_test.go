package middleware_test

import (
	"slices"
	"testing"

	"github.com/sanity-io/litter"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestGetProjectPipelines(t *testing.T) {
	m := config.Middleware

	got, err := m.GetProjectPipelines(config.Org, *config.Project.Canonical)
	if err != nil {
		t.Errorf("middleware.GetProjectPipelines() error = %v", err)
		return
	}

	index := slices.IndexFunc(got, func(p *models.Pipeline) bool {
		return *p.Component.Canonical == *config.Component.Canonical
	})

	if index == -1 {
		t.Fatalf("failed to find created component in the list: %v\nexpected component: %v", litter.Sdump(got), litter.Sdump(config.Component))
	}
}
