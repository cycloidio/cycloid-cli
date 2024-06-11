//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMembers(t *testing.T) {
	LoginToRootOrg()

	// Cleanup invites in case of a previous test
	t.Run("CleanupPreviousInvites", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"members",
			"list-invites",
		})

		require.Nil(t, cmdErr)
		ids, err := JsonListExtractFields(cmdOut, "id", "email", "^foo@bli.fr$")
		require.Nil(t, err)

		for _, id := range ids {
			executeCommand([]string{
				"--output", "json",
				"--org", CY_TEST_ROOT_ORG,
				"members",
				"delete-invite",
				"--invite", id,
			})
		}
	})

	t.Run("SuccessMembersList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"members",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "email\": \"cycloidio@cycloid.io")
	})

	t.Run("SuccessMembersGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"members",
			"get",
			"--name", "cycloidio",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "email\": \"cycloidio@cycloid.io")
	})

	t.Run("SuccessMembersInvite", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"members",
			"invite",
			"--email", "foo@bli.fr",
			"--role", "organization-admin",
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})

	t.Run("SuccessMembersListInvite", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"members",
			"list-invites",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "email\": \"foo@bli.fr")
	})
}
