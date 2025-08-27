package jsonformatter

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type JSONFormatter struct {
	// set the indentation of the JSON output
	indentSize int

	// escape the JSON output if true, will remove indentation
	escape bool

	// if true, will make a compact json on one line and ignore identation
	compact bool
}

// New will parse params and return a JSONFormatter
func New(params map[string][]string) *JSONFormatter {
	var f = &JSONFormatter{
		indentSize: 2,
		escape:     false,
		compact:    false,
	}

	for option, valueArray := range params {
		var value = ""
		if len(valueArray) > 0 {
			value = valueArray[0]
		}

		switch option {
		case "indent_size":
			indentSize, err := strconv.ParseInt(value, 10, 64)
			if err != nil || indentSize < 0 {
				f.indentSize = 2
				continue
			}
			f.indentSize = int(indentSize)
		case "json_escape":
			escape, err := strconv.ParseBool(value)
			if err != nil {
				// if the option is present, we escape
				f.escape = true
				continue
			}
			f.escape = escape
		case "json_compact":
			compact, err := strconv.ParseBool(value)
			if err != nil {
				// if the option is present, we compact
				f.compact = true
				continue
			}
			f.compact = compact
		}
	}

	return f
}

func (f *JSONFormatter) Format(data []any) (string, error) {
	var output any
	if len(data) == 1 {
		// if the final value is a string, print it directly
		if outStr, ok := data[0].(string); ok {
			return outStr, nil
		}

		output = data[0]
	} else {
		output = data
	}

	var jsonOutput []byte
	var err error
	if f.escape || f.compact {
		jsonOutput, err = json.Marshal(output)
	} else {
		jsonOutput, err = json.MarshalIndent(output, "", strings.Repeat(" ", f.indentSize))
	}
	if err != nil {
		return "", fmt.Errorf("failed to serialize output to JSON: %w", err)
	}

	if f.escape {
		return strconv.Quote(string(jsonOutput)), nil
	}

	return string(jsonOutput), nil
}
