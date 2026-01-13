package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/matryer/is"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/custommodels"
)

func TestStacks(t *testing.T) {
	t.Run("SuccessStacksListJSON", func(t *testing.T) {
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
	})

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
			t.Errorf("[preparation]: json serialization shouldn't fail: %s", err.Error())
		}

		// team management is not implemented on the CLI, so making the call ourselves
		request, err := http.NewRequest("POST", fmt.Sprintf("%s/organizations/%s/teams", config.APIUrl, config.Org), bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Errorf("[preparation]: request creationg shoudn't fail: %s", err.Error())
		}

		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))
		request.Header.Add("Content-Type", "application/vnd.cycloid.io.v1+json")

		client := &http.Client{}
		_, err = client.Do(request)
		if err != nil {
			t.Errorf("[Preparation]: request to create teams shouldn't fail: %s", err.Error())
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
			t.Errorf("setup failed: error while writing test forms at '%s'", testForms.Name())
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

	// Completion tests
	t.Run("SuccessCompleteStackVersionCommitHash", func(t *testing.T) {
		is := is.New(t)
		// Test completion for component create with stack-ref
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "create",
			"--org", config.Org,
			"--stack-ref", testStackRef,
			"--stack-commit-hash", "",
		})
		is.NoErr(cmdErr)
		// The output should contain commit hashes
		// Format: commitHash\tDescription\n:4\nCompletion ended with directive: ShellCompDirectiveNoFileComp
		is.True(len(cmdOut) > 0)                                        // Should have some output
		is.True(!bytes.Contains([]byte(cmdOut), []byte("_activeHelp"))) // Should not have error help
		is.True(bytes.Contains([]byte(cmdOut), []byte(":4")))           // Should have NoFileComp directive (4)
	})

	t.Run("SuccessCompleteStackVersionCommitHashWithPartialInput", func(t *testing.T) {
		is := is.New(t)
		// First, get a valid commit hash from the stack versions
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "create",
			"--org", config.Org,
			"--stack-ref", testStackRef,
			"--stack-commit-hash", "",
		})
		is.NoErr(cmdErr)

		// Parse the first commit hash (before the tab character)
		lines := bytes.Split([]byte(cmdOut), []byte("\n"))
		if len(lines) > 0 && len(lines[0]) > 3 {
			// Take first 3 characters as partial input
			partial := string(lines[0][:3])

			// Test completion with partial input
			cmdOut2, cmdErr2 := executeCommand([]string{
				"__complete",
				"component", "create",
				"--org", config.Org,
				"--stack-ref", testStackRef,
				"--stack-commit-hash", partial,
			})
			is.NoErr(cmdErr2)
			is.True(len(cmdOut2) > 0)                                        // Should have some output
			is.True(!bytes.Contains([]byte(cmdOut2), []byte("_activeHelp"))) // Should not have error help
			is.True(bytes.Contains([]byte(cmdOut2), []byte(partial)))        // Should contain the partial string
		}
	})

	t.Run("FailCompleteStackVersionCommitHashNoStackRef", func(t *testing.T) {
		is := is.New(t)
		// Test completion without stack-ref should provide helpful error
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "create",
			"--org", config.Org,
			"--stack-commit-hash", "",
		})
		// The command itself doesn't error, but completion should show active help
		is.NoErr(cmdErr)
		is.True(bytes.Contains([]byte(cmdOut), []byte("_activeHelp")))       // Should have error help
		is.True(bytes.Contains([]byte(cmdOut), []byte("missing stack-ref"))) // Should mention missing stack-ref
	})

	t.Run("SuccessCompleteStackVersionCommitHashForUpdate", func(t *testing.T) {
		is := is.New(t)
		// For component update, if we provide stack-ref it should work
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "update",
			"--org", config.Org,
			"--stack-ref", testStackRef,
			"--stack-commit-hash", "",
		})
		is.NoErr(cmdErr)
		is.True(len(cmdOut) > 0)                                        // Should have some output
		is.True(!bytes.Contains([]byte(cmdOut), []byte("_activeHelp"))) // Should not have error help
		is.True(bytes.Contains([]byte(cmdOut), []byte(":4")))           // Should have NoFileComp directive
	})

	// Stack version tag completion tests
	t.Run("SuccessCompleteStackVersionTag", func(t *testing.T) {
		is := is.New(t)
		// Test completion for component create with stack-ref
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "create",
			"--org", config.Org,
			"--stack-ref", testStackRef,
			"--stack-tag", "",
		})
		is.NoErr(cmdErr)
		// The output should contain version tags
		is.True(len(cmdOut) > 0)                                        // Should have some output
		is.True(!bytes.Contains([]byte(cmdOut), []byte("_activeHelp"))) // Should not have error help
		is.True(bytes.Contains([]byte(cmdOut), []byte(":4")))           // Should have NoFileComp directive (4)
	})

	t.Run("SuccessCompleteStackVersionTagWithPartialInput", func(t *testing.T) {
		is := is.New(t)
		// First, get a valid version tag from the stack versions
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "create",
			"--org", config.Org,
			"--stack-ref", testStackRef,
			"--stack-tag", "",
		})
		is.NoErr(cmdErr)

		// Parse the first version tag (before the tab character)
		lines := bytes.Split([]byte(cmdOut), []byte("\n"))
		if len(lines) > 0 && len(lines[0]) > 0 {
			// Take first character as partial input
			parts := bytes.Split(lines[0], []byte("\t"))
			if len(parts) > 0 && len(parts[0]) > 0 {
				partial := string(parts[0][:1])

				// Test completion with partial input
				cmdOut2, cmdErr2 := executeCommand([]string{
					"__complete",
					"component", "create",
					"--org", config.Org,
					"--stack-ref", testStackRef,
					"--stack-tag", partial,
				})
				is.NoErr(cmdErr2)
				is.True(len(cmdOut2) > 0)                                        // Should have some output
				is.True(!bytes.Contains([]byte(cmdOut2), []byte("_activeHelp"))) // Should not have error help
			}
		}
	})

	t.Run("FailCompleteStackVersionTagNoStackRef", func(t *testing.T) {
		is := is.New(t)
		// Test completion without stack-ref should provide helpful error
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "create",
			"--org", config.Org,
			"--stack-tag", "",
		})
		// The command itself doesn't error, but completion should show active help
		is.NoErr(cmdErr)
		is.True(bytes.Contains([]byte(cmdOut), []byte("_activeHelp")))       // Should have error help
		is.True(bytes.Contains([]byte(cmdOut), []byte("missing stack-ref"))) // Should mention missing stack-ref
	})

	t.Run("SuccessCompleteStackVersionTagForUpdate", func(t *testing.T) {
		is := is.New(t)
		// For component update, if we provide stack-ref it should work
		cmdOut, cmdErr := executeCommand([]string{
			"__complete",
			"component", "update",
			"--org", config.Org,
			"--stack-ref", testStackRef,
			"--stack-tag", "",
		})
		is.NoErr(cmdErr)
		is.True(len(cmdOut) > 0)                                        // Should have some output
		is.True(!bytes.Contains([]byte(cmdOut), []byte("_activeHelp"))) // Should not have error help
		is.True(bytes.Contains([]byte(cmdOut), []byte(":4")))           // Should have NoFileComp directive
	})
}
