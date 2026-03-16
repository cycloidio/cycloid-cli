package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRefreshToken(t *testing.T) {
	m := config.Middleware

	session, _, err := m.RefreshToken(&config.Org, nil, config.APIKey)
	require.NoError(t, err)
	assert.NotNil(t, session)
	assert.NotEmpty(t, session.Token)
}
