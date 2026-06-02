package middleware

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSanitizeBody(t *testing.T) {
	t.Run("NilInput", func(t *testing.T) {
		assert.Nil(t, sanitizeBody(nil))
	})

	t.Run("EmptyInput", func(t *testing.T) {
		assert.Nil(t, sanitizeBody([]byte{}))
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		input := []byte("not json")
		assert.Equal(t, input, sanitizeBody(input))
	})

	t.Run("RedactsKnownKeys", func(t *testing.T) {
		input := []byte(`{"name":"my-cred","password":"secret123","type":"ssh"}`)
		out := sanitizeBody(input)
		require.NotNil(t, out)

		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		assert.Equal(t, "my-cred", m["name"])
		assert.Equal(t, "[REDACTED]", m["password"])
		assert.Equal(t, "ssh", m["type"])
	})

	t.Run("RedactsSSHKey", func(t *testing.T) {
		input := []byte(`{"ssh_key":"-----BEGIN RSA PRIVATE KEY-----"}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		assert.Equal(t, "[REDACTED]", m["ssh_key"])
	})

	t.Run("RawKeyIsRedacted", func(t *testing.T) {
		// "raw" is a sensitive catch-all field in CredentialRaw
		input := []byte(`{"name":"cred","raw":{"ssh_key":"secret","type":"ssh"}}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		assert.Equal(t, "[REDACTED]", m["raw"], "raw key itself should be redacted")
	})

	t.Run("NestedSensitiveField", func(t *testing.T) {
		input := []byte(`{"credential":{"password":"x","name":"foo"}}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		nested := m["credential"].(map[string]interface{})
		assert.Equal(t, "[REDACTED]", nested["password"])
		assert.Equal(t, "foo", nested["name"])
	})

	t.Run("ArrayElements", func(t *testing.T) {
		input := []byte(`[{"token":"abc"},{"token":"def","name":"foo"}]`)
		out := sanitizeBody(input)
		var arr []map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &arr))
		assert.Equal(t, "[REDACTED]", arr[0]["token"])
		assert.Equal(t, "[REDACTED]", arr[1]["token"])
		assert.Equal(t, "foo", arr[1]["name"])
	})

	t.Run("CaseInsensitiveKey", func(t *testing.T) {
		input := []byte(`{"Password":"xxx","SECRET_KEY":"yyy"}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		assert.Equal(t, "[REDACTED]", m["Password"])
		assert.Equal(t, "[REDACTED]", m["SECRET_KEY"])
	})

	t.Run("NoSensitiveKeys", func(t *testing.T) {
		input := []byte(`{"name":"foo","canonical":"bar"}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		assert.Equal(t, "foo", m["name"])
		assert.Equal(t, "bar", m["canonical"])
	})

	t.Run("RedactsOIDCSecrets", func(t *testing.T) {
		// `beta oidc integration set` sends oidc_client_secret / oidc_ca_cert in
		// the request body; these must never reach DEBUG or error output. Non-secret
		// OIDC fields must stay visible so debugging stays useful.
		input := []byte(`{"config":{"type":"AuthenticationOIDC","enabled":true,"oidc_issuer":"https://idp.example.com","oidc_client_id":"public-id","oidc_client_secret":"s3cr3t","oidc_ca_cert":"-----BEGIN CERTIFICATE-----"}}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		cfg := m["config"].(map[string]interface{})
		assert.Equal(t, "[REDACTED]", cfg["oidc_client_secret"], "oidc_client_secret must be redacted")
		assert.Equal(t, "[REDACTED]", cfg["oidc_ca_cert"], "oidc_ca_cert must be redacted")
		assert.Equal(t, "https://idp.example.com", cfg["oidc_issuer"], "non-secret oidc_issuer must stay visible")
		assert.Equal(t, "public-id", cfg["oidc_client_id"], "non-secret oidc_client_id must stay visible")
	})

	t.Run("RedactsSecretSuffixKeys", func(t *testing.T) {
		// Defense-in-depth: any *_secret / *_ca_cert field is redacted by suffix,
		// so a new provider's secret field does not leak by default.
		input := []byte(`{"saml_client_secret":"x","some_secret":"y","azure_ca_cert":"z","name":"ok"}`)
		out := sanitizeBody(input)
		var m map[string]interface{}
		require.NoError(t, json.Unmarshal(out, &m))
		assert.Equal(t, "[REDACTED]", m["saml_client_secret"])
		assert.Equal(t, "[REDACTED]", m["some_secret"])
		assert.Equal(t, "[REDACTED]", m["azure_ca_cert"])
		assert.Equal(t, "ok", m["name"])
	})
}
