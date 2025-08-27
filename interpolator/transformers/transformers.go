package transformers

import (
	"encoding/base64"
	"strconv"
	"strings"
)

type Transformer func(input string, options []string) (string, error)

var (
	AvailableTransformers = map[string]Transformer{
		"indent":       Indent,
		"nindent":      NIndent,
		"base64":       Base64,
		"base64encode": Base64,
	}
)

func Transform(input string, params map[string][]string) (string, error) {
	var out = input
	var err error
	for param, options := range params {
		if ts, ok := AvailableTransformers[param]; ok {
			out, err = ts(out, options)
			if err != nil {
				return out, err
			}
		}
	}

	return out, nil
}

func Indent(input string, options []string) (string, error) {
	var out string
	var pad = 0
	if len(options) >= 1 {
		pad, _ = strconv.Atoi(options[0])
	}

	for line := range strings.Lines(input) {
		out = out + strings.Repeat(" ", pad) + line
	}
	return out, nil
}

func NIndent(input string, options []string) (string, error) {
	out, err := Indent(input, options)
	if err != nil {
		return out, err
	}

	return "\n" + out, err
}

func Base64(input string, _ []string) (string, error) {
	enc := base64.StdEncoding
	return enc.EncodeToString([]byte(input)), nil
}
