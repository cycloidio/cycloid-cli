package e2e_test

import (
	"encoding/json"
	"slices"
	"testing"

	"github.com/sanity-io/litter"
	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestAPIKeysCmd(t *testing.T) {
	var (
		testKeyCanonical = randomCanonical("test-admin")
		createdAPIKey    models.APIKey
	)
	t.Run("CreateAPIKeyAdminOk", func(t *testing.T) {
		args := []string{
			"--output", "json",
			"--org", config.Org,
			"api-key", "create",
			"--canonical", testKeyCanonical,
			"--description", "hello world",
			"--rules", `[{"action":"organization:**","effect":"allow","resources":[]}]`,
		}
		createOut, createErr := executeCommand(args)
		if createErr != nil {
			t.Errorf("failed to create api key admin: %s", createErr)
		}

		defer t.Run("DeleteAPIKeyOk", func(t *testing.T) {
			args := []string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "delete",
				"--canonical", testKeyCanonical,
			}
			_, deleteErr := executeCommand(args)
			if deleteErr != nil {
				t.Errorf("failed to delete api key admin: %s", deleteErr)
			}
		})

		err := json.Unmarshal([]byte(createOut), &createdAPIKey)
		if err != nil {
			t.Errorf("CLI output can't be serialize to models.APIKey, out:\n%s\nerr:\n%s", createOut, err)
		}

		assert.Equal(t, *createdAPIKey.Canonical, testKeyCanonical)

		var APIKeyList []*models.APIKey
		t.Run("ListAPIKeyOk", func(t *testing.T) {
			args := []string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "list",
			}
			listOut, listErr := executeCommand(args)
			if listErr != nil {
				t.Errorf("failed to list api keys: %s", listErr)
			}

			err := json.Unmarshal([]byte(listOut), &APIKeyList)
			if err != nil {
				t.Errorf("failed to parse cli output as list of api keys, out: %s\nerr: %s", listOut, err)
			}

			if index := slices.IndexFunc(APIKeyList, func(a *models.APIKey) bool {
				return *a.Canonical == testKeyCanonical
			}); index == -1 {
				t.Errorf("failed to find the API key with can '%s' in list: %s", testKeyCanonical, litter.Sdump(APIKeyList))
			}
		})

		t.Run("GetAPIKeyOk", func(t *testing.T) {
			args := []string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "get",
				"--canonical", testKeyCanonical,
			}
			getOut, getErr := executeCommand(args)
			if getErr != nil {
				t.Errorf("failed to get api key: %s", getErr)
			}

			var gotKey models.APIKey
			if err := json.Unmarshal([]byte(getOut), &gotKey); err != nil {
				t.Errorf("CLI output can't be serialized to models.APIKey, out:\n%s\nerr:\n%s", getOut, err)
			}
			assert.Equal(t, testKeyCanonical, *gotKey.Canonical)
		})

		t.Run("GetWithPositionalArg", func(t *testing.T) {
			args := []string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "get",
				testKeyCanonical, // positional, no --canonical
			}
			getOut, getErr := executeCommand(args)
			if getErr != nil {
				t.Errorf("failed to get api key via positional arg: %s", getErr)
			}

			var gotKey models.APIKey
			if err := json.Unmarshal([]byte(getOut), &gotKey); err != nil {
				t.Errorf("CLI output can't be serialized to models.APIKey, out:\n%s\nerr:\n%s", getOut, err)
			}
			assert.Equal(t, testKeyCanonical, *gotKey.Canonical)
		})

		t.Run("RecreateAPIKey", func(t *testing.T) {
			tmpCan := randomCanonical("test-recreate")
			rules := `[{"action":"organization:**","effect":"allow","resources":[]}]`

			// Create
			_, err := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "create",
				"--canonical", tmpCan,
				"--rules", rules,
			})
			if err != nil {
				t.Errorf("setup: failed to create api key for recreate test: %s", err)
			}

			defer t.Run("RecreateCleanup", func(t *testing.T) {
				executeCommand([]string{
					"--output", "json",
					"--org", config.Org,
					"api-key", "delete",
					tmpCan,
				})
			})

			// Second create without --recreate should fail
			_, dupErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "create",
				"--canonical", tmpCan,
				"--rules", rules,
			})
			assert.Error(t, dupErr, "second create without --recreate should fail")

			// Third create with --recreate should succeed
			_, recreateErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "create",
				"--canonical", tmpCan,
				"--rules", rules,
				"--recreate",
			})
			assert.NoError(t, recreateErr, "create --recreate should succeed for existing api key")
		})

		t.Run("DeleteWithPositionalArg", func(t *testing.T) {
			tmpCan := randomCanonical("test-del-pos")
			_, err := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "create",
				"--canonical", tmpCan,
				"--rules", `[{"action":"organization:**","effect":"allow","resources":[]}]`,
			})
			if err != nil {
				t.Errorf("setup: failed to create api key: %s", err)
			}

			_, deleteErr := executeCommand([]string{
				"--output", "json",
				"--org", config.Org,
				"api-key", "delete",
				tmpCan, // positional, no --canonical
			})
			assert.NoError(t, deleteErr, "delete via positional arg should succeed")
		})
	})
}
