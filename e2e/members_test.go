package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMembers(t *testing.T) {
	t.Skip()
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

		var memberList []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &memberList)
		require.Nil(t, err, "unmarshalling cli json output")

		ok := JsonListFindObjectValue(memberList, "email", "cycloidio@cycloid.io")
		assert.True(t, ok, fmt.Sprint("member with cycloidio@cycloid.io email address not found in json:\n", cmdOut))
	})

	t.Run("SuccessMembersGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"members",
			"get",
			"--id", "1",
		})

		require.Nil(t, cmdErr, "CLI should not error on this action.")

		var member map[string]any
		err := json.Unmarshal([]byte(cmdOut), &member)
		require.Nil(t, err, "we should be able to serialized the CLI json output: ", cmdOut)

		assert.Equal(t, member["email"], "cycloidio@cycloid.io",
			"member with cycloidio@cycloid.io email address not found in json:\n", cmdOut)
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

		var memberList []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &memberList)
		require.Nil(t, err, "unmarshalling cli json output")

		ok := JsonListFindObjectValue(memberList, "invitation_email", "foo@bli.fr")
		assert.True(t, ok, fmt.Sprint("member with foo@bli.fr email address not found in json:\n", cmdOut))
	})
}
