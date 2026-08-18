package engine

import (
	"fmt"
	"regexp"
)

var functionNotDefinedRe = regexp.MustCompile("function \"(.*?)\" not defined")

// wrapInterpolatorErr produces a friendly message for template parse errors.
// Adapted from youdeploy-http-api/utils/interpolator_error.go: the backend wraps
// these in its yderr/errtmpl taxonomy; offline we only need the human-readable
// message, so it collapses to fmt.Errorf. Rendered output is unaffected.
func wrapInterpolatorErr(templateName string, err error) error {
	rm := functionNotDefinedRe.FindStringSubmatch(err.Error())
	if len(rm) != 0 {
		args := rm[1:]
		return fmt.Errorf(
			"interpolation error: %s function %q not defined, did you mean %q - a variable?",
			templateName,
			args[0],
			"."+args[0],
		)
	}

	return fmt.Errorf("interpolation error: %s", err.Error())
}
