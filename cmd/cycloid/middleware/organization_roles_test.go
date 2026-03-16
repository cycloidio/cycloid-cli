package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

func TestRolesCRUD(t *testing.T) {
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-role")

	rule := &models.NewRule{Action: ptr.Ptr("*"), Effect: ptr.Ptr(models.NewRuleEffectAllow)}
	created, _, err := m.CreateRole(config.Org, ptr.Ptr("Test Role"), ptr.Ptr(canonical), ptr.Ptr("test role description"), []*models.NewRule{rule})
	require.NoError(t, err, "CreateRole should succeed")
	require.NotNil(t, created)

	defer func() {
		_, err := m.DeleteRole(config.Org, canonical)
		require.NoError(t, err, "DeleteRole should succeed")
	}()

	got, _, err := m.GetRole(config.Org, canonical)
	require.NoError(t, err, "GetRole should succeed")
	assert.Equal(t, canonical, *got.Canonical)

	list, _, err := m.ListRoles(config.Org)
	require.NoError(t, err, "ListRoles should succeed")
	assert.NotEmpty(t, list)
}
