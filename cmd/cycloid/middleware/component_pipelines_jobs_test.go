package middleware_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetJobs(t *testing.T) {
	m := config.Middleware

	pipelineName := fmt.Sprintf("%s-%s-%s", *config.Project.Canonical, *config.Environment.Canonical, *config.Component.Canonical)

	jobs, _, err := m.GetJobs(config.Org, *config.Project.Canonical, *config.Environment.Canonical, *config.Component.Canonical, pipelineName)
	require.NoError(t, err)
	assert.NotNil(t, jobs)
}
