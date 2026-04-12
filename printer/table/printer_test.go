package table

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/cycloidio/cycloid-cli/client/models"
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

		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "A")
		assert.Contains(t, out, "B")
		assert.Contains(t, out, "C")
		assert.Contains(t, out, "value a")
		assert.Contains(t, out, "value b")
		assert.Contains(t, out, "abc, def")
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
				A: int64Ptr(now.Unix()),
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
		// Nested struct field is skipped by headersFromStruct → no data columns.
		// Since there are no data columns, headers is empty → no output.
		assert.Empty(t, b.String())
	})
	t.Run("SuccessSlicePtr", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
			obj = struct {
				A []*struct{}
			}{
				A: []*struct{}{
					{},
					{},
				},
			}
		)

		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "A")
		assert.Contains(t, out, "2")
	})
}

func TestTablePrinter_StringSlice(t *testing.T) {
	t.Run("DefaultColumnName", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
		)
		err := tab.Print([]string{"foo", "bar"}, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "Value", "default column header should be 'Value'")
		assert.Contains(t, out, "foo")
		assert.Contains(t, out, "bar")
	})

	t.Run("CustomColumnName", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
		)
		err := tab.Print([]string{"proj-one", "proj-two"}, printer.Options{Columns: []string{"Canonical"}}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "Canonical", "column header should be 'Canonical'")
		assert.Contains(t, out, "proj-one")
		assert.Contains(t, out, "proj-two")
	})

	t.Run("EmptySlice", func(t *testing.T) {
		var (
			tab Table
			b   bytes.Buffer
		)
		err := tab.Print([]string{}, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Empty(t, b.String(), "empty slice should produce no output")
	})
}

// apiErrNilPayload mimics *middleware.APIResponseError with no parsed JSON error payload (e.g. HTML body).
type apiErrNilPayload struct {
	StatusCode int
}

func (*apiErrNilPayload) GetPayload() *models.ErrorPayload {
	return nil
}

func TestTablePrinter_APIErrorNilPayloadDoesNotPanic(t *testing.T) {
	var (
		tab Table
		b   bytes.Buffer
	)
	err := tab.Print(&apiErrNilPayload{StatusCode: 404}, printer.Options{}, &b)
	require.NoError(t, err)
	assert.NotEmpty(t, b.String())
}

func TestTablePrinter_BorderOption(t *testing.T) {
	t.Run("DefaultNoBorder", func(t *testing.T) {
		type row struct{ Name string }
		obj := []row{{"alice"}}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.NotContains(t, out, "╭", "default should not have rounded corner")
		assert.NotContains(t, out, "╮", "default should not have rounded corner")
	})

	t.Run("BorderMode", func(t *testing.T) {
		type row struct{ Name string }
		obj := []row{{"alice"}}
		var b bytes.Buffer
		tab := Table{opts: printer.TableOptions{Border: true}}
		err := tab.Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "╭", "border mode should have rounded top-left corner")
		assert.Contains(t, out, "╰", "border mode should have rounded bottom-left corner")
		assert.Contains(t, out, "│", "border mode should have vertical separators")
	})
}

func TestRenderValue(t *testing.T) {
	t.Run("Struct", func(t *testing.T) {
		type inner struct {
			A string
			B int
		}
		v := reflect.ValueOf(inner{"x", 1})
		result := renderValue(v)
		assert.Equal(t, "{record 2 fields}", result)
	})

	t.Run("PtrToStruct", func(t *testing.T) {
		type inner struct {
			A string
			B int
			C bool
		}
		val := inner{}
		v := reflect.ValueOf(&val)
		result := renderValue(v)
		assert.Equal(t, "{record 3 fields}", result)
	})

	t.Run("Int", func(t *testing.T) {
		v := reflect.ValueOf(int(42))
		assert.Equal(t, "42", renderValue(v))
	})

	t.Run("Int32", func(t *testing.T) {
		v := reflect.ValueOf(int32(7))
		assert.Equal(t, "7", renderValue(v))
	})

	t.Run("Float32", func(t *testing.T) {
		v := reflect.ValueOf(float32(3.14))
		result := renderValue(v)
		assert.Contains(t, result, "3.14")
	})

	t.Run("Float64", func(t *testing.T) {
		v := reflect.ValueOf(float64(2.718))
		result := renderValue(v)
		assert.Contains(t, result, "2.718")
	})
}

func TestEstimateWidth(t *testing.T) {
	headers := []string{"Name", "Status"}
	rows := [][]string{{"alice", "active"}}

	withoutBorder := estimateWidth(headers, rows, false)
	withBorder := estimateWidth(headers, rows, true)

	// Border adds numCols+1 = 3 chars
	assert.Equal(t, withoutBorder+3, withBorder, "border mode should add numCols+1 chars")
}

func TestExpandColumns(t *testing.T) {
	type full struct {
		Canonical string
		Name      string
		Extra1    string
		Extra2    string
	}
	obj := []full{{"c1", "n1", "e1", "e2"}}
	curated := []string{"Canonical", "Name"}

	expanded := expandColumns(obj, curated)

	assert.Equal(t, []string{"Canonical", "Name"}, expanded[:2], "curated cols should come first")
	assert.Len(t, expanded, 4, "should include all 4 fields")
	assert.Contains(t, expanded, "Extra1")
	assert.Contains(t, expanded, "Extra2")
}

func TestFitToWidth_ProtectedCount(t *testing.T) {
	headers := []string{"Canonical", "Name", "Extra1", "Extra2"}
	rows := [][]string{
		{"c1", "n1", strings.Repeat("x", 30), strings.Repeat("y", 30)},
	}

	// protectedCount=2 protects Canonical+Name; Extra1 and Extra2 can be dropped
	fitted, _ := fitToWidth(headers, rows, "", 40, 2, false, false)

	// Canonical and Name must survive
	assert.Contains(t, fitted, "Canonical")
	assert.Contains(t, fitted, "Name")
}
