package json

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/printer"
)

// JSON implements the printer interface
type JSON struct{}

// Print will write the object in the writer using the given options
func (j JSON) Print(obj interface{}, opts printer.Options, w io.Writer) error {
	payload, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return errors.Wrap(err, "unable to marshal object")
	}

	// Add a newline to avoid ugly console output for user
	payload = append(payload, '\n')

	if _, err = w.Write(payload); err != nil {
		return errors.Wrap(err, "unable to write JSON in the writer")
	}
	return nil
}
