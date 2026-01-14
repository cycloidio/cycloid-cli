package e2e_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExternalBackends(t *testing.T) {
	t.Skip()

	// Prepare a running project
	t.Run("CleanupAndPrepare", func(t *testing.T) {

		// // Clean external backends if exist
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

		// Create ssh cred
		WriteFile("/tmp/test_cli-ssh", TestGitSSHKey)
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"creds",
			"create",
			"ssh",
			"--name", "git-project-creds",
			"--ssh-key", "/tmp/test_cli-ssh",
		})
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

		// Create config repo
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"config-repo",
			"create",
			"--name", "project-config",
			"--branch", config.ConfigRepo.Branch,
			"--cred", "git-project-creds",
			"--url", *config.ConfigRepo.URL,
		})

		// Provide service catalog public
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"catalog-repository",
			"create",
			"--branch", "master",
			"--url", "https://github.com/cycloid-community-catalog/stack-dummy.git",
			"--name", "dummy",
		})

		// Create project
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)
		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"project",
			"create",
			"--name", "eb-test",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", config.Org),
			"--config-repo", "project-config",
		})

		executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"project",
			"create-env",
			"--project", "eb-test",
			"--env", "test",
			"--use-case", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		// Ensure the catalog is present
		cmdOut, cmdErr = executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"project",
			"get",
			"--project", "eb-test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"eb-test")
	})

	t.Run("SuccessExternalBackendsCreateAWSRemoteTFState", func(t *testing.T) {
		// TODO: Fix tests when components are implemented
		t.Skip()

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
			"--project", "eb-test",
			"--env", "test",
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
			"--project", "eb-test",
			"--region", "eu-west-1",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "purpose\": \"logs")
	})

	t.Run("SuccessExternalBackendsList", func(t *testing.T) {
		// TODO: Fix tests when components are implemented
		t.Skip()

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
