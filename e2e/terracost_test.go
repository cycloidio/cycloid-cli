package e2e_test

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestTerracost(t *testing.T) {
	t.Run("SuccessTerracostEstimate", func(t *testing.T) {
		tmpDir := t.TempDir()
		planPath := filepath.Join(tmpDir, "test-plan.json")
		WriteFile(planPath, TestTerraformPlanSample)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"terracost",
			"estimate",
			"--plan-path", planPath,
		})

		require.Nil(t, cmdErr, "terracost estimate should not fail, out: %s", cmdOut)

		// Output should deserialize to a cost estimation result
		var result models.CostEstimationResult
		err := json.Unmarshal([]byte(cmdOut), &result)
		assert.Nil(t, err, "output should deserialize to models.CostEstimationResult, out: %s", cmdOut)
	})
}
