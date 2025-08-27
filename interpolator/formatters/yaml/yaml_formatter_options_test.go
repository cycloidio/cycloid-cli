package yamlformatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYAMLFormatterOptions(t *testing.T) {
	tcs := []struct {
		name   string
		params map[string][]string
		expect *YAMLFormatter
	}{
		{
			"DefaultsOk",
			map[string][]string{},
			&YAMLFormatter{indentSize: 2},
		},
		{
			"GoodIndentOk",
			map[string][]string{
				"output":      {"yaml"},
				"indent_size": {"4"},
			},
			&YAMLFormatter{indentSize: 4},
		},
		{
			"BadIndentShouldDefault",
			map[string][]string{
				"output":      {"yaml"},
				"indent_size": {"-15"},
			},
			&YAMLFormatter{indentSize: 2},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := New(tc.params)
			assert.Equal(t, tc.expect, got, "Values must match")
		})
	}
}
