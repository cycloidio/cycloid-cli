package factory

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/json"
	"github.com/cycloidio/cycloid-cli/printer/table"
	"github.com/cycloidio/cycloid-cli/printer/yaml"
)

var (
	// printer is the lookup table to associate a printer
	// to a printer type
	// TODO: use constant + go enum to normalize the printer types
	printers = map[string]printer.Printer{
		"yaml":  yaml.YAML{},
		"yml":   yaml.YAML{},
		"json":  json.JSON{},
		"table": table.Table{},
	}
)

// GetPrinter is a helper to return the printer associated to the
// printer type passed in argument of the method
func GetPrinter(printer string) (printer.Printer, error) {
	p, ok := printers[printer]
	if !ok {
		return nil, fmt.Errorf("printer does not exist: %s", printer)
	}
	return p, nil
}
