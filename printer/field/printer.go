package field

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/cycloidio/cycloid-cli/printer"
)

// Field implements the printer interface by extracting a single attribute.
// One value per line. Supports dot notation for nested access (e.g. "owner.username").
type Field struct {
	name string
}

// New returns a Field printer for the given attribute name.
func New(name string) *Field {
	return &Field{name: name}
}

// Print marshals obj to JSON, extracts the named field, and writes one value per line.
func (f *Field) Print(obj interface{}, opt printer.Options, w io.Writer) error {
	raw, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("unable to marshal to JSON: %w", err)
	}
	var v interface{}
	if err := json.Unmarshal(raw, &v); err != nil {
		return fmt.Errorf("unable to unmarshal JSON: %w", err)
	}

	switch val := v.(type) {
	case []interface{}:
		for _, item := range val {
			if m, ok := item.(map[string]interface{}); ok {
				if fv, found := lookup(m, f.name); found {
					fmt.Fprintln(w, formatValue(fv))
				}
			}
		}
	case map[string]interface{}:
		if fv, found := lookup(val, f.name); found {
			fmt.Fprintln(w, formatValue(fv))
		}
	}
	return nil
}

// lookup finds a key in a map case-insensitively.
// Supports dot notation: "owner.username" walks nested maps.
func lookup(m map[string]interface{}, key string) (interface{}, bool) {
	parts := strings.SplitN(key, ".", 2)
	target := strings.ToLower(parts[0])

	for k, v := range m {
		if strings.ToLower(k) == target {
			if len(parts) == 1 {
				return v, true
			}
			// Recurse into nested map
			if nested, ok := v.(map[string]interface{}); ok {
				return lookup(nested, parts[1])
			}
			return nil, false
		}
	}
	return nil, false
}

// formatValue converts a JSON value to a display string (no JSON quoting for scalars).
func formatValue(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case float64:
		// JSON numbers are float64; display integers without decimal point
		if val == float64(int64(val)) {
			return fmt.Sprintf("%d", int64(val))
		}
		return fmt.Sprintf("%g", val)
	case bool:
		return fmt.Sprintf("%t", val)
	case nil:
		return ""
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}
