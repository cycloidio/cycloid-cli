package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

// TestLHSFilterDiscovery verifies which attributes the backend accepts as LHS
// filters on each List route that supports them. Field names follow the
// backend's filtersMappings convention: <table_alias>_<column>.
//
// Routes confirmed to NOT support LHS filters (backend handler does not call
// lhs.ParseQuery): ListTeamMembers, ListRoles, ListCredentials,
// ListCatalogRepositories, ListConfigRepositories, ListStackVersions,
// ListStackUseCases.
//
// A subtest failure means the attribute is not supported for that route.
// Results inform the godoc on each List* middleware method.
func TestLHSFilterDiscovery(t *testing.T) {
	m := config.Middleware
	org := config.Org

	t.Run("Projects", func(t *testing.T) {
		projA := testcfg.RandomCanonical("lhs-disc-a")
		projB := testcfg.RandomCanonical("lhs-disc-b")
		nameA := "LHS Discovery Alpha " + projA
		nameB := "LHS Discovery Beta " + projB
		configRepo := *config.ConfigRepo.Canonical

		_, _, err := m.CreateProject(org, nameA, projA, "desc a", configRepo, "", "", "default", "world")
		require.NoError(t, err, "setup: create projA")
		defer func() { _, _ = m.DeleteProject(org, projA, middleware.DeleteOptions{}) }()

		_, _, err = m.CreateProject(org, nameB, projB, "desc b", configRepo, "", "", "default", "world")
		require.NoError(t, err, "setup: create projB")
		defer func() { _, _ = m.DeleteProject(org, projB, middleware.DeleteOptions{}) }()

		t.Run("project_canonical[eq]", func(t *testing.T) {
			results, _, err := m.ListProjects(org, middleware.LHSFilter{Attribute: "project_canonical", Condition: "eq", Value: projA})
			require.NoError(t, err)
			require.NotEmpty(t, results)
			for _, p := range results {
				assert.Equal(t, projA, *p.Canonical)
			}
		})

		t.Run("project_name[eq]", func(t *testing.T) {
			results, _, err := m.ListProjects(org, middleware.LHSFilter{Attribute: "project_name", Condition: "eq", Value: nameA})
			require.NoError(t, err)
			require.NotEmpty(t, results)
			for _, p := range results {
				assert.Equal(t, nameA, *p.Name)
			}
		})

		t.Run("project_name[rlike]", func(t *testing.T) {
			results, _, err := m.ListProjects(org, middleware.LHSFilter{Attribute: "project_name", Condition: "rlike", Value: "LHS Discovery .*"})
			require.NoError(t, err)
			var found int
			for _, p := range results {
				if *p.Canonical == projA || *p.Canonical == projB {
					found++
				}
			}
			assert.Equal(t, 2, found, "rlike must match both lhs-disc projects")
		})
	})

	t.Run("Stacks", func(t *testing.T) {
		// Stack creation requires complex catalog-repo setup — rely on the
		// test stack seeded by testcfg.
		testCanonical := "stack-e2e-stackforms"
		testRef := org + ":" + testCanonical

		all, _, err := m.ListStacks(org)
		require.NoError(t, err)
		require.NotEmpty(t, all, "test stacks must exist")

		t.Run("service_catalog_ref[eq]", func(t *testing.T) {
			results, _, err := m.ListStacks(org, middleware.LHSFilter{Attribute: "service_catalog_ref", Condition: "eq", Value: testRef})
			require.NoError(t, err)
			require.NotEmpty(t, results)
			for _, s := range results {
				assert.Equal(t, testRef, *s.Ref)
			}
			assert.Less(t, len(results), len(all), "filtered must be smaller than full list")
		})

		t.Run("service_catalog_visibility[eq]", func(t *testing.T) {
			// valid values: local, shared, hidden (handled with custom logic in backend)
			results, _, err := m.ListStacks(org, middleware.LHSFilter{Attribute: "service_catalog_visibility", Condition: "eq", Value: "local"})
			require.NoError(t, err)
			for _, s := range results {
				assert.Equal(t, "local", *s.Visibility)
			}
		})

		t.Run("service_catalog_blueprint[eq]", func(t *testing.T) {
			results, _, err := m.ListStacks(org, middleware.LHSFilter{Attribute: "service_catalog_blueprint", Condition: "eq", Value: "false"})
			require.NoError(t, err)
			assert.NotNil(t, results)
		})
	})

	t.Run("Components", func(t *testing.T) {
		if config.Component == nil || config.Environment == nil {
			t.Skip("skipping: env/component fixtures unavailable")
		}
		knownCanonical := *config.Component.Canonical
		proj := *config.Project.Canonical
		env := *config.Environment.Canonical

		all, _, err := m.ListComponents(org, proj, env)
		require.NoError(t, err)
		require.NotEmpty(t, all, "test component must exist")

		t.Run("component_canonical[eq]", func(t *testing.T) {
			results, _, err := m.ListComponents(org, proj, env, middleware.LHSFilter{Attribute: "component_canonical", Condition: "eq", Value: knownCanonical})
			require.NoError(t, err)
			require.NotEmpty(t, results)
			for _, c := range results {
				assert.Equal(t, knownCanonical, *c.Canonical)
			}
		})

		t.Run("service_catalog_ref[eq]", func(t *testing.T) {
			stackRef := org + ":stack-e2e-stackforms"
			results, _, err := m.ListComponents(org, proj, env, middleware.LHSFilter{Attribute: "service_catalog_ref", Condition: "eq", Value: stackRef})
			require.NoError(t, err)
			assert.NotNil(t, results)
		})

		t.Run("environment_canonical[eq]", func(t *testing.T) {
			results, _, err := m.ListComponents(org, proj, env, middleware.LHSFilter{Attribute: "environment_canonical", Condition: "eq", Value: env})
			require.NoError(t, err)
			assert.NotEmpty(t, results)
		})
	})

	t.Run("Members", func(t *testing.T) {
		// List all members to get a known canonical for filtering
		all, _, err := m.ListMembers(org)
		require.NoError(t, err)
		require.NotEmpty(t, all, "org must have at least one member")
		knownCanonical := all[0].Username

		t.Run("user_canonical[eq]", func(t *testing.T) {
			results, _, err := m.ListMembers(org, middleware.LHSFilter{Attribute: "user_canonical", Condition: "eq", Value: knownCanonical})
			require.NoError(t, err)
			require.NotEmpty(t, results)
			for _, member := range results {
				assert.Equal(t, knownCanonical, member.Username)
			}
		})

		t.Run("invitation_state[eq]", func(t *testing.T) {
			// accepted is the normal state for active members
			results, _, err := m.ListMembers(org, middleware.LHSFilter{Attribute: "invitation_state", Condition: "eq", Value: "accepted"})
			require.NoError(t, err)
			assert.NotNil(t, results)
		})
	})

	t.Run("Teams", func(t *testing.T) {
		// List all teams; test env may or may not have teams, so we just verify
		// the filter is accepted without error. If teams exist, verify filtering.
		all, _, err := m.ListTeams(org, nil, nil, nil, nil)
		require.NoError(t, err)

		t.Run("team_canonical[eq]", func(t *testing.T) {
			if len(all) == 0 {
				t.Skip("no teams in test org")
			}
			knownCanonical := *all[0].Canonical
			results, _, err2 := m.ListTeams(org, nil, nil, nil, nil, middleware.LHSFilter{Attribute: "team_canonical", Condition: "eq", Value: knownCanonical})
			require.NoError(t, err2)
			for _, team := range results {
				assert.Equal(t, knownCanonical, *team.Canonical)
			}
		})

		t.Run("team_name[rlike]", func(t *testing.T) {
			if len(all) == 0 {
				t.Skip("no teams in test org")
			}
			_, _, err2 := m.ListTeams(org, nil, nil, nil, nil, middleware.LHSFilter{Attribute: "team_name", Condition: "rlike", Value: ".*"})
			require.NoError(t, err2)
		})
	})

	t.Run("InventoryResources", func(t *testing.T) {
		label := "lhs-disc-label"
		provider := "custom_resources"
		resType := "custom_instance"
		nameA := testcfg.RandomCanonical("lhs-res-a")
		nameB := testcfg.RandomCanonical("lhs-res-b")

		resA, _, err := m.CreateInventoryResource(org, &models.NewInventoryResource{
			Label:    &label,
			Name:     &nameA,
			Provider: &provider,
			Type:     &resType,
		})
		require.NoError(t, err, "setup: create inventory resource A")
		defer func() { _, _ = m.DeleteInventoryResource(org, resA.ID) }()

		resB, _, err := m.CreateInventoryResource(org, &models.NewInventoryResource{
			Label:    &label,
			Name:     &nameB,
			Provider: &provider,
			Type:     &resType,
		})
		require.NoError(t, err, "setup: create inventory resource B")
		defer func() { _, _ = m.DeleteInventoryResource(org, resB.ID) }()

		t.Run("resources_provider[eq]", func(t *testing.T) {
			results, _, err := m.ListInventoryResources(org, middleware.LHSFilter{Attribute: "resources_provider", Condition: "eq", Value: provider})
			require.NoError(t, err)
			for _, r := range results {
				assert.Equal(t, provider, *r.Provider)
			}
		})

		t.Run("resources_type[eq]", func(t *testing.T) {
			results, _, err := m.ListInventoryResources(org, middleware.LHSFilter{Attribute: "resources_type", Condition: "eq", Value: resType})
			require.NoError(t, err)
			for _, r := range results {
				assert.Equal(t, resType, *r.Type)
			}
		})

		t.Run("resources_label[eq]", func(t *testing.T) {
			results, _, err := m.ListInventoryResources(org, middleware.LHSFilter{Attribute: "resources_label", Condition: "eq", Value: label})
			require.NoError(t, err)
			require.GreaterOrEqual(t, len(results), 2)
			for _, r := range results {
				assert.Equal(t, label, r.Label)
			}
		})

		t.Run("resources_name[rlike]", func(t *testing.T) {
			results, _, err := m.ListInventoryResources(org, middleware.LHSFilter{Attribute: "resources_name", Condition: "rlike", Value: "lhs-res-.*"})
			require.NoError(t, err)
			var found int
			for _, r := range results {
				if *r.Name == nameA || *r.Name == nameB {
					found++
				}
			}
			assert.Equal(t, 2, found, "rlike must match both lhs-disc inventory resources")
		})
	})

	t.Run("InventoryOutputs", func(t *testing.T) {
		// Outputs come from TF states — cannot be created directly. We verify
		// the filter attributes are accepted by the route without error.
		t.Run("output_key[eq] accepted", func(t *testing.T) {
			_, _, err := m.ListInventoryOutputs(org, middleware.LHSFilter{Attribute: "output_key", Condition: "eq", Value: "nonexistent"})
			require.NoError(t, err)
		})

		t.Run("output_is_pinned[eq] accepted", func(t *testing.T) {
			_, _, err := m.ListInventoryOutputs(org, middleware.LHSFilter{Attribute: "output_is_pinned", Condition: "eq", Value: "false"})
			require.NoError(t, err)
		})

		t.Run("output_type[eq] accepted", func(t *testing.T) {
			_, _, err := m.ListInventoryOutputs(org, middleware.LHSFilter{Attribute: "output_type", Condition: "eq", Value: "string"})
			require.NoError(t, err)
		})

		t.Run("project_canonical[eq] accepted", func(t *testing.T) {
			_, _, err := m.ListInventoryOutputs(org, middleware.LHSFilter{Attribute: "project_canonical", Condition: "eq", Value: "nonexistent"})
			require.NoError(t, err)
		})
	})

	t.Run("APIKeys", func(t *testing.T) {
		all, _, err := m.ListAPIKeys(org)
		require.NoError(t, err)

		t.Run("organization_canonical[eq]", func(t *testing.T) {
			// organization_canonical is always org in this context
			results, _, err2 := m.ListAPIKeys(org, middleware.LHSFilter{Attribute: "organization_canonical", Condition: "eq", Value: org})
			require.NoError(t, err2)
			assert.Equal(t, len(all), len(results), "filtering by org canonical should return same count")
		})

		t.Run("user_canonical[eq]", func(t *testing.T) {
			if len(all) == 0 {
				t.Skip("no API keys in test org")
			}
			// Verify filter is accepted; nonexistent user returns empty not error
			_, _, err2 := m.ListAPIKeys(org, middleware.LHSFilter{Attribute: "user_canonical", Condition: "eq", Value: "nonexistent"})
			require.NoError(t, err2)
		})
	})
}
