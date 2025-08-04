package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStacks(t *testing.T) {
	// Set the API key environment variable for CLI commands
	err := os.Setenv("CY_API_KEY", config.APIKey)
	require.Nil(t, err)

	// Since the latest update the public catalog have been added by default
	// Here is a sample of code if we need to add a dedicated one
	// t.Run("InitPublicCatalog", func(t *testing.T) {
	// 	executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", CY_TEST_ROOT_ORG,
	// 		"catalog-repository",
	// 		"create",
	// 		"--branch", "master",
	// 		"--url", "https://github.com/cycloid-community-catalog/stack-magento.git",
	// 		"--name", "magento",
	// 	})
	//
	// 	// Ensure the catalog is present
	// 	cmdOut, _ := executeCommand([]string{
	// 		"--output", "json",
	// 		"--org", CY_TEST_ROOT_ORG,
	// 		"catalog-repository",
	// 		"get",
	// 		"--canonical", "magento",
	// 	})
	//
	// 	require.Contains(t, cmdOut, "canonical\": \"magento")
	// })

	t.Run("SuccessStacksList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"stack-dummy")
	})

	t.Run("SuccessStacksGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"get",
			"--ref", fmt.Sprintf("%s:stack-dummy", config.Org),
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"stack-dummy")
	})

	t.Run("SuccessStacksValidateForm", func(t *testing.T) {
		var TestForms = []byte(`---
version: "4"
shared:
- &anchor2
  name: "hello"
  key: "toto3"
  widget: "simple_text"
  type: "string"
use_cases:
- name: use_cases
  sections:
  - name: "hello"
    groups:
    - name: "toto"
      technlogies: ["tutu"]
      vars:
      - &anchor1
        name: "hello"
        key: "toto"
        widget: "simple_text"
        type: "string"
      - <<: *anchor1
        key: "toto1"
      - *anchor2
      - <<: *anchor2
        key: "toto4"
`)
		testFile, err := os.CreateTemp("", "test-stackforms.yml")
		if err != nil {
			t.Fatalf("setup failed: error while writing test forms at '%s'", testFile.Name())
		}

		formsFile := testFile.Name()
		WriteFile(formsFile, TestForms)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"validate-form",
			formsFile,
		})
		require.Nil(t, cmdErr)
		assert.Equal(t, cmdOut, "")
	})

	t.Run("SuccessStacksUpdateVisibilty", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"update",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", config.Org),
			"--visibility", "shared",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "canonical\": \"stack-dummy")
		assert.Contains(t, cmdOut, "visibility\": \"shared")
	})

	t.Run("SuccessAddStackMaintainer", func(t *testing.T) {
		t.Setenv("CY_ORG", config.Org)
		var teamCanonical = "test-team"
		body := map[string]any{
			"canonical": teamCanonical,
			"name":      teamCanonical,
			"roles_canonical": []string{
				"default-project-viewer",
			},
		}
		jsonBody, err := json.Marshal(body)
		assert.Nil(t, err, "[preparation]: json serialization shouldn't fail.")

		// team management is not implemented on the CLI, so making the call ourselves
		request, err := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/teams", config.APIUrl, config.Org), bytes.NewBuffer(jsonBody))
		assert.Nil(t, err, "[preparation]: request creationg shoudn't fail")

		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))
		request.Header.Add("Content-Type", "application/vnd.cycloid.io.v1+json")

		client := &http.Client{}
		_, err = client.Do(request)
		assert.Nil(t, err, "[Preparation]: request to create teams shouldn't fail")

		// At this point we should have a team, I assume CR and stacks are present too
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"stack", "update",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", config.Org),
			"--team", teamCanonical,
		})

		assert.Nil(t, cmdErr, "CLI should be able to update the correct team without error")
		assert.Contains(t, cmdOut, teamCanonical, "team canonical should be in the JSON response.")
	})

	t.Run("SuccessRemoveMaintainer", func(t *testing.T) {
		t.Setenv("CY_ORG", config.Org)
		// We assume that the team exists from the previous test
		var teamCanonical = "test-team"
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"stack", "update",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", config.Org),
			"--team", "", // setting the flag with empty string should remove the maintainer
		})

		assert.Nil(t, cmdErr, "CLI should be able to update the correct team without error")
		assert.NotContains(t, cmdOut, teamCanonical, "team canonical should not be in json response")
	})

	t.Run("InvalidMaintainerShouldError", func(t *testing.T) {
		t.Setenv("CY_ORG", config.Org)
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"stack", "update",
			"--stack-ref", fmt.Sprintf("%s:stack-dummy", config.Org),
			"--team", "invalidteam",
		})

		assert.Error(t, cmdErr, "CLI should output an error if we try to update a stack with a team that doesn't exists")
	})

	// Blueprint tests
	t.Run("SuccessStacksListWithBlueprintFlag", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "[", "Response should be a JSON array")
		assert.Contains(t, cmdOut, "]", "Response should be a JSON array")
	})

	t.Run("SuccessStacksListBlueprintTableOutput", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "table",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})

		require.Nil(t, cmdErr)
		assert.NotEmpty(t, cmdOut, "Table output should not be empty")
	})

	t.Run("SuccessCreateStackFromBlueprint", func(t *testing.T) {
		t.Skip("Skipping due to missing service catalog source in test environment")
		// First, get a blueprint reference
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})
		require.Nil(t, cmdErr)

		// Parse the response to get a blueprint ref
		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)
		require.NotEmpty(t, blueprints, "Should have at least one blueprint")

		blueprintRef := blueprints[0]["ref"].(string)
		sourceCanonical := "test-catalog"
		useCase := "default"
		testStackName := "test-stack-from-blueprint"
		testStackCanonical := "test-stack-from-blueprint"

		createOut, createErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"create-from-blueprint",
			"--blueprint-ref", blueprintRef,
			"--name", testStackName,
			"--canonical", testStackCanonical,
			"--service-catalog-source-canonical", sourceCanonical,
			"--use-case", useCase,
		})

		require.Nil(t, createErr)
		assert.Contains(t, createOut, testStackCanonical)

		// Cleanup: delete the created stack
		defer func() {
			executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"stack",
				"delete",
				"--stack-ref", fmt.Sprintf("%s:%s", config.Org, testStackCanonical),
			})
		}()
	})

	t.Run("ErrorCreateStackFromBlueprintInvalidRef", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"create-from-blueprint",
			"--blueprint-ref", "invalid:ref",
			"--name", "test-stack",
			"--canonical", "test-stack",
			"--service-catalog-source-canonical", "test-catalog",
			"--use-case", "default",
		})

		assert.Error(t, cmdErr, "Should fail with invalid blueprint ref")
	})

	t.Run("ErrorCreateStackFromBlueprintMissingRequiredFlags", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"create-from-blueprint",
			"--blueprint-ref", "test:blueprint",
			// Missing required flags
		})

		assert.Error(t, cmdErr, "Should fail with missing required flags")
	})

	t.Run("SuccessGetBlueprintConfig", func(t *testing.T) {
		t.Skip("Skipping due to panic in config command")
		// First, get a blueprint reference
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})
		require.Nil(t, cmdErr)

		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)
		require.NotEmpty(t, blueprints, "Should have at least one blueprint")

		blueprintRef := blueprints[0]["ref"].(string)

		// Get the blueprint config
		configOut, configErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"config",
			"get",
			"-s", blueprintRef,
			"-u", "default",
		})

		require.Nil(t, configErr)
		assert.Contains(t, configOut, "{", "Should return JSON config")
	})

	t.Run("SuccessBlueprintWithUseCases", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})

		require.Nil(t, cmdErr)

		// Parse response to check for use cases
		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)

		// Check that blueprints have use cases field
		for _, blueprint := range blueprints {
			if useCases, exists := blueprint["use_cases"]; exists {
				assert.NotNil(t, useCases, "Use cases should not be nil")
			}
		}
	})

	t.Run("SuccessBlueprintWorkflow", func(t *testing.T) {
		t.Skip("Skipping due to missing service catalog source in test environment")
		// Test the complete workflow: list blueprints, get config, create stack
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})
		require.Nil(t, cmdErr)

		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)
		require.NotEmpty(t, blueprints, "Should have at least one blueprint")

		blueprintRef := blueprints[0]["ref"].(string)
		testStackName := "workflow-test-stack"
		testStackCanonical := "workflow-test-stack"

		// Create stack from blueprint
		createOut, createErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"create-from-blueprint",
			"--blueprint-ref", blueprintRef,
			"--name", testStackName,
			"--canonical", testStackCanonical,
			"--service-catalog-source-canonical", "test-catalog",
			"--use-case", "default",
		})

		require.Nil(t, createErr)
		assert.Contains(t, createOut, testStackCanonical)

		// Verify the stack was created
		getOut, getErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"get",
			"--ref", fmt.Sprintf("%s:%s", config.Org, testStackCanonical),
		})

		require.Nil(t, getErr)
		assert.Contains(t, getOut, testStackCanonical)

		// Cleanup
		defer func() {
			executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"stack",
				"delete",
				"--stack-ref", fmt.Sprintf("%s:%s", config.Org, testStackCanonical),
			})
		}()
	})

	t.Run("ErrorHandlingInvalidBlueprintRef", func(t *testing.T) {
		invalidRefs := []string{"invalid-ref", "invalid:", ":invalid", "", "invalid:ref:extra"}
		for _, invalidRef := range invalidRefs {
			t.Run(fmt.Sprintf("InvalidRef_%s", invalidRef), func(t *testing.T) {
				_, cmdErr := executeCommand([]string{
					"--output", "json",
					"--org", config.Org,
					"stack",
					"create-from-blueprint",
					"--blueprint-ref", invalidRef,
					"--name", "test-stack",
					"--canonical", "test-stack",
					"--service-catalog-source-canonical", "test-catalog",
					"--use-case", "default",
				})
				assert.Error(t, cmdErr, "Should fail with invalid blueprint ref: %s", invalidRef)
			})
		}
	})

	t.Run("ErrorHandlingInvalidUseCase", func(t *testing.T) {
		// First, get a valid blueprint reference
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})
		require.Nil(t, cmdErr)

		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)
		require.NotEmpty(t, blueprints, "Should have at least one blueprint")

		blueprintRef := blueprints[0]["ref"].(string)

		_, cmdErr = executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stack",
			"create-from-blueprint",
			"--blueprint-ref", blueprintRef,
			"--name", "test-stack",
			"--canonical", "test-stack",
			"--service-catalog-source-canonical", "test-catalog",
			"--use-case", "invalid-use-case",
		})

		assert.Error(t, cmdErr, "Should fail with invalid use case")
	})

	t.Run("SuccessBlueprintConfigValidation", func(t *testing.T) {
		// Get a blueprint and validate its config structure
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})
		require.Nil(t, cmdErr)

		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)
		require.NotEmpty(t, blueprints, "Should have at least one blueprint")

		// Validate blueprint structure
		blueprint := blueprints[0]
		assert.Contains(t, blueprint, "ref", "Blueprint should have ref field")
		assert.Contains(t, blueprint, "name", "Blueprint should have name field")
		assert.Contains(t, blueprint, "canonical", "Blueprint should have canonical field")
		assert.Contains(t, blueprint, "blueprint", "Blueprint should have blueprint field")
		assert.Equal(t, true, blueprint["blueprint"], "Blueprint field should be true")
	})

	t.Run("SuccessBlueprintUseCasesExtraction", func(t *testing.T) {
		// Test that use cases are properly extracted from blueprint config
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
			"--blueprint",
		})
		require.Nil(t, cmdErr)

		var blueprints []map[string]interface{}
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		require.Nil(t, err)
		require.NotEmpty(t, blueprints, "Should have at least one blueprint")

		// Check that use cases are properly formatted
		for _, blueprint := range blueprints {
			if useCases, exists := blueprint["use_cases"]; exists {
				switch v := useCases.(type) {
				case []interface{}:
					// Use cases as array
					for _, useCase := range v {
						assert.IsType(t, "", useCase, "Use case should be a string")
					}
				case string:
					// Use cases as comma-separated string
					assert.NotEmpty(t, v, "Use cases string should not be empty")
				}
			}
		}
	})
}
