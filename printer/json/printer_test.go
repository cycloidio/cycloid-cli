package json

import (
	"bytes"
	"testing"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
