package parsers

import (
	"fmt"
	"regexp"

	"github.com/cycloidio/cycloid-cli/interpolator/resolvers"
)

func ReplaceFile(resolver resolvers.ResourceResolver, file string) (string, error) {
	regex := `cy:\/\/[a-zA-Z0-9{}\/]+[[:graph:]]+\b`
	re, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("failed to compile regex '%s': %s", regex, err.Error())
	}

	return re.ReplaceAllStringFunc(file, func(uri string) string {
		out, err := resolver.Interpolate(uri)
		if err != nil {
			fmt.Println(err.Error())
			err = nil
			return uri
		}
		return out
	}), nil
}
