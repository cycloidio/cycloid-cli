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

// ---------------------------------------------------------------------------
// Nil / panic-safety tests
// ---------------------------------------------------------------------------

func TestTablePrinter_NilSafety(t *testing.T) {
	t.Run("NilInterface", func(t *testing.T) {
		var b bytes.Buffer
		err := (&Table{}).Print(nil, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Empty(t, b.String())
	})

	t.Run("TypedNilPointer", func(t *testing.T) {
		type item struct{ Name string }
		var p *item // typed nil
		var b bytes.Buffer
		err := (&Table{}).Print(p, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Empty(t, b.String())
	})

	t.Run("SliceOfAllNilPointers", func(t *testing.T) {
		type item struct{ Name string }
		objs := []*item{nil, nil, nil}
		var b bytes.Buffer
		err := (&Table{}).Print(objs, printer.Options{}, &b)
		require.NoError(t, err)
		// All nil → no headers discoverable → no output
		assert.Empty(t, b.String())
	})

	t.Run("SliceWithMixedNilPointers", func(t *testing.T) {
		type item struct{ Name string }
		objs := []*item{nil, {Name: "alice"}, nil}
		var b bytes.Buffer
		err := (&Table{}).Print(objs, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "Name")
		assert.Contains(t, out, "alice")
	})

	t.Run("SliceOfNilInterfaces", func(t *testing.T) {
		objs := []interface{}{nil, nil}
		var b bytes.Buffer
		err := (&Table{}).Print(objs, printer.Options{}, &b)
		require.NoError(t, err)
	})

	t.Run("SliceWithMixedNilInterfaces", func(t *testing.T) {
		type item struct{ Name string }
		objs := []interface{}{nil, &item{Name: "bob"}, nil}
		var b bytes.Buffer
		err := (&Table{}).Print(objs, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Contains(t, b.String(), "bob")
	})

	t.Run("EmptyStruct", func(t *testing.T) {
		obj := &struct{}{}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Empty(t, b.String())
	})

	t.Run("StructAllNilFields", func(t *testing.T) {
		type item struct {
			Name   *string
			Count  *int64
			Nested *struct{ X string }
		}
		obj := &item{}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		// Name and Count are ptr-to-scalar (shown), Nested is ptr-to-struct (hidden)
		out := b.String()
		assert.Contains(t, out, "Name")
		assert.Contains(t, out, "Count")
	})

	t.Run("NilSliceField", func(t *testing.T) {
		type item struct {
			Name string
			Tags []string
		}
		obj := &item{Name: "test", Tags: nil}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Contains(t, b.String(), "test")
	})

	t.Run("NilMapField", func(t *testing.T) {
		type item struct {
			Name   string
			Config map[string]string
		}
		obj := &item{Name: "test", Config: nil}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Contains(t, b.String(), "test")
	})

	t.Run("PopulatedMapField", func(t *testing.T) {
		type item struct {
			Name   string
			Config map[string]string
		}
		obj := &item{Name: "test", Config: map[string]string{"a": "1", "b": "2"}}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "test")
		assert.Contains(t, out, "{2 entries}")
	})

	t.Run("NonStructScalarValue", func(t *testing.T) {
		// Passing a bare string should return an error, not panic
		var b bytes.Buffer
		err := (&Table{}).Print("just a string", printer.Options{}, &b)
		assert.Error(t, err, "bare string should return error")
	})

	t.Run("NonStructIntValue", func(t *testing.T) {
		var b bytes.Buffer
		err := (&Table{}).Print(42, printer.Options{}, &b)
		assert.Error(t, err, "bare int should return error")
	})

	t.Run("DotNotationOnNilNestedField", func(t *testing.T) {
		type Owner struct{ Username string }
		type item struct {
			Name  string
			Owner *Owner
		}
		obj := &item{Name: "test", Owner: nil}
		var b bytes.Buffer
		tab := NewWithOptions(printer.TableOptions{Columns: []string{"Name", "Owner.Username"}})
		err := tab.Print(obj, printer.Options{}, &b)
		require.NoError(t, err)
		assert.Contains(t, b.String(), "test")
	})

	t.Run("DotNotationOnNilNestedInSlice", func(t *testing.T) {
		type Owner struct{ Username string }
		type item struct {
			Name  string
			Owner *Owner
		}
		objs := []*item{
			{Name: "alice", Owner: &Owner{Username: "alice_u"}},
			{Name: "bob", Owner: nil},
		}
		var b bytes.Buffer
		tab := NewWithOptions(printer.TableOptions{Columns: []string{"Name", "Owner.Username"}})
		err := tab.Print(objs, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "alice_u")
		assert.Contains(t, out, "bob")
	})

	t.Run("TransformReturnsNilMap", func(t *testing.T) {
		type item struct{ Name string }
		obj := &item{Name: "test"}
		var b bytes.Buffer
		err := (&Table{}).Print(obj, printer.Options{
			Transform: func(interface{}) map[string]string { return nil },
		}, &b)
		require.NoError(t, err)
		// nil map → empty headers → no output
		assert.Empty(t, b.String())
	})

	t.Run("ExpandColumnsNilSliceElement", func(t *testing.T) {
		type item struct{ Name string }
		objs := []*item{nil}
		result := expandColumns(objs, []string{"Name"})
		assert.Equal(t, []string{"Name"}, result, "should return curated cols when all nil")
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
