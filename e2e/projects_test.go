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

	t.Run("SuccessProjectsCreateStdin", func(t *testing.T) {
		cmdOut, cmdErr := executeCommandWithStdin([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-stackforms-env",
			"--project", "snowy",
			"--env", "testStackformsStdin",
			"--use-case", "default",
			"-f", "-",
		}, `{ "pipeline": { "config": { "message": "filledFromStdin" } } }`)

		assert.Nil(t, cmdErr)
		var data = new(models.Project)
		err := json.Unmarshal([]byte(cmdOut), data)
		assert.NoError(t, err)

		var found = false
		for _, env := range data.Environments {
			if *env.Canonical == "testStackformsStdin" {
				found = true
			}
		}

		if !found {
			t.Errorf("testStackformsStdin not found in create project output")
		}
	})

	t.Run("SuccessProjectGetConfigAsJSON", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"get-env-config",
			"--project", "snowy",
			"--env", "testStackformsStdin",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "message\": \"filledFromStdin")
	})

	t.Run("SuccessProjectsCreateStdin", func(t *testing.T) {
		cmdOut, cmdErr := executeCommandWithStdin([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-stackform-env",
			"--project", "snowy",
			"--env", "testStackformsStdin",
			"-f", "-",
		}, `{ "pipeline": { "config": { "message": "filledFromStdin" } } }`)

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
		var data models.Project
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.Nil(t, err)
	})

	t.Run("SuccessProjectGetStackformEnvStdin", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project", "get-env-config",
			"-p", "snowy", "-e", "testStackformsStdin",
		})

		assert.Nil(t, cmdErr)

		// Output should be in json by default
		var data = make(map[string]map[string]map[string]any)
		err := json.Unmarshal([]byte(cmdOut), &data)
		assert.NoError(t, err)

		message, ok := data["pipeline"]["config"]["message"]
		assert.True(t, ok)
		assert.Equal(t, "filledFromStdin", message)
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
