package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoles(t *testing.T) {
	t.Run("SuccessRolesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"roles",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"organization-member")
	})

	t.Run("SuccessRolesCreate", func(t *testing.T) {
		ruleFile, err := WriteTempFile(`[
		{"action": "organization:update", "effect": "allow", "resources": []},
		{"action": "organization:list", "effect": "allow", "resources": []}
		]`)
		assert.NoError(t, err, "test setup failed to write temp file", ruleFile)

		roleCan := randomCanonical("role")
		cmd := []string{
			"--output", "json",
			"--org", config.Org,
			"roles",
			"create",
			"--name", "Team leader",
			"--role", roleCan,
			"--description", "big boss",
			"--rule-json", `{"action": "organization:delete", "effect": "allow", "resources": []}`,
			"--rule-json", `{"action": "organization:create", "effect": "allow", "resources": []}`,
			"--rule-file", ruleFile,
		}
		cmdOut, cmdErr := executeCommand(cmd)

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, roleCan)

		defer t.Run("SuccessRoleDelete", func(t *testing.T) {
			_, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"roles",
				"delete",
				roleCan,
			})
			require.Nil(t, cmdErr)
		})
	})

	t.Run("SuccessRolesGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"roles",
			"get",
			"organization-member",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"organization-member")
	})
}
