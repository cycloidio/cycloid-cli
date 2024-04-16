//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Run("SuccessOrgLogin", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"login",
			"--org", CY_TEST_ROOT_ORG,
			"--api-key", CY_TEST_ROOT_API_KEY,
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})

	t.Run("ErrorMissingRequiredFlag", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"login",
			"--org", CY_TEST_ROOT_ORG,
		})

		require.NotNil(t, cmdErr)
		assert.Contains(t, string(cmdErr.Error()), "required flag(s) \"api-key\" not set")
	})
}
