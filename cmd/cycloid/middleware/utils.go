package middleware

import "fmt"

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
