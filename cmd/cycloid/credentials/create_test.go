package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func TestDefaultCredentialPath(t *testing.T) {
	t.Run("keep explicit path", func(t *testing.T) {
		assert.Equal(t, "explicit-path", defaultCredentialPath("explicit-path", "canonical", "name"))
	})

	t.Run("fallback to canonical", func(t *testing.T) {
		assert.Equal(t, "my-canonical", defaultCredentialPath("", "my-canonical", "name"))
	})

	t.Run("fallback to name when canonical missing", func(t *testing.T) {
		assert.Equal(t, "my-name", defaultCredentialPath("", "", "my-name"))
	})

	t.Run("CLI-104 name-only derives path like create", func(t *testing.T) {
		assert.Equal(t, "testflotmp", defaultCredentialPath("", "", "testflotmp"))
	})

	t.Run("empty when no identifier", func(t *testing.T) {
		assert.Equal(t, "", defaultCredentialPath("", "", ""))
	})
}

func TestFindCredentialForUpdate(t *testing.T) {
	credentials := []*models.CredentialSimple{
		{
			Canonical: ptr("cred-canonical"),
			Path:      ptr("cred-path"),
			Name:      ptr("cred-name"),
		},
		{
			Canonical: ptr("another-canonical"),
			Path:      ptr("another-path"),
			Name:      ptr("another-name"),
		},
	}

	t.Run("match by canonical first", func(t *testing.T) {
		got := findCredentialForUpdate(credentials, "cred-canonical", "another-path", "another-name")
		assert.NotNil(t, got)
		assert.Equal(t, "cred-canonical", *got.Canonical)
	})

	t.Run("match by path when canonical missing", func(t *testing.T) {
		got := findCredentialForUpdate(credentials, "", "cred-path", "")
		assert.NotNil(t, got)
		assert.Equal(t, "cred-canonical", *got.Canonical)
	})

	t.Run("match by name when canonical and path missing", func(t *testing.T) {
		got := findCredentialForUpdate(credentials, "", "", "cred-name")
		assert.NotNil(t, got)
		assert.Equal(t, "cred-canonical", *got.Canonical)
	})

	// CLI-104: `create --name X --update` leaves canonical empty; path defaults from name (e.g. testflotmp).
	t.Run("match by path from name defaulting (CLI-104)", func(t *testing.T) {
		issue := []*models.CredentialSimple{
			{
				Canonical: ptr("testflotmp"),
				Path:      ptr("testflotmp"),
				Name:      ptr("testflotmp"),
			},
		}
		got := findCredentialForUpdate(issue, "", "testflotmp", "testflotmp")
		assert.NotNil(t, got)
		assert.Equal(t, "testflotmp", *got.Canonical)
	})

	t.Run("no match", func(t *testing.T) {
		got := findCredentialForUpdate(credentials, "", "missing-path", "missing-name")
		assert.Nil(t, got)
	})
}

func ptr(value string) *string {
	return &value
}
