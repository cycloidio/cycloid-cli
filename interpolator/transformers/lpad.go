package transformers

import (
	"encoding/base64"
	"strconv"
	"strings"
)

type Transformer func(input string, options []string) string

var (
	AvailableTransformers = map[string]Transformer{
		"lpad":   LPad,
		"nlpad":  NLPad,
		"base64": Base64,
	}
)

func Transform(input string, params map[string][]string) string {
	var out = input
	for param, options := range params {
		if ts, ok := AvailableTransformers[param]; ok {
			out = ts(input, options)
		}
	}

	return out
}

func LPad(input string, options []string) string {
	var out string
	var pad = 0
	if len(options) >= 1 {
		pad, _ = strconv.Atoi(options[0])
	}
	for line := range strings.Lines(input) {
		out = out + strings.Repeat(" ", pad) + line
	}
	return out
}

func NLPad(input string, options []string) string {
	return "\n" + LPad(input, options)
}

func Base64(input string, _ []string) string {
	enc := base64.StdEncoding
	return enc.EncodeToString([]byte(input))
}
