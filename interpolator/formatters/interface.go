package formatters

import (
	jsonformatter "github.com/cycloidio/cycloid-cli/interpolator/formatters/json"
	yamlformatter "github.com/cycloidio/cycloid-cli/interpolator/formatters/yaml"
)

type AvailableFormat string

const (
	JSONFormat AvailableFormat = "json"
	YAMLFormat AvailableFormat = "yaml"
)

var (
	AvailableFormats = [...]AvailableFormat{
		JSONFormat,
		YAMLFormat,
	}
)

type Formatter interface {
	Format(data []any) (string, error)
}

func New(params map[string][]string) Formatter {
	outputFormat, ok := params["output"]
	if !ok || len(outputFormat) == 0 {
		return jsonformatter.New(params)
	}

	switch AvailableFormat(outputFormat[0]) {
	case YAMLFormat:
		return yamlformatter.New(params)
	default:
		return jsonformatter.New(params)
	}
}
