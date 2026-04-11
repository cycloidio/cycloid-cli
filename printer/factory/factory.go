package factory

import (
	"strings"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/field"
	"github.com/cycloidio/cycloid-cli/printer/jq"
	jsonprinter "github.com/cycloidio/cycloid-cli/printer/json"
	"github.com/cycloidio/cycloid-cli/printer/table"
	yamlprinter "github.com/cycloidio/cycloid-cli/printer/yaml"
)

var staticPrinters = map[string]printer.Printer{
	"yaml": yamlprinter.YAML{},
	"yml":  yamlprinter.YAML{},
	"json": jsonprinter.JSON{},
}

// GetPrinter returns a Printer for the given --output value.
//
// Dispatch rules:
//  1. jq=<expr>           → JQ printer (gojq expression)
//  2. table[=cols][:opts] → Table printer with options
//  3. json | yaml | yml   → static printers
//  4. <anything else>     → Field extractor (any attribute name, dot notation OK)
func GetPrinter(output string) (printer.Printer, error) {
	// JQ expression
	if strings.HasPrefix(output, "jq=") {
		return jq.New(strings.TrimPrefix(output, "jq="))
	}

	// Table with options (bare "table", "table=cols", or "table:opts")
	if output == "table" || strings.HasPrefix(output, "table=") || strings.HasPrefix(output, "table:") {
		opts := parseTableOptions(output)
		return table.NewWithOptions(opts), nil
	}

	// Known static printers
	if p, ok := staticPrinters[output]; ok {
		return p, nil
	}

	// Field extraction — any attribute name (e.g. "canonical", "id", "owner.username")
	return field.New(output), nil
}

// parseTableOptions parses the table output grammar into printer.TableOptions.
func parseTableOptions(output string) printer.TableOptions {
	var opts printer.TableOptions
	rest := strings.TrimPrefix(output, "table")
	if rest == "" {
		return opts
	}

	// Handle =cols shorthand: extract up to the first ':'
	if strings.HasPrefix(rest, "=") {
		rest = rest[1:]
		colonIdx := strings.Index(rest, ":")
		var colStr string
		if colonIdx >= 0 {
			colStr = rest[:colonIdx]
			rest = rest[colonIdx:]
		} else {
			colStr = rest
			rest = ""
		}
		if colStr != "" {
			opts.Columns = splitTrim(colStr, ",")
		}
	}

	// Parse :key=value and :flag options
	for _, part := range strings.Split(rest, ":") {
		part = strings.TrimSpace(part)
		switch {
		case part == "":
			// skip
		case part == "noheader":
			opts.NoHeader = true
		case strings.HasPrefix(part, "cols="):
			opts.Columns = splitTrim(strings.TrimPrefix(part, "cols="), ",")
		}
	}
	return opts
}

func splitTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	return out
}
