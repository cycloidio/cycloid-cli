package middleware

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// encodeQuery encodes query parameters from either a url.Values or a struct
// with `url` tags into a url.Values.
func encodeQuery(query any) (url.Values, error) {
	if query == nil {
		return nil, nil
	}

	if v, ok := query.(url.Values); ok {
		return v, nil
	}

	result := url.Values{}
	rv := reflect.ValueOf(query)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return nil, nil
		}
		rv = rv.Elem()
	}

	if rv.Kind() != reflect.Struct {
		return nil, fmt.Errorf("encodeQuery: unsupported query type %T, must be url.Values or struct with url tags", query)
	}

	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		value := rv.Field(i)

		tag := field.Tag.Get("url")
		if tag == "" || tag == "-" {
			continue
		}

		name := strings.Split(tag, ",")[0]
		if name == "" {
			continue
		}

		if !encodeValue(result, name, value) {
			continue
		}
	}

	return result, nil
}

func encodeValue(result url.Values, name string, value reflect.Value) bool {
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return false
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.String:
		s := value.String()
		if s != "" {
			result.Set(name, s)
		}
	case reflect.Bool:
		result.Set(name, strconv.FormatBool(value.Bool()))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		result.Set(name, strconv.FormatInt(value.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result.Set(name, strconv.FormatUint(value.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		result.Set(name, strconv.FormatFloat(value.Float(), 'f', -1, 64))
	case reflect.Slice:
		for j := 0; j < value.Len(); j++ {
			elem := value.Index(j)
			if elem.Kind() == reflect.Ptr {
				if elem.IsNil() {
					continue
				}
				elem = elem.Elem()
			}
			if elem.Kind() == reflect.String {
				result.Add(name, elem.String())
			}
		}
	default:
		return false
	}

	return true
}
