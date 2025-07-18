//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigRepositories(t *testing.T) {
	t.Skip()
	LoginToRootOrg()

	t.Run("SuccessConfigRepositoriesCreate", func(t *testing.T) {
		// Cleanup just in case
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"delete",
			"--canonical", "default-config",
		})

		// Create required ssh cred
		WriteFile("/tmp/test_cli-ssh", TestGitSshKey)
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"create",
			"ssh",
			"--name", "git-creds",
			"--ssh-key", "/tmp/test_cli-ssh",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"create",
			"--name", "default-config",
			"--branch", CY_TEST_GIT_CR_BRANCH,
			"--cred", "git-creds",
			"--url", CY_TEST_GIT_CR_URL,
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"default-config")
	})

	t.Run("SuccessConfigRepositoriesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"default-config")
	})

	t.Run("SuccessConfigRepositoriesGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"get",
			"--canonical", "default-config",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"default-config")
	})

	t.Run("SuccessConfigRepositoriesDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"delete",
			"--canonical", "default-config",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", string(cmdOut))
	})
}
