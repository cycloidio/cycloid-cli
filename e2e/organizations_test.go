package e2e

import (
	"fmt"
	"testing"

	"github.com/cycloidio/cycloid-cli/internal/testcfg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganizations(t *testing.T) {
	t.Run("SuccessOrganizationsGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"organization",
			"get",
		})

		require.NoError(t, cmdErr)
		assert.Contains(t, cmdOut, fmt.Sprintf("canonical\": \"%s", config.Org))
	})

	childOrg := testcfg.RandomCanonical(t.Name())
	t.Run("SuccessOrganizationsCreateChild", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", childOrg,
			"--parent-org", config.Org,
			"organization",
			"create-child",
		})

		require.NoError(t, cmdErr)
		assert.Contains(t, cmdOut, "[]")
	})

	t.Run("SuccessOrganizationsListChildrens", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"organization",
			"list-childrens",
		})

		require.NoError(t, cmdErr)
		assert.Contains(t, cmdOut, childOrg)
	})

	t.Run("SuccessOrganizationsListWorkers", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"organization",
			"list-workers",
		})

		require.NoError(t, cmdErr)
		assert.Contains(t, cmdOut, "[]")
	})
}
