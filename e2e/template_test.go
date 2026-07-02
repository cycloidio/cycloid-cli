package e2e_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// report mirrors internal/templating.Report for decoding JSON output.
type report struct {
	Name     string   `json:"name"`
	Rendered string   `json:"rendered"`
	Unset    []string `json:"unset_vars"`
	Warnings []string `json:"warnings"`
	Error    string   `json:"error"`
}

// TestTemplateRender exercises `cy beta template render` end to end. The command
// is fully offline (no API calls), so it does not depend on backend fixtures.
func TestTemplateRender(t *testing.T) {
	dir := t.TempDir()
	tplPath := filepath.Join(dir, "main.tpl")
	require.NoError(t, os.WriteFile(tplPath,
		[]byte("project=($ .project $)\nenv=($ .env $)\nupper=($ .project | upper $)\n"), 0o600))

	t.Run("set flags and placeholder for unset known var", func(t *testing.T) {
		out, err := executeCommand([]string{"beta", "template", "render", "-f", tplPath, "--set", "project=my-app", "-o", "json"})
		require.NoError(t, err)
		var r report
		require.NoError(t, json.Unmarshal([]byte(out), &r))
		assert.Equal(t, "project=my-app\nenv=<placeholder:$env>\nupper=MY-APP\n", r.Rendered)
		assert.Equal(t, []string{"env"}, r.Unset)
	})

	t.Run("yaml context file with dotted set override", func(t *testing.T) {
		ctxPath := filepath.Join(dir, "ctx.yaml")
		require.NoError(t, os.WriteFile(ctxPath, []byte("project: from-file\nenv_vars:\n  region: us\n  zone: a\n"), 0o600))
		rtpl := filepath.Join(dir, "r.tpl")
		require.NoError(t, os.WriteFile(rtpl, []byte("r=($ .env_vars.region $) z=($ .env_vars.zone $)\n"), 0o600))

		out, err := executeCommand([]string{"beta", "template", "render", "-f", rtpl, "--context-file", ctxPath, "--set", "env_vars.region=eu", "-o", "json"})
		require.NoError(t, err)
		var r report
		require.NoError(t, json.Unmarshal([]byte(out), &r))
		// region overridden by --set; zone preserved from the file (deep merge).
		assert.Equal(t, "r=eu z=a\n", r.Rendered)
	})

	t.Run("stdin json context", func(t *testing.T) {
		out, _, err := executeCommandStdin(`{"project":"piped"}`,
			[]string{"beta", "template", "render", "-f", tplPath, "-o", "json"})
		require.NoError(t, err)
		var r report
		require.NoError(t, json.Unmarshal([]byte(out), &r))
		assert.Contains(t, r.Rendered, "project=piped")
	})

	t.Run("parse error sets error field and nonzero exit", func(t *testing.T) {
		bad := filepath.Join(dir, "bad.tpl")
		require.NoError(t, os.WriteFile(bad, []byte("($ range .items $)"), 0o600))
		out, err := executeCommand([]string{"beta", "template", "render", "-f", bad, "-o", "json"})
		assert.Error(t, err)
		var r report
		require.NoError(t, json.Unmarshal([]byte(out), &r))
		assert.NotEmpty(t, r.Error)
	})
}
