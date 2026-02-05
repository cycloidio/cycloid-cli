package middleware

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

// from https://github.com/cycloidio/youdeploy-http-api/blob/develop/utils/convert.go

func convertInterfaceArray(ia []any) []any {
	res := make([]any, len(ia))
	for i, v := range ia {
		res[i] = ConvertMapInterfaceToMapString(v)
	}
	return res
}

func convertInterfaceMap(mii map[any]any) map[string]any {
	ms := make(map[string]any)
	for k, v := range mii {
		ms[fmt.Sprintf("%v", k)] = ConvertMapInterfaceToMapString(v)
	}
	return ms
}

// ConvertMapInterfaceToMapString is a method necessary to have JSON marshal
// working properly JSON doesn't know how to handle map[any]any
// so instead of keeping it, we are conveting it to map[string]any.
// If the interface p isn't a map then we simply return it as it is.
func ConvertMapInterfaceToMapString(p any) any {
	switch v := p.(type) {
	case []any:
		return convertInterfaceArray(v)
	case map[any]any:
		return convertInterfaceMap(v)
	case string:
		return v
	default:
		return v
	}
}

var allowedCanonicalCharRegex = regexp.MustCompile("[^a-zA-Z0-9-_]")

// ToCanonical convert a name to a valid canonical
func ToCanonical(name string) string {
	replacer := strings.NewReplacer(
		" ", "_",
	)

	replaced := replacer.Replace(name)
	filtered := allowedCanonicalCharRegex.ReplaceAllString(replaced, "")
	trimmed := strings.Trim(filtered, "-_")
	return strings.ToLower(trimmed)
}

// NameOrCanonical will process name and canonical argument and return both
// if name is set, the canonical will be inferred from it
// if canonical is set, name will be a Capitalized version of the canonical
// if both are empty, you will get an error
func NameOrCanonical(name, canonical *string) (string, string, error) {
	if name == nil && canonical == nil {
		return "", "", fmt.Errorf("name or canonical is required, both are empty")
	}

	nameValue := ptr.Value(name)
	canonicalValue := ptr.Value(canonical)

	if canonicalValue == "" {
		return nameValue, ToCanonical(nameValue), nil
	}

	if nameValue == "" {
		return strings.ToUpper(canonicalValue[:1]) + canonicalValue[1:], canonicalValue, nil
	}

	return nameValue, canonicalValue, nil
}
