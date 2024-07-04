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
			"--org", CY_TEST_ROOT_ORG,
			"--api-key", CY_TEST_ROOT_API_KEY,
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})

	t.Run("SuccessOrgLogin", func(t *testing.T) {
		err := os.Setenv("CY_API_KEY", CY_TEST_ROOT_API_KEY)
		require.Nil(t, err)

		cmdOut, cmdErr := executeCommand([]string{
			"login",
			"--org", CY_TEST_ROOT_ORG,
		})

		require.Nil(t, cmdErr)
		assert.Equal(t, "", cmdOut)
	})
}
