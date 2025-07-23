package e2e_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIKeyEnvVar(t *testing.T) {
	for _, envVar := range []string{"CY_API_KEY", "CY_API_TOKEN", "TOKEN"} {
		// We do a project list to check if we are authenticater
		t.Run("SuccessProjectListWithEnvVarAuth", func(t *testing.T) {
			os.Setenv(envVar, config.APIKey)
			_, err := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"project",
				"list",
			})

			assert.Nil(t, err, "Command should not fail using an env var for authentication")
			os.Unsetenv(envVar)
		})
	}
}
