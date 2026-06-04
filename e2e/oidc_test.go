package e2e_test

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func TestOIDC(t *testing.T) {
	// --------------------------------------------------------------------
	// Mappings CRUD
	// --------------------------------------------------------------------

	// We need a team to map to. Reuse the fixture team created by TestTeams
	// if available, or create a disposable one here.
	teamCan := randomCanonical("oidc-team")

	t.Run("SetupTeam", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"team", "create",
			"--name", teamCan,
			"--role", "organization-member",
		})
		require.NoError(t, cmdErr, "setup: team create should succeed")
	})

	defer t.Run("TeardownTeam", func(t *testing.T) {
		executeCommand([]string{ //nolint:errcheck
			"--output", "json",
			"--org", config.Org,
			"team", "delete", teamCan,
		})
	})

	var mappingID uint32

	t.Run("SuccessOIDCMappingCreate", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "mappings", "create",
			"--group-name", "my-oidc-group",
			"--team", teamCan,
		})
		require.NoError(t, cmdErr, "mappings create should succeed")

		var m middleware.OIDCGroupMapping
		err := json.Unmarshal([]byte(cmdOut), &m)
		require.NoError(t, err, "output should be a valid OIDCGroupMapping JSON")
		assert.NotZero(t, m.ID, "created mapping should have a non-zero ID")
		assert.Equal(t, "my-oidc-group", m.GroupName, "group name should match")
		require.NotNil(t, m.Team, "team should be set")
		assert.Equal(t, teamCan, m.Team.Canonical, "team canonical should match")

		mappingID = m.ID
	})

	t.Run("SuccessOIDCMappingList", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "mappings", "list",
		})
		require.NoError(t, cmdErr, "mappings list should succeed")

		var mappings []middleware.OIDCGroupMapping
		err := json.Unmarshal([]byte(cmdOut), &mappings)
		require.NoError(t, err, "output should be a valid JSON array of OIDCGroupMapping")

		found := false
		for _, m := range mappings {
			if m.ID == mappingID {
				found = true
				assert.Equal(t, "my-oidc-group", m.GroupName)
				break
			}
		}
		assert.True(t, found, fmt.Sprintf("mapping ID %d should appear in list output", mappingID))
	})

	t.Run("SuccessOIDCMappingDelete", func(t *testing.T) {
		_, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "mappings", "delete",
			"--mapping-id", strconv.FormatUint(uint64(mappingID), 10),
		})
		require.NoError(t, cmdErr, "mappings delete should succeed")
	})

	// Verify deletion
	t.Run("SuccessOIDCMappingListAfterDelete", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "mappings", "list",
		})
		require.NoError(t, cmdErr, "mappings list after delete should succeed")

		var mappings []middleware.OIDCGroupMapping
		err := json.Unmarshal([]byte(cmdOut), &mappings)
		require.NoError(t, err, "output should parse as JSON array")

		for _, m := range mappings {
			assert.NotEqual(t, mappingID, m.ID, "deleted mapping should no longer appear in list")
		}
	})

	// --------------------------------------------------------------------
	// Settings get / set
	// --------------------------------------------------------------------

	t.Run("SuccessOIDCSettingsGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "settings", "get",
		})
		require.NoError(t, cmdErr, "settings get should succeed")

		var s middleware.OIDCOrganizationSettings
		err := json.Unmarshal([]byte(cmdOut), &s)
		require.NoError(t, err, "output should be a valid OIDCOrganizationSettings JSON")
	})

	t.Run("SuccessOIDCSettingsSet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "settings", "set",
			"--oidc-managed=false",
			"--no-match-policy", "keep_membership",
		})
		require.NoError(t, cmdErr, "settings set should succeed")

		var s middleware.OIDCOrganizationSettings
		err := json.Unmarshal([]byte(cmdOut), &s)
		require.NoError(t, err, "output should be a valid OIDCOrganizationSettings JSON")
		assert.Equal(t, "keep_membership", s.OIDCNoMatchPolicy)
		assert.False(t, s.OIDCManaged)
	})

	t.Run("SuccessOIDCSettingsGetAfterSet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "settings", "get",
		})
		require.NoError(t, cmdErr, "settings get after set should succeed")

		var s middleware.OIDCOrganizationSettings
		err := json.Unmarshal([]byte(cmdOut), &s)
		require.NoError(t, err, "output should parse as JSON")
		assert.Equal(t, "keep_membership", s.OIDCNoMatchPolicy, "no-match-policy should persist")
	})

	// --------------------------------------------------------------------
	// Integration get / set
	// --------------------------------------------------------------------

	t.Run("SuccessOIDCIntegrationSet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "integration", "set",
			"--issuer", "https://idp.example.com",
			"--client-id", "test-client",
			"--enabled",
		})
		require.NoError(t, cmdErr, "integration set should succeed")

		var i middleware.OIDCIntegration
		err := json.Unmarshal([]byte(cmdOut), &i)
		require.NoError(t, err, "output should be a valid OIDCIntegration JSON")
		assert.True(t, i.Enabled, "integration should be enabled")
		assert.Equal(t, "test-client", i.OidcClientID, "client ID should match")
		assert.Equal(t, "https://idp.example.com", i.OidcIssuer, "issuer should match")
	})

	t.Run("SuccessOIDCIntegrationGet", func(t *testing.T) {
		cmdOut, cmdErr := executeCommand([]string{
			"--output", "json",
			"--org", config.Org,
			"oidc", "integration", "get",
		})
		require.NoError(t, cmdErr, "integration get should succeed")

		var i middleware.OIDCIntegration
		err := json.Unmarshal([]byte(cmdOut), &i)
		require.NoError(t, err, "output should be a valid OIDCIntegration JSON")
		assert.True(t, i.Enabled, "integration should still be enabled after get")
		assert.Equal(t, "test-client", i.OidcClientID, "client ID should persist")
		assert.Equal(t, "https://idp.example.com", i.OidcIssuer, "issuer should persist")
	})
}
