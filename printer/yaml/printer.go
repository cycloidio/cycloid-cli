package yaml

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/cycloidio/cycloid-cli/printer"
)

// YAML implements the printer interface
type YAML struct{}

// Print will write the object in the writer using the given options
func (y YAML) Print(obj interface{}, opts printer.Options, w io.Writer) error {
	yml, err := yaml.Marshal(obj)
	if err != nil {
		return fmt.Errorf("unable to marshal object: %w", err)
	}
	if _, err = w.Write(yml); err != nil {
		return fmt.Errorf("unable to write YAML in the writer: %w", err)
	}
	return nil
}
