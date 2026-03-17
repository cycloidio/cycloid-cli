package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/printer"
)

func TestJSONPrinter(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		var (
			j   JSON
			b   bytes.Buffer
			obj = struct {
				A string   `json:"a"`
				B string   `json:"b"`
				C []string `json:"c"`
			}{
				A: "value a",
				B: "value b",
				C: []string{"abc", "def"},
			}
		)
		exp := `{
  "a": "value a",
  "b": "value b",
  "c": [
    "abc",
    "def"
  ]
}
`

		err := j.Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, exp, b.String())

	})

}

// failMarshalHTTPError implements printer.ErrHTTPResponse and fails JSON marshaling.
type failMarshalHTTPError struct {
	code int
	body []byte
}

func (e *failMarshalHTTPError) Error() string { return "fail-marshal" }

func (e *failMarshalHTTPError) HTTPStatusCode() int { return e.code }

func (e *failMarshalHTTPError) HTTPResponseBody() []byte { return e.body }

func (e *failMarshalHTTPError) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("marshal intentionally failed")
}

type failMarshalHTTPErrorWithPath struct {
	failMarshalHTTPError
	path string
}

func (e *failMarshalHTTPErrorWithPath) HTTPRequestPath() string { return e.path }

func TestJSONPrinter_MarshalFallbackErrHTTPResponse(t *testing.T) {
	t.Run("DiagnosticJSON", func(t *testing.T) {
		body := []byte("a\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\nl")
		var j JSON
		var b bytes.Buffer
		err := j.Print(&failMarshalHTTPError{code: 422, body: body}, printer.Options{}, &b)
		require.NoError(t, err)

		var out map[string]any
		require.NoError(t, json.Unmarshal(b.Bytes(), &out))
		assert.EqualValues(t, 422, out["http_status"])
		assert.Contains(t, out["cli_marshal_error"].(string), "marshal intentionally failed")
		preview := out["api_response_preview"].(string)
		assert.Contains(t, preview, "a")
		assert.NotContains(t, preview, "k")
		_, hasPath := out["request_path"]
		assert.False(t, hasPath)
	})

	t.Run("IncludesRequestPath", func(t *testing.T) {
		var j JSON
		var b bytes.Buffer
		err := j.Print(&failMarshalHTTPErrorWithPath{
			failMarshalHTTPError: failMarshalHTTPError{code: 500, body: []byte("x")},
			path:                 "/organizations/o/projects/p",
		}, printer.Options{}, &b)
		require.NoError(t, err)
		var out map[string]any
		require.NoError(t, json.Unmarshal(b.Bytes(), &out))
		assert.Equal(t, "/organizations/o/projects/p", out["request_path"])
	})
}
