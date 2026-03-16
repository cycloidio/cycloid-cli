package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListExternalBackends(t *testing.T) {
	m := config.Middleware

	backends, _, err := m.ListExternalBackends(config.Org)
	require.NoError(t, err)
	assert.NotNil(t, backends)
}
