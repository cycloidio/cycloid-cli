// Package templating renders Cycloid stack templates locally, offline, with no
// backend. It wraps the interpolation engine vendored from youdeploy-http-api
// (internal/templating/engine) and adds the CLI-facing behavior: layered
// context input (file / stdin / --set) and placeholder rendering for unset
// variables.
//
// Unset *known* variables render as the literal "<placeholder:$name>" instead
// of erroring or rendering empty, so a template can be exercised without a full
// platform context. Unknown bare references are surfaced as warnings (likely
// typos). This wrapper never mutates the engine, keeping it byte-identical to
// the backend for the render-parity test.
package templating

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cycloidio/cycloid-cli/internal/templating/engine"
)

// reToken matches a Cycloid interpolation token: ($ ... $).
var reToken = regexp.MustCompile(`\(\$([^)]+)\$\)`)

// reBareRef matches a single dotted field reference inside a token, e.g.
// ".project". Function calls (no leading dot) and compound expressions
// (pipes, multiple segments, arguments) deliberately do not match — those are
// left to the engine.
var reBareRef = regexp.MustCompile(`^\.([a-z_][a-z0-9_]*)$`)

// Report is the outcome of rendering one template.
type Report struct {
	Name     string   `json:"name"`
	Rendered string   `json:"rendered"`
	Unset    []string `json:"unset_vars,omitempty"` // known vars referenced but not provided → placeholder
	Warnings []string `json:"warnings,omitempty"`   // unknown references and other non-fatal notes
	Error    string   `json:"error,omitempty"`      // parse/render failure (Rendered is empty)
}

// PlaceholderFor returns the sentinel rendered for an unset known variable.
func PlaceholderFor(name string) string { return fmt.Sprintf("<placeholder:$%s>", name) }

// Render interpolates tmpl with ctx using the offline engine (latest, non-deprecated
// interpolation version). ctx is not mutated. name labels the template in
// errors and the report.
func Render(name, tmpl string, ctx Context) Report {
	report := Report{Name: name}

	// Work on a copy so placeholder injection never leaks into the caller's
	// shared context across multiple files.
	data := Merge(Context{}, ctx)

	known := engine.KnownKeys()
	for _, ref := range bareRefs(tmpl) {
		if _, provided := data[ref]; provided {
			continue
		}
		if _, ok := known[ref]; ok {
			data[ref] = PlaceholderFor(ref)
			report.Unset = append(report.Unset, ref)
		} else {
			report.Warnings = append(report.Warnings,
				fmt.Sprintf("unknown variable %q referenced; rendered empty", ref))
		}
	}

	interp := engine.Interpolator{Version: engine.Latest}
	out, err := interp.InterpolateWithExtraData(tmpl, name, map[string]any(data))
	if err != nil {
		report.Error = err.Error()
		return report
	}
	report.Rendered = out
	return report
}

// bareRefs returns the de-duplicated set of single-field references (without the
// leading dot) used in tmpl.
func bareRefs(tmpl string) []string {
	seen := map[string]struct{}{}
	var refs []string
	for _, m := range reToken.FindAllStringSubmatch(tmpl, -1) {
		inner := strings.TrimSpace(m[1])
		sub := reBareRef.FindStringSubmatch(inner)
		if sub == nil {
			continue
		}
		ref := sub[1]
		if _, ok := seen[ref]; ok {
			continue
		}
		seen[ref] = struct{}{}
		refs = append(refs, ref)
	}
	return refs
}
