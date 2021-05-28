//+build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreds(t *testing.T) {
	LoginToRootOrg()

	t.Run("SuccessCredsList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"vault")
	})

	t.Run("SuccessCredsGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"get",
			"--canonical", "vault",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"vault")
	})

	t.Run("SuccessCredsCreateCustom", func(t *testing.T) {
		// Cleanup just in case
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"delete",
			"--canonical", "cli-custom",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"create",
			"custom",
			"--name", "cli-custom",
			"--field", "foo=bar",
			"--field", "int=1",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessCredsCreateSSH", func(t *testing.T) {
		// Cleanup just in case
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"delete",
			"--canonical", "cli-ssh",
		})

		WriteFile("/tmp/test_cli-ssh", TestGitSshKey)
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"create",
			"ssh",
			"--name", "cli-ssh",
			"--ssh-key", "/tmp/test_cli-ssh",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})
}
