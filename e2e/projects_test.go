// //go:build e2e
// // +build e2e
package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

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

	t.Run("UnAuthorizedCreateProjectWithEnv", func(t *testing.T) {
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)

		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"delete",
			"--project", "snowy-invalid",
		})

		_, cmdErr := executeCommand([]string{
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

		assert.ErrorContains(t, cmdErr, "Creating an environment when creating a project is not possible anymore", "Creating a project with an env is prohibited now.")
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
			"--output", "json",
		})

		assert.Nil(t, cmdErr)

		var expectedData map[string]any
		err := json.Unmarshal([]byte(cmdOut), &expectedData)
		assert.Nil(t, err, "whe should be able to parse json output")
		require.Equal(t,
			"snowy",
			expectedData["canonical"],
			"project canonical should be in json output: ", cmdOut)
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
			"-j", `{"pipeline": {"config": {"message": "filledFromVars"}}}`,
		})

		assert.Nil(t, cmdErr)
		var data map[string]map[string]map[string]string
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.Nil(t, err)
		assert.Equal(t, "filledFromVars", data["pipeline"]["config"]["message"])
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
			"-V", `pipeline.config.message=filledFromExtraVars`,
		})

		assert.Nil(t, cmdErr)
		var data map[string]map[string]map[string]string
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.Nil(t, err)
		assert.Equal(t, "filledFromExtraVars", data["pipeline"]["config"]["message"])
	})

	// Extra vars
	t.Run("SuccessProjectsCreateExtraVarsUpdate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-stackforms-env",
			"--project", "snowy",
			"--env", "sf-extra-vars",
			"--use-case", "default",
			"-V", `pipeline.config.message=filledFromExtraVars`,
			"-V", `pipeline.config.message=filledFromExtraVars2`,
			"--update",
		})

		assert.Nil(t, cmdErr)
		var data map[string]map[string]map[string]string
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.Nil(t, err)
		assert.Equal(t, "filledFromExtraVars2", data["pipeline"]["config"]["message"])
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
