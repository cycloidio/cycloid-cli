package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStacks(t *testing.T) {
	t.Skip()

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
		var TestGitSshKey = []byte(`---
default:
  pipeline:
    AWS:
      - name: "Default Region"
        key: aws_default_region
        type: string
        widget: dropdown
        description: "In which region you would like your project to run"
        default: "eu-west-1"
        values: ["eu-west-1", "eu-west-2", "eu-west3", "eu-south1", "eu-north1", "eu-central1"]
        required: true
`)
		WriteFile("/tmp/test_ci_form", TestGitSshKey)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"validate-form",
			"--forms", "/tmp/test_ci_form",
		})

		require.Nil(t, cmdErr)
		assert.Contains(t, cmdOut, "errors\": []")
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
}
