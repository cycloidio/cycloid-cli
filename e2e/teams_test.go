package e2e_test

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTeams(t *testing.T) {
	teamCan := randomCanonical("team")
	t.Run("SuccessTeamCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"team", "create",
			"--name", teamCan,
			"--role", "organization-admin",
		})
		require.NoError(t, cmdErr, "cmd should succeed")

		defer t.Run("SuccessTeamDelete", func(t *testing.T) {
			_, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"team", "delete", teamCan,
			})
			require.NoError(t, cmdErr, "cmd delete should succeed")
		})

		var outTeam models.Team
		err := json.Unmarshal([]byte(cmdOut), &outTeam)
		assert.NoError(t, err, "unmarshal output should work")
		assert.NotNil(t, outTeam, "team should not be empty")
		assert.Equal(t, teamCan, ptr.Value(outTeam.Canonical))

		t.Run("SuccessTeamList", func(t *testing.T) {
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"team", "list",
			})
			require.NoError(t, cmdErr, "cmd should succeed")

			var outTeams []*models.Team
			err := json.Unmarshal([]byte(cmdOut), &outTeams)
			assert.NoError(t, err, "unmarshal output should work")
			index := slices.IndexFunc(outTeams, func(t *models.Team) bool {
				return ptr.Value(t.Canonical) == teamCan
			})
			assert.NotEqual(t, -1, index, "we should find our team index in the out array")
			assert.Equal(t, teamCan, ptr.Value(outTeams[index].Canonical))
		})

		t.Run("SuccessTeamUpdate", func(t *testing.T) {
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"team", "update",
				"--name", "newName",
				"--team", teamCan,
				"--role", "organization-admin",
			})
			require.NoError(t, cmdErr, "cmd should succeed")
			var outTeam models.Team
			err := json.Unmarshal([]byte(cmdOut), &outTeam)
			assert.NoError(t, err, "unmarshal output should work")
			assert.NotNil(t, outTeam, "team should not be empty")
			assert.Equal(t, teamCan, ptr.Value(outTeam.Canonical))
			assert.Equal(t, "newName", ptr.Value(outTeam.Name))
		})

		t.Run("SuccessTeamGet", func(t *testing.T) {
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"team", "get", teamCan,
			})
			require.NoError(t, cmdErr, "cmd should succeed")

			var outTeam models.Team
			err := json.Unmarshal([]byte(cmdOut), &outTeam)
			assert.NoError(t, err, "unmarshal output should work")
			assert.Equal(t, teamCan, ptr.Value(outTeam.Canonical))
		})

		t.Run("SuccessTeamMemberAssign", func(t *testing.T) {
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"team", "member", "assign", "admin@cycloid.io",
				"--team", teamCan,
			})
			require.NoError(t, cmdErr, "cmd should succeed")
			defer t.Run("SuccessTeamMemberUnAssign", func(t *testing.T) {
				_, cmdErr := executeCommand([]string{
					"--output", "json",
					"--org", config.Org,
					"team", "member", "unassign", "admin@cycloid.io",
					"--team", teamCan,
				})
				require.NoError(t, cmdErr, "cmd should succeed")
			})

			var outMember []models.MemberTeam
			litter.Dump(cmdOut)
			err := json.Unmarshal([]byte(cmdOut), &outMember)
			assert.NoError(t, err, "unmarshal output should work")
			assert.NotNil(t, outMember, "we should have a member")
			assert.Len(t, outMember, 1, "There should be one member in the team")
			assert.Equal(t, "admin@cycloid.io", ptr.Value(outMember[0].Email).String())
		})
	})
}
