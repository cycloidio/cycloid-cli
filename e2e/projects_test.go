// //go:build e2e
// // +build e2e
package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

		// Add the CLI test catalog
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"catalog-repository",
			"create",
			"--branch", "stacks",
			"--url", "https://github.com/cycloidio/cycloid-cli-test-catalog.git",
			"--name", "cli-test",
		})

		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"catalog-repository",
			"refresh",
			"--canonical", "cli-test",
		})

		// Ensure the catalog is present
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"catalog-repository",
			"get",
			"--canonical", "cli-test",
		})

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"cli-test")

		// Here is an example if you want to add a specific catalog.
		// Since the latest update we have by default all the public stacks
		// Provide service catalog public
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

	t.Run("CreateProjectWithEnvShouldFail", func(t *testing.T) {
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
			"create-env",
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

	t.Run("SuccessCreateEnvLegacy", func(t *testing.T) {
		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)

		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"delete",
			"--project", "snowy-legacy",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create",
			"--name", "snowy-legacy",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", CY_TEST_ROOT_ORG),
			"--config-repo", "project-config",
			"--output", "json",
		})

		assert.Nil(t, cmdErr, "project creation should have succeeded: ", cmdOut)

		cmdOut, cmdErr = executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create-env",
			"--project", "snowy-legacy",
			"--env", "test",
			"--pipeline", "/tmp/test_cli-pp",
			"--vars", "/tmp/test_cli-pp-vars",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
			"--use-case", "default",
		})

		assert.Nil(t, cmdErr, "createEnv should handle legacy env creation, error: ", cmdOut)
	})

	t.Run("StackformsTestSetup", func(t *testing.T) {
		var project = "stackforms-tests"

		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"delete",
			"--project", project,
		})

		// Setup
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", CY_TEST_ROOT_ORG,
			"project",
			"create",
			"--name", project,
			"--description", "test of stackforms update values",
			"--stack-ref", fmt.Sprintf("%s:stack-e2e-stackforms", CY_TEST_ROOT_ORG),
			"--config-repo", "project-config",
			"--output", "json",
		})

		assert.Nil(t, cmdErr, "project creation should have succeeded: ", cmdOut)
	})

	type TestCase struct {
		Input  string
		Output interface{}
	}

	var formsTestCases = map[string]TestCase{
		"string": {
			Input:  "my-string",
			Output: "my-string",
		},
		"integer": {
			Input:  "1",
			Output: 1,
		},
		"float": {
			Input:  "1.1",
			Output: 1.1,
		},
		"bool": {
			Input:  "true",
			Output: true,
		},
		"bool-caps": {
			Input:  "True",
			Output: true,
		},
		"map": {
			Input: `{"myString": "string", "myBool": true, "myInt": 1, "myFloat", 1.1, "myNested": {"hello", "world"}, "myArray": ["hello", "world"]}`,
			Output: map[string]any{
				"myString": "string",
				"myBool":   true,
				"myInt":    1,
				"myFloat":  1.1,
				"myNested": map[string]string{
					"hello": "world",
				},
				"myArray": []string{"hello", "world"},
			},
		},
		"array": {
			Input: `["string", 1, 1.1, true, ["hello", "world"], {"hello": "world"}]`,
			Output: []any{
				"string",
				1,
				1.1,
				true,
				[]string{"hello", "world"},
				map[string]string{"hello": "world"},
			},
		},
	}

	for _, method := range []string{"VarFlag", "JsonVars", "VarFiles"} {
		t.Run("StackformsTest"+method, func(t *testing.T) {
			var project = "stackforms-tests"

			for testKey, testValue := range formsTestCases {
				var key string
				if testKey == "bool-caps" {
					key = "bool"
				} else {
					key = testKey
				}

				// This section is for Json vars handling
				var inputAsJsonValue string
				if key == "string" {
					// since all testValue.Input are string, we need a simple way to encore json
					// as it would be from the CLI or a file
					// using json.Marshal() would not work since all Input are strings
					// So we just have to add quotes if the Input value is a string.
					inputAsJsonValue = fmt.Sprintf(`"%s"`, testValue.Input)
				} else {
					inputAsJsonValue = testValue.Input
				}

				jsonInput := fmt.Sprintf(`{"types": { "tests": { "%s" : %s } } }`, key, inputAsJsonValue)

				var extraFlag []string
				switch method {
				case "VarFlag":
					extraFlag = []string{"-V", fmt.Sprintf("%s=%s", key, testValue.Input)}

				case "JsonVars":
					extraFlag = []string{"--json-vars", jsonInput}

				case "JsonEnvVars":
					t.Setenv("CY_STACKFORMS_VARS", jsonInput)

				case "VarFiles":
					err := os.WriteFile("/tmp/jsonVar.json", []byte(jsonInput+"\n"), 0664)
					assert.Nil(t, err, "tests must be able to write to /tmp")

					extraFlag = []string{"-f", "/tmp/jsonVar.json"}
				}

				cmd := append([]string{
					"--output", "json",
					"--org", CY_TEST_ROOT_ORG,
					"project",
					"create-env",
					"--project", project,
					"--env", fmt.Sprintf("%s-%s", strings.ToLower(method), key), // One env per type
					"--use-case", "default",
					"--update",
				}, extraFlag...)

				cmdOut, cmdErr := executeCommand(cmd)
				assert.Nil(t, cmdErr, "create env should not fail", cmdOut)

				// Parse json output to ensure our value has been taken by backend
				var data = make(map[string]map[string]map[string]any)
				err := json.Unmarshal([]byte(cmdOut), &data)
				assert.NoError(t, err, "we should be able to serialize response as JSON\n", "out:\n", cmdOut, "err:\n", cmdErr)
				assert.Equal(t, testValue.Output, data["types"]["tests"][key], "response should match the expected type output, cli output: ", cmdOut)
			}
		})
	}
}
