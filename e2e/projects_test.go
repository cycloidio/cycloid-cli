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
	t.Skip()
	LoginToRootOrg()

	// Cleanup previous project if exist and prepare required catalog repository, ...
	t.Run("CleanupAndPrepare", func(t *testing.T) {
		// TODO: Fix tests when components are implemented
		t.Skip()

		// Create ssh cred
		WriteFile("/tmp/test_cli-ssh", TestGitSshKey)
		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"creds",
			"create",
			"ssh",
			"--name", "git-project-creds",
			"--ssh-key", "/tmp/test_cli-ssh",
		})

		// Create config repo
		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"config-repo",
			"create",
			"--name", "project-config",
			"--branch", CyTestCatalogRepoBranch,
			"--cred", "git-project-creds",
			"--url", CyTestCatalogRepoURL,
		})

		// Add the CLI test catalog
		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"catalog-repository",
			"create",
			"--branch", "stacks",
			"--url", "https://github.com/cycloidio/cycloid-cli-test-catalog.git",
			"--name", "cli-test",
		})

		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"catalog-repository",
			"refresh",
			"--canonical", "cli-test",
		})

		// Ensure the catalog is present
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
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
			"--org", TestRootOrg,
			"project",
			"delete",
			"--project", "snowy-invalid",
		})

		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"create",
			"--name", "snowy",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", TestRootOrg),
			"--config-repo", "project-config",
			"--env", "test",
			"--usecase", "default",
			"--vars", "/tmp/test_cli-pp-vars",
			"--pipeline", "/tmp/test_cli-pp",
			"--config", "/tmp/test_cli-pp=/snowy/test/test_cli-pp",
		})

		// TODO: Fix tests when components are implemented
		t.Skip()

		assert.ErrorContains(t, cmdErr, "Creating an environment when creating a project is not possible anymore", "Creating a project with an env is prohibited now.")
	})

	t.Run("SuccessLegacyProjectsCreate", func(t *testing.T) {
		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"delete",
			"--project", "snowy",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"create",
			"--name", "snowy",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", TestRootOrg),
			"--config-repo", "project-config",
			"--output", "json",
		})

		// TODO: Fix tests when components are implemented
		t.Skip()

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
			"--org", TestRootOrg,
			"project",
			"list",
		})

		// TODO: Fix tests when components are implemented
		t.Skip()

		assert.Nil(t, cmdErr)
		require.Contains(t, cmdOut, "canonical\": \"snowy")
	})

	// Vars
	t.Run("SuccessProjectsCreateVars", func(t *testing.T) {
		// TODO: Fix tests when components are implemented
		t.SkipNow()

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
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
		// TODO: Fix tests when components are implemented
		t.Skip()

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
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
		// TODO: Fix tests when components are implemented
		t.Skip()

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
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
		// TODO: Fix tests when components are implemented
		t.Skip()

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
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
		// TODO: Fix tests when components are implemented
		t.Skip()

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"delete",
			"--project", "snowy",
		})

		assert.Nil(t, cmdErr)
		require.Equal(t, "", cmdOut)
	})

	t.Run("SuccessCreateEnvLegacy", func(t *testing.T) {
		// TODO: Fix tests when components are implemented
		t.Skip()

		WriteFile("/tmp/test_cli-pp-vars", TestPipelineVariables)
		WriteFile("/tmp/test_cli-pp", TestPipelineSample)

		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"delete",
			"--project", "snowy-legacy",
		})

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"create",
			"--name", "snowy-legacy",
			"--description", "this is a test project",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", TestRootOrg),
			"--config-repo", "project-config",
			"--output", "json",
		})

		assert.Nil(t, cmdErr, "project creation should have succeeded: ", cmdOut)

		cmdOut, cmdErr = executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
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
		// TODO: Fix tests when components are implemented
		t.Skip()

		var project = "stackforms-tests"

		// Cleanup
		executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"delete",
			"--project", project,
		})

		// Setup
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", TestRootOrg,
			"project",
			"create",
			"--name", project,
			"--description", "test of stackforms update values",
			"--stack-ref", fmt.Sprintf("%s:stack-e2e-stackforms", TestRootOrg),
			"--config-repo", "project-config",
			"--output", "json",
		})

		// TODO: Fix tests when components are implemented
		t.Skip()

		assert.Nil(t, cmdErr, "project creation should have succeeded: ", cmdOut)
	})

	type TestCase struct {
		Keys   []string
		Type   string
		Input  string
		Output interface{}
	}

	var formsTestCases = []TestCase{
		{
			Keys:   []string{"types", "tests", "string"},
			Type:   "string",
			Input:  "my-string",
			Output: "my-string",
		},
		{
			Keys:   []string{"types", "tests", "integer"},
			Type:   "integer",
			Input:  "1",
			Output: 1,
		},
		{
			Keys:   []string{"types", "tests", "integer"},
			Type:   "integer",
			Input:  "-29291",
			Output: -29291,
		},
		{
			Keys:   []string{"types", "tests", "float"},
			Type:   "float",
			Input:  "1.1",
			Output: 1.1,
		},
		{
			Keys:   []string{"types", "tests", "float"},
			Type:   "float",
			Input:  "-666.423",
			Output: -666.423,
		},
		{
			Keys:   []string{"types", "tests", "bool"},
			Type:   "bool",
			Input:  "true",
			Output: true,
		},
		{
			Keys:   []string{"types", "tests", "bool"},
			Type:   "bool",
			Input:  "True",
			Output: true,
		},
		{
			Keys:   []string{"types", "tests", "string"},
			Type:   "null",
			Input:  "null",
			Output: "stringValue", // Since we send nil, the forms will receive the default value.
		},
		{
			Keys:  []string{"types", "tests", "map"},
			Type:  "map",
			Input: `{"myString": "string", "myBool": true, "myInt": 1, "myFloat": 1.1, "myNested": {"hello": "world"}, "myArray": ["hello", "world"]}`,
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
		{
			Keys:  []string{"types", "tests", "array"},
			Type:  "array",
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
		{
			Keys:   []string{"section spaces AND CAPS", "group spaces AND CAPS", "no_spaces_no_caps"},
			Type:   "string",
			Input:  "thisIsAStringYay",
			Output: "thisIsAStringYay",
		},
		{
			Keys:   []string{"can two sections have same name with different caps ?", "can two groups have same name with different caps ?", "group1"},
			Type:   "string",
			Input:  "lowkey",
			Output: "lowkey",
		},
		{
			Keys:   []string{"CAN TWO SECTIONS HAVE SAME NAME WITH DIFFERENT CAPS ?", "CAN TWO GROUPS HAVE SAME NAME WITH DIFFERENT CAPS ?", "group2"},
			Type:   "string",
			Input:  "SHOUT",
			Output: "SHOUT",
		},
	}

	for _, method := range []string{"VarFlag", "JsonVars", "VarFiles"} {
		t.Run("StackformsTest"+method, func(t *testing.T) {
			var project = "stackforms-tests"

			// TODO: Fix tests when components are implemented
			t.Skip()

			for _, testValue := range formsTestCases {

				// This section is for Json vars handling
				var inputAsJsonValue string
				switch testValue.Type {
				case "string":
					// since all testValue.Input are string, we need a simple way to encore json
					// as it would be from the CLI or a file
					// using json.Marshal() would not work since all Input are strings
					// So we just have to add quotes if the Input value is a string.
					inputAsJsonValue = fmt.Sprintf(`"%s"`, testValue.Input)

				case "bool":
					// Lower the capital bool test case as it would be invalid JSON
					inputAsJsonValue = fmt.Sprintf(`%s`, strings.ToLower(testValue.Input))

				default:
					inputAsJsonValue = testValue.Input
				}

				// Lazy way to just build the JSON string as it would be inputted via the CLI
				var stringJsonInput string
				keysLen := len(testValue.Keys)
				for index, key := range testValue.Keys {
					if index < keysLen-1 {
						stringJsonInput += fmt.Sprintf(`{"%s": `, key)
						continue
					}

					// Close the json
					stringJsonInput += fmt.Sprintf(`{"%s": %s %s`, key, inputAsJsonValue, strings.Repeat("}", keysLen))
				}

				var extraFlag []string
				switch method {
				case "VarFlag":
					if testValue.Type == "string" {
						extraFlag = []string{"-V", fmt.Sprintf(`%s=%s`, strings.Join(testValue.Keys, "."), testValue.Input)}
					} else {
						extraFlag = []string{"-V", fmt.Sprintf("%s=%s", strings.Join(testValue.Keys, "."), testValue.Input)}
					}

				case "JsonVars":
					extraFlag = []string{"--json-vars", stringJsonInput}

				case "JsonEnvVars":
					t.Setenv("CY_STACKFORMS_VARS", stringJsonInput)

				case "VarFiles":
					filename := fmt.Sprintf("/tmp/cli-%s-%s.json", method, testValue.Type)
					err := os.WriteFile(filename, []byte(stringJsonInput+"\n"), 0664)
					assert.Nil(t, err, "tests must be able to write to ", filename)

					extraFlag = []string{"-f", filename}
				}

				cmd := append([]string{
					"--output", "json",
					"--org", TestRootOrg,
					"project",
					"create-env",
					"--project", project,
					"--env", fmt.Sprintf("%s-%s", strings.ToLower(method), testValue.Type), // One env per type
					"--use-case", "default",
					"--update",
				}, extraFlag...)

				cmdOut, cmdErr := executeCommand(cmd)
				assert.Nil(t, cmdErr, "create env should not fail", cmdOut)

				// Parse json output to ensure our value has been taken by backend
				var data = make(map[string]map[string]map[string]any)
				err := json.Unmarshal([]byte(cmdOut), &data)
				assert.NoError(t, err, "we should be able to serialize response as JSON\n", "out:\n", cmdOut, "err:\n", cmdErr)

				cliResult := data[testValue.Keys[0]][testValue.Keys[1]][testValue.Keys[2]]

				switch testValue.Type {
				case "string", "bool", "float", "integer", "null":
					assert.EqualValues(t, testValue.Output, cliResult, fmt.Sprintf("%v", testValue))
				case "map", "array":
					jsonResult, err := json.MarshalIndent(cliResult, "", "  ")
					assert.NoErrorf(t, err, "CLI JSON output, for key types.tests.%s should be json serializable", testValue.Type)

					jsonExpect, err := json.MarshalIndent(cliResult, "", "  ")
					assert.NoError(t, err, "the expected test output should be JSON serializable")

					assert.JSONEq(t, string(jsonExpect), string(jsonResult), "The two json should be the same.")
				}
			}
		})
	}
}
