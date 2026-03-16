package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

func TestTeamMembersCRUD(t *testing.T) {
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-team-members")

	_, _, err := m.CreateTeam(config.Org, ptr.Ptr("Test Team Members"), ptr.Ptr(canonical), nil, []string{})
	require.NoError(t, err, "CreateTeam should succeed")

	defer func() {
		_, err := m.DeleteTeam(config.Org, canonical)
		require.NoError(t, err, "DeleteTeam cleanup should succeed")
	}()

	members, _, err := m.ListTeamMembers(config.Org, canonical)
	require.NoError(t, err, "ListTeamMembers should succeed on empty team")
	assert.NotNil(t, members)

	assigned, _, err := m.AssignMemberToTeam(config.Org, canonical, ptr.Ptr("administrator"), nil)
	require.NoError(t, err, "AssignMemberToTeam should succeed")
	require.NotNil(t, assigned)

	got, _, err := m.GetTeamMember(config.Org, canonical, *assigned.ID)
	require.NoError(t, err, "GetTeamMember should succeed")
	assert.Equal(t, *assigned.ID, *got.ID)

	_, err = m.UnAssignMemberFromTeam(config.Org, canonical, *assigned.ID)
	require.NoError(t, err, "UnAssignMemberFromTeam should succeed")
}
