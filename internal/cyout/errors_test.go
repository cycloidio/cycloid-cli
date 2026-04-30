package cyout

import (
	"bytes"
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// fakeAPIError implements apiErrorInfo + apiErrorPayloader for tests.
type fakeAPIError struct {
	statusCode int
	method     string
	path       string
	reqBody    []byte
	respBody   []byte
	payload    *models.ErrorPayload
}

func (e *fakeAPIError) Error() string                    { return "fake api error" }
func (e *fakeAPIError) HTTPStatusCode() int              { return e.statusCode }
func (e *fakeAPIError) HTTPRequestMethod() string        { return e.method }
func (e *fakeAPIError) HTTPRequestPath() string          { return e.path }
func (e *fakeAPIError) HTTPRequestBody() []byte          { return e.reqBody }
func (e *fakeAPIError) HTTPResponseBody() []byte         { return e.respBody }
func (e *fakeAPIError) GetPayload() *models.ErrorPayload { return e.payload }

func ptrString(s string) *string { return &s }

func newTestCmd() *cobra.Command {
	cmd := &cobra.Command{Use: "test"}
	var buf bytes.Buffer
	cmd.SetErr(&buf)
	return cmd
}

func captureError(cmd *cobra.Command, err error) string {
	var buf bytes.Buffer
	cmd.SetErr(&buf)
	printError(cmd, err)
	return buf.String()
}

func TestPrintError_LocalError(t *testing.T) {
	cmd := newTestCmd()
	out := captureError(cmd, errors.New("missing flag: --project"))
	// Local error: only command hint (no "API Error" line), main.go handles the Error: line
	assert.NotContains(t, out, "API Error")
}

func TestPrintError_APIErrorWithPayload(t *testing.T) {
	cmd := newTestCmd()
	err := &fakeAPIError{
		statusCode: 422,
		method:     "GET",
		path:       "/organizations",
		respBody:   []byte(`{"errors":[{"code":"Invalid","message":"This endpoint cannot be used with API Key"}]}`),
		payload: &models.ErrorPayload{
			Errors: []*models.ErrorDetailsItem{
				{Code: ptrString("Invalid"), Message: ptrString("This endpoint cannot be used with API Key")},
			},
		},
	}
	out := captureError(cmd, err)
	assert.Contains(t, out, "API Error 422")
	assert.Contains(t, out, "GET")
	assert.Contains(t, out, "/organizations")
	assert.Contains(t, out, "Invalid")
	assert.Contains(t, out, "This endpoint cannot be used with API Key")
}

func TestPrintError_APIErrorMultipleErrors(t *testing.T) {
	cmd := newTestCmd()
	err := &fakeAPIError{
		statusCode: 422,
		method:     "POST",
		path:       "/organizations/myorg/projects",
		payload: &models.ErrorPayload{
			Errors: []*models.ErrorDetailsItem{
				{Code: ptrString("RequiredField"), Message: ptrString("name: must not be empty")},
				{Code: ptrString("InvalidType"), Message: ptrString("color: must be a valid hex color")},
			},
		},
	}
	out := captureError(cmd, err)
	assert.Contains(t, out, "RequiredField")
	assert.Contains(t, out, "name: must not be empty")
	assert.Contains(t, out, "InvalidType")
	assert.Contains(t, out, "color: must be a valid hex color")
}

func TestPrintError_APIErrorWithDetails(t *testing.T) {
	cmd := newTestCmd()
	err := &fakeAPIError{
		statusCode: 422,
		method:     "POST",
		path:       "/organizations/myorg/projects",
		payload: &models.ErrorPayload{
			Errors: []*models.ErrorDetailsItem{
				{
					Code:    ptrString("InvalidValue"),
					Message: ptrString("invalid color"),
					Details: []string{"accepted values are #RRGGBB hex codes"},
				},
			},
		},
	}
	out := captureError(cmd, err)
	assert.Contains(t, out, "invalid color")
	assert.Contains(t, out, "accepted values are #RRGGBB hex codes")
}

func TestPrintError_APIErrorWithRequestBody(t *testing.T) {
	cmd := newTestCmd()
	err := &fakeAPIError{
		statusCode: 422,
		method:     "POST",
		path:       "/organizations/myorg/credentials",
		reqBody:    []byte(`{"name":"my-cred","ssh_key":"[REDACTED]"}`),
		payload: &models.ErrorPayload{
			Errors: []*models.ErrorDetailsItem{
				{Code: ptrString("InvalidField"), Message: ptrString("unsupported type")},
			},
		},
	}
	out := captureError(cmd, err)
	assert.Contains(t, out, "Body:")
	assert.Contains(t, out, "my-cred")
	assert.Contains(t, out, "[REDACTED]")
}

func TestPrintError_APIErrorRawBodyFallback(t *testing.T) {
	// No structured payload — falls back to raw response body
	cmd := newTestCmd()
	err := &fakeAPIError{
		statusCode: 503,
		method:     "GET",
		path:       "/health",
		respBody:   []byte("Service Unavailable"),
	}
	out := captureError(cmd, err)
	assert.Contains(t, out, "API Error 503")
	assert.Contains(t, out, "Service Unavailable")
}

func TestPrintError_APIErrorWithRequestID(t *testing.T) {
	cmd := newTestCmd()
	err := &fakeAPIError{
		statusCode: 500,
		method:     "GET",
		path:       "/organizations",
		payload: &models.ErrorPayload{
			RequestID: "550e8400-e29b-41d4-a716-446655440000",
			Errors: []*models.ErrorDetailsItem{
				{Code: ptrString("InternalError"), Message: ptrString("unexpected failure")},
			},
		},
	}
	out := captureError(cmd, err)
	assert.Contains(t, out, "Request-ID")
	assert.Contains(t, out, "550e8400-e29b-41d4-a716-446655440000")
}
