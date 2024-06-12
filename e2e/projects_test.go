// //go:build e2e
// // +build e2e
package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjects(t *testing.T) {
	LoginToRootOrg()

	// Cleanup previous project if exist and prepare required catalog repository, ...
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

		// Here is an example if you want to add a specific catalog.
		// Since the latest update we have by default all the public stacks
		// // Provide service catalog public
		// executeCommand([]string{
		// 	"--output", "json",
		// 	"--org", CY_TEST_ROOT_ORG,
		// 	"catalog-repository",
		// 	"create",
		// 	"--branch", "master",
		// 	"--url", "https://github.com/cycloid-community-catalog/stack-dummy.git",
		// 	"--name", "dummy",
		// })
		//
		// // Ensure the catalog is present
		// cmdOut, cmdErr := executeCommand([]string{
		// 	"--output", "json",
		// 	"--org", CY_TEST_ROOT_ORG,
		// 	"catalog-repository",
		// 	"get",
		// 	"--canonical", "dummy",
		// })
		//
		// assert.Nil(t, cmdErr)
		// require.Contains(t, cmdOut, "canonical\": \"dummy")
	})

	t.Run("SuccessLegacyProjectsCreate", func(t *testing.T) {
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)

		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"delete",
			"--project", "snowy",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create",
			"--name", "snowy",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", CY_TEST_ROOT_ORG),
			"--config-repo", "project-config",
			"--env", "test",
			"--usecase", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
	})

	t.Run("SuccessLegacyProjectsCreateEnv", func(t *testing.T) {
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-env",
			"--project", "snowy",
			"--env", "test2",
			"--usecase", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
	})

	t.Run("SuccessProjectsList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"list",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
	})

	// Vars
	t.Run("SuccessProjectsCreateVars", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-stackforms-env",
			"--project", "snowy",
			"--env", "sf-vars",
			"--use-case", "default",
			"-V", `{"pipeline": {"config": {"message": "filledFromVars"}}}`,
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
		var data models.Project
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.Nil(t, err)
	})

	t.Run("SuccessProjectGetStackformConfigVars", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project", "get-env-config",
			"-p", "snowy", "-e", "sf-vars",
		})

		assert.Nil(t, cmdErr)

		// Output should be in json by default
		var data = make(map[string]map[string]map[string]any)
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.NoError(t, err)

		message, ok := data["pipeline"]["config"]["message"]
		assert.True(t, ok)
		assert.Equal(t, "filledFromVars", message)
	})

	// Extra vars
	t.Run("SuccessProjectsCreateExtraVars", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-stackforms-env",
			"--project", "snowy",
			"--env", "sf-extra-vars",
			"--use-case", "default",
			"-x", `pipeline.config.message=filledFromExtraVars`,
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
		var data models.Project
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.Nil(t, err)
	})

	t.Run("SuccessProjectGetStackformConfigExtraVars", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project", "get-env-config",
			"-p", "snowy", "-e", "sf-extra-vars",
		})

		assert.Nil(t, cmdErr)

		// Output should be in json by default
		var data = make(map[string]map[string]map[string]any)
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.NoError(t, err)

		message, ok := data["pipeline"]["config"]["message"]
		assert.True(t, ok)
		assert.Equal(t, "filledFromExtraVars", message)
	})

	t.Run("SuccessProjectsDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"delete",
			"--project", "snowy",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

}
