package templating

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Context is the bag of template variables fed to the interpolation engine.
// Keys are the snake_case names used in templates (e.g. "project",
// "env_vars"). Values are decoded Go values (string, bool, float64, []any,
// map[string]any).
type Context map[string]any

// LoadContextFile reads a JSON or YAML file into a Context. Format is chosen by
// extension (.json → JSON, .yaml/.yml → YAML); any other extension is decoded
// as YAML, which is a superset of JSON.
func LoadContextFile(path string) (Context, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read context file %q: %w", path, err)
	}
	ctx := Context{}
	switch strings.ToLower(filepath.Ext(path)) {
	case ".json":
		if err := json.Unmarshal(raw, &ctx); err != nil {
			return nil, fmt.Errorf("failed to parse JSON context file %q: %w", path, err)
		}
	default:
		if err := yaml.Unmarshal(raw, &ctx); err != nil {
			return nil, fmt.Errorf("failed to parse YAML context file %q: %w", path, err)
		}
	}
	return normalizeContext(ctx), nil
}

// ParseContextString decodes a raw JSON object string into a Context.
func ParseContextString(s string) (Context, error) {
	if strings.TrimSpace(s) == "" {
		return Context{}, nil
	}
	ctx := Context{}
	if err := json.Unmarshal([]byte(s), &ctx); err != nil {
		return nil, fmt.Errorf("failed to parse --context JSON: %w", err)
	}
	return normalizeContext(ctx), nil
}

// normalizeContext rewrites every nested map to map[string]any regardless of
// how the decoder typed it (yaml.v3 reuses the named Context type for nested
// mappings; yaml.v2 can produce map[any]any). Without this, Merge's
// map[string]any type assertion would fail on nested maps and a partial --set
// override would clobber sibling keys instead of deep-merging.
func normalizeContext(ctx Context) Context {
	out := Context{}
	for k, v := range ctx {
		out[k] = normalizeValue(v)
	}
	return out
}

func normalizeValue(v any) any {
	switch m := v.(type) {
	case Context:
		return map[string]any(normalizeContext(m))
	case map[string]any:
		return map[string]any(normalizeContext(Context(m)))
	case map[any]any:
		nm := map[string]any{}
		for k, vv := range m {
			nm[fmt.Sprint(k)] = normalizeValue(vv)
		}
		return nm
	case []any:
		for i := range m {
			m[i] = normalizeValue(m[i])
		}
		return m
	default:
		return v
	}
}

// ParseSet turns repeatable `key=value` pairs into a Context. Dotted keys
// address nested maps: "env_vars.region=eu-west-1" becomes
// {"env_vars": {"region": "eu-west-1"}}. Values are kept as strings; wrap a
// whole-object override in --context or a context file when richer types are
// needed.
func ParseSet(pairs []string) (Context, error) {
	ctx := Context{}
	for _, p := range pairs {
		eq := strings.IndexByte(p, '=')
		if eq < 0 {
			return nil, fmt.Errorf("invalid --set %q: expected key=value", p)
		}
		key, val := p[:eq], p[eq+1:]
		if key == "" {
			return nil, fmt.Errorf("invalid --set %q: empty key", p)
		}
		setPath(ctx, strings.Split(key, "."), val)
	}
	return ctx, nil
}

// setPath assigns val at the dotted path within m, creating intermediate maps.
func setPath(m map[string]any, path []string, val any) {
	for i := 0; i < len(path)-1; i++ {
		next, ok := m[path[i]].(map[string]any)
		if !ok {
			next = map[string]any{}
			m[path[i]] = next
		}
		m = next
	}
	m[path[len(path)-1]] = val
}

// Merge deep-merges src into dst in place and returns dst. Nested maps are
// merged recursively; for any other type src overwrites dst. Call with sources
// in ascending precedence (later wins) — the §2.5 layering contract.
func Merge(dst, src Context) Context {
	if dst == nil {
		dst = Context{}
	}
	for k, sv := range src {
		if sm, ok := sv.(map[string]any); ok {
			if dm, ok := dst[k].(map[string]any); ok {
				Merge(dm, sm)
				continue
			}
		}
		dst[k] = sv
	}
	return dst
}
