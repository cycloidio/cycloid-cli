//+build e2e

package e2e

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganizations(t *testing.T) {
	LoginToRootOrg()

	t.Run("SuccessOrganizationsGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"organization",
			"get",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, fmt.Sprintf("canonical\":\"%s",  CY_TEST_ROOT_ORG))
	})

	t.Run("SuccessOrganizationsListChildrens", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"organization",
			"list-childrens",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "[]")
	})

	t.Run("SuccessOrganizationsListWorkers", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"organization",
			"list-workers",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "[]")
	})
}
