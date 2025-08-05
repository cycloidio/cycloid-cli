package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/custommodels"
	"github.com/matryer/is"
)

func TestStacks(t *testing.T) {
	t.Run("SuccessStacksList", func(t *testing.T) {
		is := is.New(t)
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"stacks",
			"list",
		})

		is.NoErr(cmdErr)
		var stackList []*models.ServiceCatalog
		err := json.Unmarshal([]byte(cmdOut), &stackList)
		is.NoErr(err)
		is.True(len(stackList) >= 1) // We should have at least one stack in our test org
		var testStackRef = config.Org + ":stack-e2e-stackforms"

		t.Run("SuccessStacksGet", func(t *testing.T) {
			is := is.New(t)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"stacks",
				"get",
				"--stack-ref", testStackRef,
			})
			is.NoErr(cmdErr)
			var outStack *models.ServiceCatalog
			err := json.Unmarshal([]byte(cmdOut), &outStack)
			is.NoErr(err)
		})

		t.Run("SuccessStacksUpdateVisibilty", func(t *testing.T) {
			is := is.New(t)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"stack",
				"update",
				"--stack-ref", testStackRef,
				"--visibility", "shared",
			})
			is.NoErr(cmdErr)

			var updatedStack *models.ServiceCatalog
			err := json.Unmarshal([]byte(cmdOut), &updatedStack)
			is.NoErr(err)
			is.Equal(*updatedStack.Visibility, "shared")
		})

		t.Run("SuccessAddStackMaintainer", func(t *testing.T) {
			var teamCanonical = "test-team"
			body := map[string]any{
				"canonical": teamCanonical,
				"name":      teamCanonical,
				"roles_canonical": []string{
					"default-project-viewer",
				},
			}
			jsonBody, err := json.Marshal(body)
			if err != nil {
				t.Fatalf("[preparation]: json serialization shouldn't fail: %s", err.Error())
			}

			// team management is not implemented on the CLI, so making the call ourselves
			request, err := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/teams", config.APIUrl, config.Org), bytes.NewBuffer(jsonBody))
			if err != nil {
				t.Fatalf("[preparation]: request creationg shoudn't fail: %s", err.Error())
			}

			request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))
			request.Header.Add("Content-Type", "application/vnd.cycloid.io.v1+json")

			client := &http.Client{}
			_, err = client.Do(request)
			if err != nil {
				t.Fatalf("[Preparation]: request to create teams shouldn't fail: %s", err.Error())
			}

			// At this point we should have a team, I assume CR and stacks are present too
			is := is.New(t)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"stack", "update",
				"--stack-ref", testStackRef,
				"--team", teamCanonical,
			})
			is.NoErr(cmdErr)
			var updatedStack *models.ServiceCatalog
			err = json.Unmarshal([]byte(cmdOut), &updatedStack)
			is.NoErr(err)
			is.Equal(*updatedStack.Team.Canonical, teamCanonical) // New team canonical must match
		})

		t.Run("SuccessRemoveMaintainer", func(t *testing.T) {
			is := is.New(t)
			cmdOut, cmdErr := executeCommand([]string{
				"--output", "json",
				"stack", "update",
				"--stack-ref", testStackRef,
				"--team", "", // setting the flag with empty string should remove the maintainer
			})
			is.NoErr(cmdErr) // This command must not fail
			var updatedStack *models.ServiceCatalog
			err := json.Unmarshal([]byte(cmdOut), &updatedStack)
			is.NoErr(err)                    // We should be able to deserialize a valid model
			is.Equal(updatedStack.Team, nil) // Team should be unset
		})

		t.Run("InvalidMaintainerShouldError", func(t *testing.T) {
			is := is.New(t)
			_, cmdErr := executeCommand([]string{
				"--output", "json",
				"stack", "update",
				"--stack-ref", testStackRef,
				"--team", "invalidteam",
			})
			is.True(cmdErr != nil) // CLI should output an error if we try to update a stack with a team that doesn't exists
		})
	})

	t.Run("SuccessStacksValidateForm", func(t *testing.T) {
		is := is.New(t)
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
		testForms, err := os.CreateTemp("", "test-stackforms.yml")
		if err != nil {
			t.Fatalf("setup failed: error while writing test forms at '%s'", testForms.Name())
		}
		testFormsPath := testForms.Name()
		WriteFile(testFormsPath, TestForms)
		defer os.Remove(testFormsPath)

		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"stacks",
			"forms", "validate",
			testFormsPath,
		})
		is.NoErr(cmdErr)
		is.Equal(cmdOut, "")
	})

	t.Run("SuccessStacksListWithBlueprintFlag", func(t *testing.T) {
		is := is.New(t)
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"stacks",
			"list",
			"--blueprint",
		})
		is.NoErr(cmdErr) // cmd should not fail

		var blueprints []*custommodels.Blueprint
		err := json.Unmarshal([]byte(cmdOut), &blueprints)
		is.NoErr(err) // json output should be deserializable
	})

	t.Run("SuccessCreateStackFromBlueprint", func(t *testing.T) {
		t.Skip("Skipping due to missing way to cleanup stack afterwards")
	})
}
