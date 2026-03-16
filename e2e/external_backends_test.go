package e2e_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExternalBackends(t *testing.T) {
	// Prepare: create the AWS credential needed for the logs backend test
	t.Run("CleanupAndPrepare", func(t *testing.T) {
		// Clean external backends if exist
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"external-backends",
			"list",
		})

		require.Nil(t, cmdErr)
		cs, err := JSONListExtractFields(cmdOut, "id", "", "")
		require.Nil(t, err)

		for _, c := range cs {
			executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"eb",
				"delete",
				"--id", c,
			})
		}

		// Create AWS credential for the CloudWatch logs test
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"create",
			"aws",
			"--name", "eb-aws",
			"--access-key", "foo",
			"--secret-key", "bar",
		})
	})

	t.Run("SuccessExternalBackendsCreateAWSRemoteTFState", func(t *testing.T) {
		// TODO: Fix tests when components are implemented (BE-XXXX)
		t.Skip("blocked: requires component-level infraview support")

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"external-backends",
			"create",
			"infraview",
			"AWSRemoteTFState",
			"--bucket-name", "eb-ifraview-aws",
			"--bucket-path", "/foo",
			"--cred", "eb-aws",
			"--project", *config.Project.Canonical,
			"--env", *config.Environment.Canonical,
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "purpose\": \"remote_tfstate")
	})

	t.Run("SuccessExternalBackendsCreateAWSCloudWatchLogs", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"external-backends",
			"create",
			"logs",
			"AWSCloudWatchLogs",
			"--cred", "eb-aws",
			"--project", *config.Project.Canonical,
			"--region", "eu-west-1",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "purpose\": \"logs")
	})

	t.Run("SuccessExternalBackendsList", func(t *testing.T) {
		// TODO: Fix tests when components are implemented (BE-XXXX)
		t.Skip("blocked: list expects remote_tfstate entry from skipped infraview test")

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"external-backends",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "purpose\": \"remote_tfstate")
	})
}
