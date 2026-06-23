package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/cycloidio/cycloid-cli/printer"
)

const apiResponsePreviewLineCount = 10

// JSON implements the printer interface
type JSON struct{}

// Print will write the object in the writer using the given options
func (j JSON) Print(obj interface{}, opts printer.Options, w io.Writer) error {
	payload, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		asErr, ok := obj.(error)
		if !ok {
			return fmt.Errorf("unable to marshal object: %w", err)
		}
		var httpErr printer.ErrHTTPResponse
		if errors.As(asErr, &httpErr) {
			diag := struct {
				CLIMarshalError    string `json:"cli_marshal_error"`
				HTTPStatus         int    `json:"http_status"`
				APIResponsePreview string `json:"api_response_preview"`
				RequestPath        string `json:"request_path,omitempty"`
			}{
				CLIMarshalError:    err.Error(),
				HTTPStatus:         httpErr.HTTPStatusCode(),
				APIResponsePreview: printer.FirstNLinesFromBytes(httpErr.HTTPResponseBody(), apiResponsePreviewLineCount),
			}
			if rp, ok := httpErr.(printer.RequestPather); ok {
				if p := rp.HTTPRequestPath(); p != "" {
					diag.RequestPath = p
				}
			}
			payload, err = json.MarshalIndent(diag, "", "  ")
			if err != nil {
				return fmt.Errorf("unable to marshal object: %w", err)
			}
			payload = append(payload, '\n')
			if _, werr := w.Write(payload); werr != nil {
				return fmt.Errorf("unable to write JSON in the writer: %w", werr)
			}
			return nil
		}
		return fmt.Errorf("unable to marshal object: %w", err)
	}

	// Add a newline to avoid ugly console output for user
	payload = append(payload, '\n')

	if _, err = w.Write(payload); err != nil {
		return fmt.Errorf("unable to write JSON in the writer: %w", err)
	}
	return nil
}
