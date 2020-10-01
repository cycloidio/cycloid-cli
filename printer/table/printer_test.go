package table

import (
	"bytes"
	"testing"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func int64Ptr(i int64) *int64 {
	return &i
}

func TestTablePrinter(t *testing.T) {

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
value a	value b	slice	
`
		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, b.String(), exp)
	})
	t.Run("SuccessTimestamp", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
			obj = struct {
				A *int64 `json:"a"`
			}{
				A: int64Ptr(1578325983),
			}
		)

		exp := `A                    
06/01/2020, 16:53:03	
`
		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, b.String(), exp)
	})
	t.Run("SuccessAvoidNestedStruct", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
			obj = struct {
				A *struct{}
			}{
				A: &struct{}{},
			}
		)

		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, len(b.String()), 0)
	})
}
