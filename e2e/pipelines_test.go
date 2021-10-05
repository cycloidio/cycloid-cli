//+build e2e

package e2e

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPipelines(t *testing.T) {
	LoginToRootOrg()

	// Prepare a running project
	t.Run("CleanupAndPrepare", func(t *testing.T) {
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
			"--name", "pipeline-test",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy",  CY_TEST_ROOT_ORG),
			"--config-repo", "project-config",
			"--env", "test",
			"--usecase", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		// Ensure the catalog is present
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"get",
			"--project", "pipeline-test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"pipeline-test")
	})

	t.Run("SuccessPipelinesUpdate", func(t *testing.T) {
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"update",
			"--project", "pipeline-test",
			"--env", "test",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\":\"pipeline-test")
	})

	t.Run("SuccessPipelinesPause", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"pause",
			"--project", "pipeline-test",
			"--env", "test",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessPipelinesUnPause", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"unpause",
			"--project", "pipeline-test",
			"--env", "test",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessPipelinesList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "name\":\"pipeline-test-test")
	})

	t.Run("SuccessPipelinesListJobs", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"list-jobs",
			"--project", "pipeline-test",
			"--env", "test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "name\":\"job-hello-world")
	})

	t.Run("SuccessPipelinesGetJob", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"get-job",
			"--project", "pipeline-test",
			"--env", "test",
			"--job", "job-hello-world",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "name\":\"job-hello-world")
	})

	t.Run("SuccessPipelinesPauseJob", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"pause-job",
			"--project", "pipeline-test",
			"--env", "test",
			"--job", "job-hello-world",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessPipelinesUnpauseJob", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"unpause-job",
			"--project", "pipeline-test",
			"--env", "test",
			"--job", "job-hello-world",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})
	t.Run("SuccessPipelinesTriggerBuild", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"trigger-build",
			"--project", "pipeline-test",
			"--env", "test",
			"--job", "job-hello-world",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessPipelinesListBuilds", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"list-builds",
			"--project", "pipeline-test",
			"--env", "test",
			"--job", "job-hello-world",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "job_name\":\"job-hello-world")
	})
	t.Run("SuccessPipelinesClearTaskCache", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"pipeline",
			"clear-task-cache",
			"--project", "pipeline-test",
			"--env", "test",
			"--job", "job-hello-world",
			"--task", "hello-world",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})
}
