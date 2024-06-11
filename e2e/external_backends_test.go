//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExternalBackends(t *testing.T) {
	LoginToRootOrg()

	// Prepare a running project
	t.Run("CleanupAndPrepare", func(t *testing.T) {

		// // Clean external backends if exist
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"external-backends",
			"list",
		})

		require.Nil(t, cmdErr)
		cs, err := JsonListExtractFields(cmdOut, "id", "", "")
		require.Nil(t, err)

		for _, c := range cs {
			executeCommand([]string{
				"--output", "json",
				"--org", CY_TEST_ROOT_ORG,
				"eb",
				"delete",
				"--id", c,
			})
		}

		// Create ssh cred
		WriteFile("/tmp/test_cli-ssh", TestGitSshKey)
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"creds",
			"create",
			"ssh",
			"--name", "git-project-creds",
			"--ssh-key", "/tmp/test_cli-ssh",
		})
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
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
			"--org", CY_TEST_ROOT_ORG,
			"config-repo",
			"create",
			"--name", "project-config",
			"--branch", CY_TEST_GIT_CR_BRANCH,
			"--cred", "git-project-creds",
			"--url", CY_TEST_GIT_CR_URL,
		})

		// Provide service catalog public
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
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
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create",
			"--name", "eb-test",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", CY_TEST_ROOT_ORG),
			"--config-repo", "project-config",
			"--env", "test",
			"--usecase", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		// Ensure the catalog is present
		cmdOut, cmdErr = executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"get",
			"--project", "eb-test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"eb-test")
	})

	t.Run("SuccessExternalBackendsCreateAWSRemoteTFState", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
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
			"--org", CY_TEST_ROOT_ORG,
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
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"external-backends",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "purpose\": \"remote_tfstate")
	})
}
