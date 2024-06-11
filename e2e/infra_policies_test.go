//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfraPolicies(t *testing.T) {
	LoginToRootOrg()

	// Checks the succesfull creation of a new infrapolicy
	// The test validates that the reply of the create method
	// contains the cannonical of the created infrapolicy
	t.Run("SuccessInfraPolicyCreate", func(t *testing.T) {
		WriteFile("/tmp/test-cli-ip.rego", TestInfraPolicySample)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"ip",
			"create",
			"--policy-path", "/tmp/test-cli-ip.rego",
			"--name", "test",
			"--description", "test infrapolicy",
			"--owner", "cycloidio",
			"--severity", "advisory",
			"--enabled=1",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"test")
	})

	// Checks the succesfull get of a new infrapolicy
	// The test validates that the reply of the get method
	// contains the description of the infrapolicy
	t.Run("SuccessInfraPolicyGet", func(t *testing.T) {

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"ip",
			"get",
			"--cannonical", "test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "description\": \"test infrapolicy")
	})

	// Checks the succesfull list of infrapolicies in org
	// The test validates that the reply of the list method
	// contains the infrapolicy (previously created)
	t.Run("SuccessInfraPoliciesList", func(t *testing.T) {

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"ip",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "description\": \"test infrapolicy")
	})

	// Checks the succesfull update of a infrapolicy
	// The test validates that the reply of the update method
	// contains the changed description of the infrapolicy
	t.Run("SuccessInfraPolicyUpdate", func(t *testing.T) {
		//WriteFile("/tmp/update-test-cli-ip.rego", TestInfraPolicySample)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"ip",
			"update",
			"--cannonical", "test",
			"--policy-path", "/tmp/test-cli-ip.rego",
			"--name", "test",
			"--description", "changed description",
			"--owner", "cycloidio",
			"--severity", "advisory",
			"--enabled=false",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "description\": \"changed description")
	})

	// Checks the succesfull deletion of a infrapolicy
	// The test validates that the reply of the delete method
	// has no error
	t.Run("SuccessInfraPolicyDelete", func(t *testing.T) {

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"ip",
			"delete",
			"--cannonical", "test",
		})

		assert.Nil(t, cmdErr)
		require.Empty(t, cmdOut)
	})
}

//TODO! validate test
