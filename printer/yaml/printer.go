package yaml

import (
	"io"

	"github.com/pkg/errors"

	"github.com/cycloidio/youdeploy-cli/printer"
	"gopkg.in/yaml.v1"
)

// YAML implements the printer interface
type YAML struct{}

// Print will write the object in the writer using the given options
func (y YAML) Print(obj interface{}, opts printer.Options, w io.Writer) error {
	yml, err := yaml.Marshal(obj)
	if err != nil {
		return errors.Wrap(err, "unable to marshal object")
	}
	if _, err = w.Write(yml); err != nil {
		return errors.Wrap(err, "unable to write YAML in the writer")
	}
	return nil
}
