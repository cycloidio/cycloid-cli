package mockresolver

import (
	"fmt"
	"strings"

	"github.com/cycloidio/cycloid-cli/interpolator/formatters"
	"github.com/cycloidio/cycloid-cli/interpolator/resources"
	"github.com/cycloidio/cycloid-cli/interpolator/transformers"
)

type MockResolver map[string][]any

func NewMockResolver(data map[string][]any) MockResolver {
	return data
}

type someObject struct {
	Str     string `yaml:"str"`
	Num     int    `yaml:"num"`
	Boolean bool   `yaml:"boolean"`
}

type bigObject struct {
	Obj     someObject   `yaml:"obj"`
	List    []string     `yaml:"list"`
	ListObj []someObject `yaml:"listObj"`
}

func NewMockResolverWithDefault() MockResolver {
	return map[string][]any{
		"/simple/string": {"simple"},
		"/simple/num":    {1},
		"/simple/bool":   {true},
		"/simple/object": {someObject{"simple", 1, true}},
		"/simple/bigObject": {
			bigObject{
				Obj:  someObject{"simple", 3, false},
				List: []string{"hello", "world"},
				ListObj: []someObject{
					{"one", 1, true},
					{"two", 2, false},
					{"three", 3, true},
				},
			},
		},
		"/list/string": {
			"one",
			"two",
			"three",
		},
		"/list/object": {
			someObject{"one", 1, true},
			someObject{"two", 2, false},
			someObject{"three", 3, true},
		},
		"/list/num": {
			1,
			2,
			3,
		},
		"/list/bool": {
			true,
			false,
			true,
		},
	}
}

func (r MockResolver) Resolve(ref *resources.Reference) ([]any, error) {
	value, ok := r[ref.Path]
	if !ok {
		return nil, fmt.Errorf("did not found %q in the mockresolver", ref.Path)
	}

	return value, nil
}

func (r MockResolver) Interpolate(uri string) (string, error) {
	path := strings.TrimPrefix(uri, "cy:/")
	data, ok := r[path]
	if !ok {
		return "", fmt.Errorf("invalid uri %q for mockresolver", uri)
	}

	formatter := formatters.New(map[string][]string{})
	dataStr, err := formatter.Format(data)
	if err != nil {
		return "", fmt.Errorf("formatter failed: %w", err)
	}

	out, err := transformers.Transform(dataStr, map[string][]string{})
	if err != nil {
		return "", fmt.Errorf("transformer failed: %w", err)
	}

	return out, nil
}
