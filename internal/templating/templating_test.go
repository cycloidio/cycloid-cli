package templating

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	tcs := []struct {
		name        string
		tmpl        string
		ctx         Context
		wantOut     string
		wantUnset   []string
		wantWarnLen int
		wantErr     bool
	}{
		{
			name:    "provided variable",
			tmpl:    `p=($ .project $)`,
			ctx:     Context{"project": "my-app"},
			wantOut: `p=my-app`,
		},
		{
			name:      "unset known variable renders placeholder",
			tmpl:      `e=($ .env $)`,
			ctx:       Context{},
			wantOut:   `e=<placeholder:$env>`,
			wantUnset: []string{"env"},
		},
		{
			name:        "unknown variable warns and renders empty",
			tmpl:        `x=($ .nonsense $)`,
			ctx:         Context{},
			wantOut:     `x=<no value>`,
			wantWarnLen: 1,
		},
		{
			name:    "sprig function",
			tmpl:    `u=($ .project | upper $)`,
			ctx:     Context{"project": "app"},
			wantOut: `u=APP`,
		},
		{
			name:      "inventory_jwt unset is a placeholder, not an error",
			tmpl:      `j=($ .inventory_jwt $)`,
			ctx:       Context{},
			wantOut:   `j=<placeholder:$inventory_jwt>`,
			wantUnset: []string{"inventory_jwt"},
		},
		{
			name:    "parse error is captured in report",
			tmpl:    `($ range .items $)`,
			ctx:     Context{},
			wantErr: true,
		},
		{
			name:    "nested context value",
			tmpl:    `r=($ .env_vars.region $)`,
			ctx:     Context{"env_vars": map[string]any{"region": "eu-west-1"}},
			wantOut: `r=eu-west-1`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := Render(tc.name, tc.tmpl, tc.ctx)
			if tc.wantErr {
				assert.NotEmpty(t, r.Error)
				assert.Empty(t, r.Rendered)
				return
			}
			assert.Empty(t, r.Error)
			assert.Equal(t, tc.wantOut, r.Rendered)
			assert.Equal(t, tc.wantUnset, r.Unset)
			assert.Len(t, r.Warnings, tc.wantWarnLen)
		})
	}
}

// TestRenderDoesNotMutateContext guards that placeholder injection for one
// template never leaks into the caller's context (which is reused across files).
func TestRenderDoesNotMutateContext(t *testing.T) {
	ctx := Context{"project": "p"}
	_ = Render("t", `($ .env $)`, ctx)
	_, leaked := ctx["env"]
	assert.False(t, leaked, "placeholder for unset var must not leak into caller context")
}

func TestParseSet(t *testing.T) {
	t.Run("flat and dotted keys", func(t *testing.T) {
		got, err := ParseSet([]string{"project=app", "env_vars.region=eu", "env_vars.zone=a"})
		require.NoError(t, err)
		assert.Equal(t, Context{
			"project": "app",
			"env_vars": map[string]any{
				"region": "eu",
				"zone":   "a",
			},
		}, got)
	})

	t.Run("value may contain equals", func(t *testing.T) {
		got, err := ParseSet([]string{"token=a=b=c"})
		require.NoError(t, err)
		assert.Equal(t, Context{"token": "a=b=c"}, got)
	})

	t.Run("rejects missing equals", func(t *testing.T) {
		_, err := ParseSet([]string{"noeq"})
		assert.Error(t, err)
	})
}

func TestMergePrecedence(t *testing.T) {
	base := Context{"project": "base", "env_vars": map[string]any{"region": "us", "zone": "a"}}
	over := Context{"project": "over", "env_vars": map[string]any{"region": "eu"}}
	got := Merge(base, over)
	assert.Equal(t, "over", got["project"])
	// nested maps deep-merge: region overridden, zone preserved.
	assert.Equal(t, map[string]any{"region": "eu", "zone": "a"}, got["env_vars"])
}

func TestLoadContextFile(t *testing.T) {
	dir := t.TempDir()
	jsonPath := filepath.Join(dir, "ctx.json")
	yamlPath := filepath.Join(dir, "ctx.yaml")
	require.NoError(t, os.WriteFile(jsonPath, []byte(`{"project":"j","env_vars":{"region":"eu"}}`), 0o600))
	require.NoError(t, os.WriteFile(yamlPath, []byte("project: y\nenv_vars:\n  region: us\n"), 0o600))

	gotJSON, err := LoadContextFile(jsonPath)
	require.NoError(t, err)
	assert.Equal(t, "j", gotJSON["project"])

	gotYAML, err := LoadContextFile(yamlPath)
	require.NoError(t, err)
	assert.Equal(t, "y", gotYAML["project"])
	// YAML nested maps must normalize to map[string]any so templates can index
	// them AND so Merge can deep-merge them (see TestMergeNormalisedYAML).
	assert.Equal(t, map[string]any{"region": "us"}, gotYAML["env_vars"])
}

// TestMergeNormalisedYAML guards the layering bug where a YAML-loaded nested
// map (decoded as the named Context type) defeated Merge's map[string]any deep
// merge, so a partial --set override silently clobbered sibling keys.
func TestMergeNormalisedYAML(t *testing.T) {
	dir := t.TempDir()
	yamlPath := filepath.Join(dir, "ctx.yaml")
	require.NoError(t, os.WriteFile(yamlPath, []byte("env_vars:\n  region: us\n  zone: a\n"), 0o600))

	fileCtx, err := LoadContextFile(yamlPath)
	require.NoError(t, err)
	setCtx, err := ParseSet([]string{"env_vars.region=eu"})
	require.NoError(t, err)

	merged := Merge(fileCtx, setCtx)
	// region overridden by --set, zone preserved from the file.
	assert.Equal(t, map[string]any{"region": "eu", "zone": "a"}, merged["env_vars"])
}
