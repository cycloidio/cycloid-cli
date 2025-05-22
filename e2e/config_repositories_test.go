package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigRepositories(t *testing.T) {
	defer t.Run("SuccessConfigRepositoriesDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"delete",
			"--canonical", "test-config",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", string(cmdOut))
	})
	t.Run("SuccessConfigRepositoriesCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"create",
			"--name", "test-config",
			"--branch", "",
			"--cred", config.ConfigRepo.CredentialCanonical,
			"--url", *config.ConfigRepo.Canonical,
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"test-config")
	})

	t.Run("SuccessConfigRepositoriesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"test-config")
	})

	t.Run("SuccessConfigRepositoriesGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"get",
			"--canonical", "test-config",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"test-config")
	})

}
