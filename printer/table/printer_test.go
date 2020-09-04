package table

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cycloidio/youdeploy-cli/printer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYAMLPrinter(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		var (
			tab Table
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

		exp := `A      	B      	C 
value a	value b	
`
		err := tab.Print(&obj, printer.Options{}, &b)
		fmt.Println(b.String())
		require.NoError(t, err)
		assert.Equal(t, b.String(), exp)

	})
}
