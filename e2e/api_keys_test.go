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
	})
}
