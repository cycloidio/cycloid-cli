package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBuilds(t *testing.T) {
	m := config.Middleware

	pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *config.Component.Canonical)

	jobs, _, err := m.GetJobs(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *config.Component.Canonical, pipelineName)
	require.NoError(t, err)
	require.NotEmpty(t, jobs, "expected at least one job in the default component pipeline")

	builds, _, err := m.GetBuilds(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *config.Component.Canonical, pipelineName, *jobs[0].Name)
	require.NoError(t, err)
	assert.NotNil(t, builds)
}
