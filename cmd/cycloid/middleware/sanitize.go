package middleware

import (
	"encoding/json"
	"strings"
)

// sensitiveKeys is the set of JSON field names (lowercased) whose values are
// always redacted when displaying request bodies in error output.
// These cover credentials, tokens, and secrets across all Cycloid API models.
var sensitiveKeys = map[string]bool{
	"ssh_key":            true,
	"password":           true,
	"secret_key":         true,
	"access_key":         true,
	"client_secret":      true,
	"oidc_client_secret": true, // OIDC integration client secret
	"json_key":           true,
	"token":              true,
	"ca_cert":            true,
	"oidc_ca_cert":       true, // OIDC integration CA certificate (PEM)
	"raw":                true, // CredentialRaw.Raw — custom catch-all
	"current":            true, // password_update.current
}

// sensitiveKeySuffixes redacts any field whose (lowercased) name ends with one
// of these, so prefixed provider variants (oidc_client_secret, saml_client_secret,
// ...) are redacted by default rather than leaking into debug/error output.
var sensitiveKeySuffixes = []string{"_secret", "_ca_cert"}

// isSensitiveKey reports whether a JSON field name's value must be redacted.
func isSensitiveKey(k string) bool {
	lk := strings.ToLower(k)
	if sensitiveKeys[lk] {
		return true
	}
	for _, suf := range sensitiveKeySuffixes {
		if strings.HasSuffix(lk, suf) {
			return true
		}
	}
	return false
}

// sanitizeBody returns a compact JSON copy of body with sensitive field values
// replaced by "[REDACTED]". Returns nil if body is nil or empty. Returns body
// unchanged if it is not valid JSON (non-JSON bodies are typically not secrets).
func sanitizeBody(body []byte) []byte {
	if len(body) == 0 {
		return nil
	}
	var v interface{}
	if err := json.Unmarshal(body, &v); err != nil {
		return body
	}
	sanitized := sanitizeValue(v)
	out, err := json.Marshal(sanitized)
	if err != nil {
		return body
	}
	return out
}

func sanitizeValue(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		out := make(map[string]interface{}, len(val))
		for k, child := range val {
			// isSensitiveKey is the single redaction gate (exact-match map +
			// suffix net). Do not reintroduce a direct sensitiveKeys lookup here.
			if isSensitiveKey(k) {
				out[k] = "[REDACTED]"
			} else {
				out[k] = sanitizeValue(child)
			}
		}
		return out
	case []interface{}:
		out := make([]interface{}, len(val))
		for i, child := range val {
			out[i] = sanitizeValue(child)
		}
		return out
	default:
		return v
	}
}
