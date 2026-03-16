package e2e_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfraPolicies(t *testing.T) {
	t.Skip()
	ipCanonical := randomCanonical("test-ip")
	tmpDir := t.TempDir()
	policyPath := filepath.Join(tmpDir, "test-cli-ip.rego")
	WriteFile(policyPath, TestInfraPolicySample)

	// Checks the successful creation of a new infrapolicy
	t.Run("SuccessInfraPolicyCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"ip",
			"create",
			"--policy-path", policyPath,
			"--canonical", ipCanonical,
			"--name", ipCanonical,
			"--description", "test infrapolicy",
			"--owner", "cycloidio",
			"--severity", "advisory",
			"--enabled=true",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, ipCanonical)
	})

	// Checks the successful get of a new infrapolicy
	t.Run("SuccessInfraPolicyGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"ip",
			"get",
			"--canonical", ipCanonical,
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "description\": \"test infrapolicy")
	})

	// Checks the successful list of infrapolicies in org
	t.Run("SuccessInfraPoliciesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"ip",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "description\": \"test infrapolicy")
	})

	// Checks the successful update of a infrapolicy
	t.Run("SuccessInfraPolicyUpdate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"ip",
			"update",
			"--canonical", ipCanonical,
			"--policy-path", policyPath,
			"--name", ipCanonical,
			"--description", "changed description",
			"--owner", "cycloidio",
			"--severity", "advisory",
			"--enabled=false",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "description\": \"changed description")
	})

	// Checks that validate runs against a terraform plan
	t.Run("SuccessInfraPoliciesValidate", func(t *testing.T) {
		planPath := filepath.Join(tmpDir, "test-plan.json")
		WriteFile(planPath, TestTerraformPlanSample)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"ip",
			"validate",
			"--plan-path", planPath,
			"--project", *config.Project.Canonical,
			"--env", *config.Environment.Canonical,
		})

		assert.Nil(t, cmdErr, "validate should not fail, out: %s", cmdOut)
	})

	// Checks the successful deletion of a infrapolicy
	t.Run("SuccessInfraPolicyDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"ip",
			"delete",
			"--canonical", ipCanonical,
		})

		assert.Nil(t, cmdErr)
		require.Empty(t, cmdOut)
	})
}
