package table

import (
	"bytes"
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

		err := tab.Print(&obj, printer.Options{}, &b)
		require.NoError(t, err)
		out := b.String()
		assert.Contains(t, out, "A")
		assert.Contains(t, out, "2")
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
