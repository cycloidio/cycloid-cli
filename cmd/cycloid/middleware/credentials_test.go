package middleware_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/testcfg"
)

func TestCredentialsCRUD(t *testing.T) {
	m := config.Middleware

	canonical := testcfg.RandomCanonical("test-cred")
	name := "Test Credential"
	raw := &models.CredentialRaw{
		Password: "test-password",
		Username: "test-user",
	}

	created, _, err := m.CreateCredential(config.Org, name, "basic_auth", raw, "", canonical, "test credential")
	require.NoError(t, err, "CreateCredential should succeed")
	require.NotNil(t, created)

	defer func() {
		_, err := m.DeleteCredential(config.Org, *created.Canonical)
		require.NoError(t, err, "DeleteCredential should succeed")
	}()

	got, _, err := m.GetCredential(config.Org, *created.Canonical)
	require.NoError(t, err, "GetCredential should succeed")
	assert.Equal(t, *created.Canonical, *got.Canonical)

	list, _, err := m.ListCredentials(config.Org, "")
	require.NoError(t, err, "ListCredentials should succeed")
	assert.NotEmpty(t, list)

	updated, _, err := m.UpdateCredential(config.Org, name, "basic_auth", raw, "", canonical, "updated description")
	require.NoError(t, err, "UpdateCredential should succeed")
	assert.Equal(t, "updated description", updated.Description)
}
