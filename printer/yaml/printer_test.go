package yaml

import (
	"bytes"
	"testing"

	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYAMLPrinter(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		var (
			y   YAML
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

		exp := `a: value a
b: value b
c:
- abc
- def
`
		err := y.Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, b.String(), exp)

	})

}
