package cycloid_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/cmd"
)

// TestVersion_JSONWithBadAPIURL runs `cy version -o json` against a non-API host.
// The response body is HTML; errors are always shown in human-readable format on stderr
// regardless of --output flag, so the HTML content should appear in the error block.
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
	require.NotEmpty(t, rawErr, "expected error output on stderr (stdout=%q)", stdout.String())

	// Errors are always rendered as a human-readable block, not JSON
	lower := strings.ToLower(rawErr)
	require.True(t,
		strings.Contains(lower, "api error"),
		"expected 'API Error' header in error output, got: %q", truncateRunes(rawErr, 200))

	require.True(t,
		strings.Contains(lower, "<!doctype") || strings.Contains(lower, "<html"),
		"expected HTML snippet in error output body, got: %q", truncateRunes(rawErr, 200))
}

func truncateRunes(s string, max int) string {
	r := []rune(s)
	if len(r) <= max {
		return s
	}
	return string(r[:max]) + "…"
}
