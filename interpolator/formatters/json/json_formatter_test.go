package jsonformatter

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatterOptions(t *testing.T) {
	tcs := []struct {
		name   string
		params map[string][]string
		expect *JSONFormatter
	}{
		{
			"DefaultsOk",
			map[string][]string{},
			&JSONFormatter{indentSize: 2, escape: false, compact: false},
		},
		{
			"GoodIndentOk",
			map[string][]string{
				"output":      {"json"},
				"indent_size": {"4"},
			},
			&JSONFormatter{indentSize: 4, escape: false, compact: false},
		},
		{
			"BadIndentShouldDefault",
			map[string][]string{
				"output":      {"json"},
				"indent_size": {"-15"},
			},
			&JSONFormatter{indentSize: 2, escape: false, compact: false},
		},
		{
			"EscapeOnlyOk",
			map[string][]string{
				"output":      {"json"},
				"json_escape": {},
			},
			&JSONFormatter{indentSize: 2, escape: true, compact: false},
		},
		{
			"EscapeTrueOk",
			map[string][]string{
				"json_escape": {},
			},
			&JSONFormatter{indentSize: 2, escape: true, compact: false},
		},
		{
			"EscapeFalseOk",
			map[string][]string{
				"json_escape": {"false"},
			},
			&JSONFormatter{indentSize: 2, escape: false, compact: false},
		},
		{
			"CompactOnlyOk",
			map[string][]string{
				"json_compact": {},
			},
			&JSONFormatter{indentSize: 2, escape: false, compact: true},
		},
		{
			"CompactFalseOk",
			map[string][]string{
				"json_compact": {"false"},
			},
			&JSONFormatter{indentSize: 2, escape: false, compact: false},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := New(tc.params)
			assert.Equal(t, tc.expect, got, "Values must match, setup:", tc.params)
		})
	}
}

func TestJSONEscape(t *testing.T) {
	formatter := New(map[string][]string{"json_escape": {}})
	got, err := formatter.Format([]any{map[string]string{"hello": "world"}})
	assert.NoError(t, err, "formatting should not fail")
	assert.Equal(t, `"{\"hello\":\"world\"}"`, got, "output should be equal")

	// Unquote should work
	unquoted, err := strconv.Unquote(got)
	assert.NoError(t, err, "unquote should work")
	assert.Equal(t, `{"hello":"world"}`, unquoted)
}
