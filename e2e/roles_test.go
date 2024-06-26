//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoles(t *testing.T) {
	LoginToRootOrg()

	t.Run("SuccessRolesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"roles",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"organization-member")
	})

	t.Run("SuccessRolesGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"roles",
			"get",
			"--canonical", "organization-member",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"organization-member")
	})
}
