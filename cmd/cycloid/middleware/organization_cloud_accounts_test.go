package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/pkg/testcfg"
)

func TestCloudAccountCRUD(t *testing.T) {
	m := config.Middleware

	// Pre-create a credential — cloud accounts require one at create time.
	credCanonical := testcfg.RandomCanonical("test-ca-cred")
	cred, _, err := m.CreateCredential(config.Org, credCanonical, "aws",
		&models.CredentialRaw{AccessKey: "AKIAIOSFODNN7EXAMPLE", SecretKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"},
		"", credCanonical, "",
	)
	require.NoError(t, err, "CreateCredential should succeed")
	require.NotNil(t, cred)
	defer func() {
		// The backend cascades credential deletion when a cloud account is deleted,
		// so a 404 here is expected when the cloud account cleanup ran first.
		_, _ = m.DeleteCredential(config.Org, *cred.Canonical)
	}()

	// Create
	canonical := testcfg.RandomCanonical("test-ca")
	createBody := &models.NewCloudAccount{
		Canonical:           canonical,
		Name:                ptr.Ptr(canonical),
		CloudProvider:       ptr.Ptr("aws"),
		CredentialCanonical: ptr.Ptr(*cred.Canonical),
	}
	created, _, err := m.CreateCloudAccount(config.Org, createBody)
	require.NoError(t, err, "CreateCloudAccount should succeed")
	require.NotNil(t, created)
	assert.Equal(t, canonical, *created.Canonical)

	defer func() {
		_, err := m.DeleteCloudAccount(config.Org, *created.Canonical)
		require.NoError(t, err, "DeleteCloudAccount cleanup should succeed")
	}()

	// Get
	got, _, err := m.GetCloudAccount(config.Org, *created.Canonical)
	require.NoError(t, err, "GetCloudAccount should succeed")
	assert.Equal(t, *created.Canonical, *got.Canonical)
	assert.Equal(t, "aws", *got.CloudProvider)

	// List — created account must appear
	list, _, err := m.ListCloudAccounts(config.Org)
	require.NoError(t, err, "ListCloudAccounts should succeed")
	assert.NotEmpty(t, list)
	found := false
	for _, ca := range list {
		if ca.Canonical != nil && *ca.Canonical == *created.Canonical {
			found = true
			break
		}
	}
	assert.True(t, found, "created cloud account must appear in list")

	// Update — rename only; credential stays the same
	updatedName := canonical + "-updated"
	updateBody := &models.UpdateCloudAccount{
		Name:                ptr.Ptr(updatedName),
		CredentialCanonical: cred.Canonical,
	}
	updated, _, err := m.UpdateCloudAccount(config.Org, *created.Canonical, updateBody)
	require.NoError(t, err, "UpdateCloudAccount should succeed")
	assert.Equal(t, updatedName, *updated.Name)
}

func TestGetCloudAccount_NotFound(t *testing.T) {
	m := config.Middleware

	_, resp, err := m.GetCloudAccount(config.Org, "nonexistent-cloud-account-xyz")
	assert.Error(t, err, "GetCloudAccount should return an error for an unknown canonical")
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}
