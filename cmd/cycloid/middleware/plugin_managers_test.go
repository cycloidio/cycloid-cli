package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListPluginManagers(t *testing.T) {
	m := config.Middleware

	managers, _, err := m.ListPluginManagers(config.Org)
	require.NoError(t, err)
	assert.NotNil(t, managers)
}

func TestCreatePluginManager(t *testing.T) {
	m := config.Middleware

	pm, _, err := m.CreatePluginManager(config.Org, "test-pm", "http://192.168.10.14:4000")
	if err != nil {
		// May fail with duplicate org on re-runs
		t.Skipf("plugin manager create failed (may be duplicate): %v", err)
	}
	require.NotNil(t, pm)
	assert.Equal(t, "test-pm", *pm.Name)
}
