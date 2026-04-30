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
}
