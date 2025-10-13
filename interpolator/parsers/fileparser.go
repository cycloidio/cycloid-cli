package parsers

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/cycloidio/cycloid-cli/interpolator/resolvers"
)

var re = regexp.MustCompile(`cy:\/\/[a-zA-Z0-9{}\/]+[[:graph:]]+\b`)

func ReplaceFile(resolver resolvers.ResourceResolver, file string) (string, error) {
	var errList []error
	output := re.ReplaceAllStringFunc(file, func(uri string) string {
		out, err := resolver.Interpolate(uri)
		if err != nil {
			errList = append(errList, err)
			return uri
		}
		return out
	})
	err := errors.Join(errList...)
	if err != nil {
		return file, fmt.Errorf("failed to interpolate the whole file: %w", err)
	}

	if output == "" && file != "" {
		// If the output is empty wile the file has content,
		// something has got wrong, so send the original instead
		return file, fmt.Errorf("failed to interpolate file, output is empty: %w", err)
	}

	cleanedOut := strings.TrimRight(output, " \t\n")
	return cleanedOut, nil
}
