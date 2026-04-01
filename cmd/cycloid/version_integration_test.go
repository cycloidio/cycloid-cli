package cycloid_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd"
)

// TestVersion_JSONWithBadAPIURL runs `cy version -o json` against a non-API host.
// The response body is HTML; JSON output must include HTTP context and a snippet of that HTML.
func TestVersion_JSONWithBadAPIURL(t *testing.T) {
	t.Setenv("CY_API_URL", "http://google.com")
	t.Setenv("CY_API_KEY", "")

	root := cmd.NewRootCommand()
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	root.SetArgs([]string{"version", "-o", "json"})

	err := root.Execute()
	require.Error(t, err)

	rawErr := strings.TrimSpace(stderr.String())
	require.NotEmpty(t, rawErr, "expected error JSON on stderr (stdout=%q)", stdout.String())

	var payload map[string]any
	require.NoError(t, json.Unmarshal([]byte(rawErr), &payload), "stderr should be JSON: %q", rawErr)

	htmlSample := htmlFromVersionErrorPayload(payload)
	require.NotEmpty(t, htmlSample, "expected HTML snippet in JSON payload: %#v", payload)

	lower := strings.ToLower(htmlSample)
	require.True(t,
		strings.Contains(lower, "<!doctype") || strings.Contains(lower, "<html"),
		"expected HTML document prefix in API body preview, got: %q", truncateRunes(htmlSample, 120))
}

func htmlFromVersionErrorPayload(payload map[string]any) string {
	// Non-2xx: *APIResponseError — Body is base64 in JSON.
	if bodyStr, ok := payload["Body"].(string); ok && bodyStr != "" {
		raw, err := base64.StdEncoding.DecodeString(bodyStr)
		if err == nil && len(raw) > 0 {
			return string(raw)
		}
	}
	// 2xx decode failure diagnostic (JSON printer fallback).
	if v, ok := payload["api_response_preview"].(string); ok {
		return v
	}
	return ""
}

func truncateRunes(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "…"
}
