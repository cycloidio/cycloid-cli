package formatters_test

import (
	"testing"

	"github.com/cycloidio/cycloid-cli/interpolator/formatters"
	"github.com/stretchr/testify/assert"
)

func TestFormatter(t *testing.T) {
	complexData := []any{map[string]map[string]map[string]map[string]any{
		"one": {
			"two": {
				"three": {
					"array": []string{"one", "two"},
					"map": map[string]string{
						"hello": "world",
					},
					"int":                  1,
					"string":               "hello_world",
					"string_with_spaces":   "hello world",
					"string_with_newlines": "---\none\ntwo\nthree\n",
				},
			},
		},
	}}

	tcs := []struct {
		name     string
		params   map[string][]string
		data     []any
		expect   string
		allowErr bool
	}{
		{
			"UseJSONFormatterOk",
			map[string][]string{
				"output":       {"json"},
				"json_compact": {},
			},
			[]any{map[string]string{"hello": "world"}},
			`{"hello":"world"}`,
			false,
		},
		{
			"UseJSONFormatterComplex",
			map[string][]string{
				"output":       {"json"},
				"json_compact": {},
			},
			complexData,
			`{"one":{"two":{"three":{"array":["one","two"],"int":1,"map":{"hello":"world"},"string":"hello_world","string_with_newlines":"---\none\ntwo\nthree\n","string_with_spaces":"hello world"}}}}`,
			false,
		},
		{
			"JSONIndentOk",
			map[string][]string{
				"output":      {"json"},
				"indent_size": {"2"},
			},
			[]any{map[string]string{"hello": "world"}},
			`{
  "hello": "world"
}`,
			false,
		},
		{
			"UseYAMLFormatterOk",
			map[string][]string{
				"output": {"yaml"},
			},
			[]any{map[string]string{"hello": "world"}},
			"hello: world\n",
			false,
		},
		{
			"UseYAMLFormatterComplex",
			map[string][]string{
				"output":      {"yaml"},
				"indent_size": {"2"},
			},
			complexData,
			`one:
  two:
    three:
      array:
      - one
      - two
      int: 1
      map:
        hello: world
      string: hello_world
      string_with_newlines: |
        ---
        one
        two
        three
      string_with_spaces: hello world
`,
			false,
		},
		{
			"YAMLIndentSettingsOk",
			map[string][]string{
				"output":      {"yaml"},
				"indent_size": {"6"},
			},
			[]any{map[string]map[string]string{"one": {"two": "hello_world"}}},
			"one:\n      two: hello_world\n",
			false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			formatter := formatters.New(tc.params)
			got, err := formatter.Format(tc.data)
			if err != nil && !tc.allowErr {
				assert.NoError(t, err, "formatting should not fail")
			}

			assert.Equal(t, tc.expect, got, "output strings should match")
		})
	}
}
