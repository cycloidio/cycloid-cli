package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCatalogRepositories(t *testing.T) {
	defer t.Run("SuccessCatalogRepositoriesDelete", func(t *testing.T) {
		// Cleanup just in case
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"delete",
			"--canonical", "step-by-step",
		})

		assert.Nil(t, cmdErr, "failed to delete test catalog repo:", cmdOut)
	})

	t.Run("SuccessCatalogRepositoriesCreate", func(t *testing.T) {
		// Cleanup just in case
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"delete",
			"--canonical", "step-by-step",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"create",
			"--branch", "stack-aws",
			"--url", "https://github.com/cycloid-community-catalog/docs-step-by-step-stack.git",
			"--name", "step-by-step",
		})

		require.Nil(t, cmdErr)

		assert.Contains(t, cmdOut, "canonical\": \"stack-aws-sample")
	})

	t.Run("SuccessCatalogRepositoriesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"step-by-step")
	})

	t.Run("SuccessCatalogRepositoriesGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"get",
			"--canonical", "step-by-step",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"stack-aws-sample")
	})

	t.Run("SuccessCatalogRepositoriesRefresh", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"refresh",
			"--canonical", "step-by-step",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "updated")
	})
}
