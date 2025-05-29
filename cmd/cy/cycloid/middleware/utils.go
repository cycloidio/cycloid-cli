package middleware

import "fmt"

// from https://github.com/cycloidio/youdeploy-http-api/blob/develop/utils/convert.go

func convertInterfaceArray(ia []interface{}) []interface{} {
	res := make([]interface{}, len(ia))
	for i, v := range ia {
		res[i] = ConvertMapInterfaceToMapString(v)
	}
	return res
}

func convertInterfaceMap(mii map[interface{}]interface{}) map[string]interface{} {
	ms := make(map[string]interface{})
	for k, v := range mii {
		ms[fmt.Sprintf("%v", k)] = ConvertMapInterfaceToMapString(v)
	}
	return ms
}

// ConvertMapInterfaceToMapString is a method necessary to have JSON marshal
// working properly JSON doesn't know how to handle map[interface{}]interface{}
// so instead of keeping it, we are conveting it to map[string]interface{}.
// If the interface p isn't a map then we simply return it as it is.
func ConvertMapInterfaceToMapString(p interface{}) interface{} {
	switch v := p.(type) {
	case []interface{}:
		return convertInterfaceArray(v)
	case map[interface{}]interface{}:
		return convertInterfaceMap(v)
	case string:
		return v
	default:
		return v
	}
}
