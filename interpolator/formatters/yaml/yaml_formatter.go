package yamlformatter

import (
	"bytes"
	"fmt"
	"strconv"

	"go.yaml.in/yaml/v4"
)

type YAMLFormatter struct {
	// set the indentation of the YAML output
	// a value of 0 will not indentSize the YAML
	indentSize int
}

// New will parse params and return a YAMLFormatter
func New(params map[string][]string) *YAMLFormatter {
	var f = &YAMLFormatter{
		indentSize: 2,
	}

	for option, valueArray := range params {
		value := ""
		if len(valueArray) > 0 {
			value = valueArray[0]
		}

		switch option {
		case "indent_size":
			indentSize, err := strconv.Atoi(value)
			if err != nil || indentSize < 2 {
				f.indentSize = 2
			} else {
				f.indentSize = indentSize
			}
			f.indent = indent
		}
	}

	return f
}

func (f *YAMLFormatter) Format(data []any) (string, error) {
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

	writer := bytes.NewBuffer(nil)
	encoder := yaml.NewEncoder(writer)
	encoder.SetIndent(f.indentSize)
	encoder.CompactSeqIndent()
	defer encoder.Close()
	err := encoder.Encode(output)
	if err != nil {
		return "", fmt.Errorf("failed to serialize output to YAML: %q", err)
	}

	return writer.String(), nil
}
