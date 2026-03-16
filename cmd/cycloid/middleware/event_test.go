package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSendEvent(t *testing.T) {
	m := config.Middleware

	_, err := m.SendEvent(config.Org, "Cycloid", "test event", "middleware test", "info", map[string]string{"env": "test"}, "blue")
	require.NoError(t, err)
}

func TestListEvents(t *testing.T) {
	m := config.Middleware

	events, _, err := m.ListEvents(config.Org, nil, nil, 0, 0)
	require.NoError(t, err)
	assert.NotNil(t, events)
}
