package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEnvPipelines(t *testing.T) {
	m := config.Middleware

	pipelines, _, err := m.GetEnvPipelines(config.Org, *config.Project.Canonical, *config.Environment.Canonical)
	require.NoError(t, err)
	assert.NotNil(t, pipelines)
}
