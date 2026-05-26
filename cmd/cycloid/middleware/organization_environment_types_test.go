package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/pkg/testcfg"
)

func TestEnvironmentTypeCRUD(t *testing.T) {
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-env-type")
	name := "Test Env Type"
	color := "#3498db"

	createBody := &models.NewEnvironmentType{
		Canonical: canonical,
		Name:      ptr.Ptr(name),
		Color:     ptr.Ptr(color),
	}
	created, _, err := m.CreateEnvironmentType(config.Org, createBody)
	require.NoError(t, err, "CreateEnvironmentType should succeed")
	require.NotNil(t, created)
	assert.Equal(t, canonical, *created.Canonical)
	assert.Equal(t, name, *created.Name)

	defer func() {
		_, err := m.DeleteEnvironmentType(config.Org, *created.Canonical)
		require.NoError(t, err, "DeleteEnvironmentType cleanup should succeed")
	}()

	// Get — must round-trip through list since GET-by-canonical is list-based
	got, _, err := m.GetEnvironmentType(config.Org, *created.Canonical)
	require.NoError(t, err, "GetEnvironmentType should succeed")
	assert.Equal(t, *created.Canonical, *got.Canonical)
	assert.Equal(t, color, *got.Color)

	// List — created type must appear
	list, _, err := m.ListEnvironmentTypes(config.Org)
	require.NoError(t, err, "ListEnvironmentTypes should succeed")
	assert.NotEmpty(t, list)
	found := false
	for _, et := range list {
		if et.Canonical != nil && *et.Canonical == *created.Canonical {
			found = true
			break
		}
	}
	assert.True(t, found, "created environment type must appear in list")

	// Update
	updatedName := name + " Updated"
	updateBody := &models.UpdateEnvironmentType{
		Name:  ptr.Ptr(updatedName),
		Color: ptr.Ptr(color),
	}
	updated, _, err := m.UpdateEnvironmentType(config.Org, *created.Canonical, updateBody)
	require.NoError(t, err, "UpdateEnvironmentType should succeed")
	assert.Equal(t, updatedName, *updated.Name)
}

func TestGetEnvironmentType_NotFound(t *testing.T) {
	m := config.Middleware

	_, _, err := m.GetEnvironmentType(config.Org, "nonexistent-env-type-xyz")
	assert.Error(t, err, "GetEnvironmentType should return an error for an unknown canonical")
	assert.Contains(t, err.Error(), "nonexistent-env-type-xyz")
}
