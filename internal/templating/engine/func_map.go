package engine

import (
	"errors"
	"fmt"
	"strings"
	"text/template"
)

// recursionMaxNums defined the maximal count of nested references
const recursionMaxNums = 1000

// InitFuncMap inits a function map used by helm templating engine and sets it to the template.
// It is based on spring functions with helm extensions.
// If strict parameter is set to true, missing keys will be treated as errors.
// The content of this function is taken from
// https://github.com/helm/helm/blob/v3.19.0/pkg/engine/engine.go#L193
// The function reference can be found here:
// https://helm.sh/docs/howto/charts_tips_and_tricks/#know-your-template-functions
func FuncMap(t *template.Template, strict bool) template.FuncMap {
	funcMap := make(template.FuncMap, 3)

	// Add the template-rendering functions here so we can close over t.
	includedNames := make(map[string]int)
	funcMap["include"] = includeFun(t, includedNames)
	funcMap["tpl"] = tplFun(t, includedNames, strict)

	// Add the `required` function here so we can use lintMode
	// Reference: https://helm.sh/docs/howto/charts_tips_and_tricks/#using-the-required-function
	funcMap["required"] = func(warn string, val interface{}) (interface{}, error) {
		if val == nil {
			// if e.LintMode {
			// 	// Don't fail on missing required values when linting
			// 	slog.Warn("missing required value", "message", warn)
			// 	return "", nil
			// }
			return val, errors.New(warn)
		} else if _, ok := val.(string); ok {
			if val == "" {
				// if e.LintMode {
				// 	// Don't fail on missing required values when linting
				// 	slog.Warn("missing required values", "message", warn)
				// 	return "", nil
				// }
				return val, errors.New(warn)
			}
		}
		return val, nil
	}

	return funcMap
}

// 'include' needs to be defined in the scope of a 'tpl' template as
// well as regular file-loaded templates.
// Copied from https://github.com/helm/helm/blob/v3.19.0/pkg/engine/engine.go#L129
// Reference: https://helm.sh/docs/howto/charts_tips_and_tricks/#using-the-include-function
func includeFun(t *template.Template, includedNames map[string]int) func(string, interface{}) (string, error) {
	return func(name string, data interface{}) (string, error) {
		var buf strings.Builder
		if v, ok := includedNames[name]; ok {
			if v > recursionMaxNums {
				return "", fmt.Errorf("include recursion limit exceeded: included name: %s", name)
			}
			includedNames[name]++
		} else {
			includedNames[name] = 1
		}
		err := t.ExecuteTemplate(&buf, name, data)
		includedNames[name]--
		return buf.String(), err
	}
}

// As does 'tpl', so that nested calls to 'tpl' see the templates
// defined by their enclosing contexts.
// Copied from https://github.com/helm/helm/blob/v3.19.0/pkg/engine/engine.go#L148
// Reference: https://helm.sh/docs/howto/charts_tips_and_tricks/#using-the-tpl-function
func tplFun(parent *template.Template, includedNames map[string]int, strict bool) func(string, interface{}) (string, error) {
	return func(tpl string, vals interface{}) (string, error) {
		t, err := parent.Clone()
		if err != nil {
			return "", errors.New("cannot clone template")
		}

		// Re-inject the missingkey option, see text/template issue https://github.com/golang/go/issues/43022
		// We have to go by strict from our engine configuration, as the option fields are private in Template.
		// TODO: Remove workaround (and the strict parameter) once we build only with golang versions with a fix.
		if strict {
			t.Option("missingkey=error")
		} else {
			t.Option("missingkey=zero")
		}

		// Re-inject 'include' so that it can close over our clone of t;
		// this lets any 'define's inside tpl be 'include'd.
		t.Funcs(template.FuncMap{
			"include": includeFun(t, includedNames),
			"tpl":     tplFun(t, includedNames, strict),
		})

		// We need a .New template, as template text which is just blanks
		// or comments after parsing out defines just adds new named
		// template definitions without changing the main template.
		// https://pkg.go.dev/text/template#Template.Parse
		// Use the parent's name for lack of a better way to identify the tpl
		// text string. (Maybe we could use a hash appended to the name?)
		t, err = t.New(parent.Name()).Parse(tpl)
		if err != nil {
			return "", fmt.Errorf("cannot parse template %q", tpl)
		}

		var buf strings.Builder
		if err := t.Execute(&buf, vals); err != nil {
			return "", fmt.Errorf("error during tpl function execution for %q", tpl)
		}

		// See comment in renderWithReferences explaining the <no value> hack.
		return strings.ReplaceAll(buf.String(), "<no value>", ""), nil
	}
}
