package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListMembers(t *testing.T) {
	m := config.Middleware

	members, _, err := m.ListMembers(config.Org)
	require.NoError(t, err)
	assert.NotEmpty(t, members)
}

func TestListInvites(t *testing.T) {
	m := config.Middleware

	invites, _, err := m.ListInvites(config.Org)
	require.NoError(t, err)
	assert.NotNil(t, invites)
}

func TestGetMember(t *testing.T) {
	m := config.Middleware

	members, _, err := m.ListMembers(config.Org)
	require.NoError(t, err)
	require.NotEmpty(t, members)

	got, _,
		err := m.GetMember(config.Org, *members[0].ID)
	require.NoError(t, err)
	assert.Equal(t, *members[0].ID, *got.ID)
}
