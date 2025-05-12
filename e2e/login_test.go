package e2e

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Run("SuccessOrgLoginLegacy", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"login",
			"--org", TestRootOrg,
			"--api-key", TestAPIKey,
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})

	t.Run("SuccessOrgLogin", func(t *testing.T) {
		err := os.Setenv("CY_API_KEY", TestAPIKey)
		require.Nil(t, err)

		cmdOut, cmdErr := executeCommand([]string{
			"login",
			"--org", TestRootOrg,
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})
}
