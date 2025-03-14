package table

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/printer"
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
value a	value b	abc	
       	       	def	
`
		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, b.String(), exp)
	})
	t.Run("SuccessTimestamp", func(t *testing.T) {
		now := time.Now()
		exptNow := now.Format(time.RFC3339)
		var (
			tab Table
			b   bytes.Buffer
			obj = struct {
				A *int64 `json:"a"`
			}{
				A: int64Ptr(now.Unix()), // int64Ptr(1578325983),
			}
		)

		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Contains(t, b.String(), exptNow)
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
	t.Run("SuccessSlicePtr", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
			obj = struct {
				A []*struct{}
			}{
				A: []*struct{}{
					&struct{}{},
					&struct{}{},
				},
			}
		)

		exp := `A 
2	
`

		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Equal(t, exp, b.String())
	})
}
