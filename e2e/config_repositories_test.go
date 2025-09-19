package e2e_test

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestConfigRepositories(t *testing.T) {
	// Bug: https://linear.app/cycloid/issue/BE-981/catalog-repository-creation-fail-on-staging
	t.Skip()
	//

	var (
		CRUrl       = "git@github.com:cycloidio/cycloid-cli-test-catalog.git"
		CRBranch    = "config-e2e"
		CRCanonical = "config-repo-e2e"
		CRCred      = config.ConfigRepo.CredentialCanonical
	)

	defer t.Run("SuccessConfigRepositoriesDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"delete",
			"--canonical", CRCanonical,
		})

		assert.Nil(t, cmdErr)
		require.Empty(t, cmdOut)
	})

	t.Run("SuccessConfigRepositoriesCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"create",
			"--name", CRCanonical,
			"--branch", CRBranch,
			"--cred", CRCanonical,
			"--url", CRUrl,
		})

		assert.Nil(t, cmdErr)

		var result models.ConfigRepository
		err := json.Unmarshal([]byte(cmdOut), &result)
		assert.NoError(t, err)
		assert.Equal(t, CRUrl, *result.URL)
		assert.Equal(t, CRBranch, result.Branch)
		assert.Equal(t, CRCanonical, *result.Canonical)
		assert.Equal(t, CRCred, result.CredentialCanonical)
	})

	t.Run("SuccessConfigRepositoriesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"list",
		})

		assert.Nil(t, cmdErr)
		var list []models.ConfigRepository
		err := json.Unmarshal([]byte(cmdOut), &list)
		assert.NoError(t, err)

		if index := slices.IndexFunc(list, func(cr models.ConfigRepository) bool {
			return *cr.Canonical == CRCanonical
		}); index != -1 {
			result := list[index]
			assert.Equal(t, CRUrl, *result.URL)
			assert.Equal(t, CRBranch, result.Branch)
			assert.Equal(t, CRCanonical, *result.Canonical)
			assert.Equal(t, CRCred, result.CredentialCanonical)
		} else {
			t.Fatal("did not found our config repo in cmd output: " + cmdOut)
		}
	})

	t.Run("SuccessConfigRepositoriesGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"get",
			"--canonical", CRCanonical,
		})

		assert.Nil(t, cmdErr)

		var result models.ConfigRepository
		err := json.Unmarshal([]byte(cmdOut), &result)
		assert.NoError(t, err)
		assert.Equal(t, CRUrl, *result.URL)
		assert.Equal(t, CRBranch, result.Branch)
		assert.Equal(t, CRCanonical, *result.Canonical)
		assert.Equal(t, CRCred, result.CredentialCanonical)
	})
}
