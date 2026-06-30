package e2e_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInventory exercises the `cy inventory` command tree end-to-end.
//
// Inventory outputs and resources originate from Terraform states and cannot be
// created directly via the API, so (like the middleware LHS filter discovery
// test) these checks validate that the commands wire up correctly and that the
// filter flags are accepted by the route without error — not specific data.
func TestInventory(t *testing.T) {
	t.Run("OutputsList", func(t *testing.T) {
		out, err := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"inventory", "outputs", "list",
		})
		require.NoError(t, err)

		var outputs []map[string]interface{}
		require.NoError(t, json.Unmarshal([]byte(out), &outputs))
	})

	t.Run("OutputsListWithFilters", func(t *testing.T) {
		_, err := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"inventory", "outputs", "list",
			"--type", "string",
			"--filter", "output_is_pinned[eq]=false",
		})
		require.NoError(t, err)
	})

	t.Run("OutputsGetNotFound", func(t *testing.T) {
		_, err := executeCommand([]string{
			"--org", config.Org,
			"inventory", "outputs", "get", "definitely-nonexistent-output-key",
		})
		assert.Error(t, err, "getting a nonexistent output key must error")
	})

	t.Run("ResourcesList", func(t *testing.T) {
		out, err := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"inventory", "resources", "list",
		})
		require.NoError(t, err)

		var resources []map[string]interface{}
		require.NoError(t, json.Unmarshal([]byte(out), &resources))
	})

	t.Run("ResourcesListWithFilters", func(t *testing.T) {
		_, err := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"inventory", "resources", "list",
			"--provider", "aws",
			"--filter", "resources_mode[eq]=managed",
		})
		require.NoError(t, err)
	})
}
