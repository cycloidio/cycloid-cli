package e2e

import (
	"testing"
)

func TestPipelines(t *testing.T) {
	t.Parallel()

	// Preparation

	t.Run("ListOrgPipelines", func(t *testing.T) {

	})

	// // Prepare a running project
	// t.Run("CleanupAndPrepare", func(t *testing.T) {
	// 	// Create ssh cred
	// 	WriteFile("/tmp/test_cli-ssh", TestGitSshKey)
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"creds",
	// 		"create",
	// 		"ssh",
	// 		"--name", "git-project-creds",
	// 		"--ssh-key", "/tmp/test_cli-ssh",
	// 	})
	//
	// 	// Create config repo
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"config-repo",
	// 		"create",
	// 		"--name", "project-config",
	// 		"--branch", CyTestCatalogRepoBranch,
	// 		"--cred", "git-project-creds",
	// 		"--url", CyTestCatalogRepoURL,
	// 	})
	//
	// 	// Provide service catalog public
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"catalog-repository",
	// 		"create",
	// 		"--branch", "master",
	// 		"--url", "https://github.com/cycloid-community-catalog/stack-dummy.git",
	// 		"--name", "dummy",
	// 	})
	//
	// 	// Create project
	// 	WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
	// 	WriteFile("/tmp/test_cli-pp", TestPipelineSample)
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"project",
	// 		"create",
	// 		"--name", "pipeline-test",
	// 		"--description", "this is a test project",
	// 		"--stack-ref", fmt.Sprintf("%s:stack-dummy", TestRootOrg),
	// 		"--config-repo", "project-config",
	// 	})
	//
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"project",
	// 		"create-env",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--use-case", "default",
	// 		"--vars", "/tmp/test_cli-pp-vars",
	// 		"--pipeline", "/tmp/test_cli-pp",
	// 		"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
	// 	})
	//
	// 	// Ensure the catalog is present
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"project",
	// 		"get",
	// 		"--project", "pipeline-test",
	// 	})
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "canonical\": \"pipeline-test")
	// })
	//
	// t.Run("SuccessPipelinesUpdate", func(t *testing.T) {
	// 	WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
	// 	WriteFile("/tmp/test_cli-pp", TestPipelineSample)
	//
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"update",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--vars", "/tmp/test_cli-pp-vars",
	// 		"--pipeline", "/tmp/test_cli-pp",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "canonical\": \"pipeline-test")
	// })
	//
	// t.Run("SuccessPipelinesPause", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"pause",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Equal(t, "", cmdOut)
	// })
	//
	// t.Run("SuccessPipelinesUnPause", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"unpause",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Equal(t, "", cmdOut)
	// })
	//
	// t.Run("SuccessPipelinesList", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"list",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "name\": \"pipeline-test-test")
	// })
	//
	// t.Run("SuccessPipelinesGet", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"get",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "name\": \"pipeline-test-test")
	// })
	//
	// t.Run("SuccessPipelinesListJobs", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"list-jobs",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "name\": \"job-hello-world")
	// })
	//
	// t.Run("SuccessPipelinesGetJob", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"get-job",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--job", "job-hello-world",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "name\": \"job-hello-world")
	// })
	//
	// t.Run("SuccessPipelinesPauseJob", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"pause-job",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--job", "job-hello-world",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Equal(t, "", cmdOut)
	// })
	//
	// t.Run("SuccessPipelinesUnpauseJob", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"unpause-job",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--job", "job-hello-world",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Equal(t, "", cmdOut)
	// })
	// t.Run("SuccessPipelinesTriggerBuild", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"trigger-build",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--job", "job-hello-world",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Equal(t, "", cmdOut)
	// })
	//
	// t.Run("SuccessPipelinesListBuilds", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"list-builds",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--job", "job-hello-world",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Contains(t, cmdOut, "job_name\": \"job-hello-world")
	// })
	//
	// t.Run("SuccessPipelinesClearTaskCache", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"clear-task-cache",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 		"--job", "job-hello-world",
	// 		"--task", "hello-world",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	require.Equal(t, "", cmdOut)
	// })
	//
	// t.Run("SuccessPipelinesSynced", func(t *testing.T) {
	// 	cmdOut, cmdErr := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", TestRootOrg,
	// 		"pipeline",
	// 		"synced",
	// 		"--project", "pipeline-test",
	// 		"--env", "test",
	// 	})
	//
	// 	// TODO: Fix tests when components are implemented
	// 	t.Skip()
	//
	// 	assert.Nil(t, cmdErr)
	// 	// Note: we expect no diff because the pipeline from helpers.go is the same as the dummy-stack.
	// 	// This mean if someone change the code from the dummy stack, this test could fail because the helper
	// 	// pipeline will differ from the one in the dummy stack
	// 	require.Contains(t, cmdOut, "jobs\": null")
	// })
}
