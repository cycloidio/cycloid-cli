package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

func TestTeamsCRUD(t *testing.T) {
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-team")

	created, _, err := m.CreateTeam(config.Org, ptr.Ptr("Test Team"), ptr.Ptr(canonical), nil, []string{})
	require.NoError(t, err, "CreateTeam should succeed")
	require.NotNil(t, created)

	defer func() {
		_, err := m.DeleteTeam(config.Org, canonical)
		require.NoError(t, err, "DeleteTeam should succeed")
	}()

	got, _, err := m.GetTeam(config.Org, canonical)
	require.NoError(t, err, "GetTeam should succeed")
	assert.Equal(t, canonical, *got.Canonical)

	list, _, err := m.ListTeams(config.Org, nil, nil, nil, nil)
	require.NoError(t, err, "ListTeams should succeed")
	assert.NotEmpty(t, list)

	newName := ptr.Ptr("Updated Team Name")
	updated, _, err := m.UpdateTeam(config.Org, newName, ptr.Ptr(canonical), nil, []string{})
	require.NoError(t, err, "UpdateTeam should succeed")
	assert.Equal(t, *newName, *updated.Name)

}
