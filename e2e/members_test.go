package e2e_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMembers(t *testing.T) {
	// Cleanup invites in case of a previous test
	t.Run("CleanupPreviousInvites", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"members",
			"list-invites",
		})

		require.Nil(t, cmdErr)
		ids, err := JSONListExtractFields(cmdOut, "id", "email", "^foo@bli.fr$")
		require.Nil(t, err)

		for _, id := range ids {
			executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"members",
				"delete-invite",
				"--invite", id,
			})
		}
	})

	t.Run("SuccessMembersList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"members",
			"list",
		})

		require.Nil(t, cmdErr)

		var memberList []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &memberList)
		require.Nil(t, err, "unmarshalling cli json output")

		ok := JSONListFindObjectValue(memberList, "email", "admin@cycloid.io")
		assert.True(t, ok, fmt.Sprint("member with admin@cycloid.io email address not found in json:\n", cmdOut))
	})

	t.Run("SuccessMembersGet", func(t *testing.T) {
		// Look up the admin member's numeric ID from the list
		listOut, listErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"members",
			"list",
		})
		require.Nil(t, listErr, "members list should not fail")

		ids, err := JSONListExtractFields(listOut, "id", "email", "^admin@cycloid.io$")
		require.Nil(t, err)
		require.NotEmpty(t, ids, "admin member should be in the list")

		adminID := ids[0]
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"members",
			"get",
			"--id", adminID,
		})

		require.Nil(t, cmdErr, "CLI should not error on members get")

		var member map[string]any
		err = json.Unmarshal([]byte(cmdOut), &member)
		require.Nil(t, err, "we should be able to serialize the CLI json output: ", cmdOut)

		assert.Equal(t, "admin@cycloid.io", member["email"],
			"member with admin@cycloid.io email address not found in json:\n", cmdOut)
	})

	t.Run("SuccessMembersUpdate", func(t *testing.T) {
		// Look up the admin member's numeric ID from the list
		listOut, listErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"members",
			"list",
		})
		require.Nil(t, listErr, "members list should not fail")

		ids, err := JSONListExtractFields(listOut, "id", "email", "^admin@cycloid.io$")
		require.Nil(t, err)
		require.NotEmpty(t, ids, "admin member should be in the list")

		adminID := ids[0]
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"members",
			"update",
			"--id", adminID,
			"--role", "organization-admin",
		})

		require.Nil(t, cmdErr, "members update should not fail")
		assert.NotEmpty(t, cmdOut, "members update should return the updated member")
	})

	t.Run("SuccessMembersInvite", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
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
			"--org", config.Org,
			"members",
			"list-invites",
		})

		require.Nil(t, cmdErr)

		var memberList []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &memberList)
		require.Nil(t, err, "unmarshalling cli json output")

		ok := JSONListFindObjectValue(memberList, "invitation_email", "foo@bli.fr")
		assert.True(t, ok, fmt.Sprint("member with foo@bli.fr email address not found in json:\n", cmdOut))
	})
}
