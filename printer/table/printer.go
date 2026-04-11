package table

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/charmbracelet/lipgloss"
	libtable "github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/printer"
)

// Table is the default printer, rendering structured data as a terminal table.
type Table struct {
	opts printer.TableOptions
}

// NewWithOptions returns a Table printer configured with the parsed table options
// (from --output table[=cols][:flags]).
func NewWithOptions(opts printer.TableOptions) *Table {
	return &Table{opts: opts}
}

// Print renders obj as a terminal table to w.
func (t *Table) Print(obj interface{}, opts printer.Options, w io.Writer) error {
	// Unwrap API error payloads into their error list for display
	if apiErr, ok := obj.(interface {
		GetPayload() *models.ErrorPayload
	}); ok {
		if p := apiErr.GetPayload(); p != nil && reflect.TypeOf(p) == reflect.TypeOf(&models.ErrorPayload{}) && len(p.Errors) > 0 {
			obj = p.Errors
		}
	}

	headers, rows, err := build(obj, t.opts, opts)
	if err != nil {
		return err
	}
	if len(headers) == 0 {
		return nil
	}

	termWidth := terminalWidth()
	headers, rows = fitToWidth(headers, rows, opts.Identifier, termWidth, len(t.opts.Columns) > 0)

	tbl := libtable.New().
		Border(lipgloss.HiddenBorder()).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == libtable.HeaderRow {
				return lipgloss.NewStyle().Bold(true)
			}
			return lipgloss.NewStyle()
		})

	if !t.opts.NoHeader {
		displayHeaders := make([]string, len(headers))
		for i, h := range headers {
			displayHeaders[i] = camelToTitle(h)
		}
		tbl = tbl.Headers(displayHeaders...)
	}

	for _, row := range rows {
		tbl = tbl.Row(row...)
	}

	fmt.Fprint(w, tbl.Render())
	if !strings.HasSuffix(tbl.Render(), "\n") {
		fmt.Fprintln(w)
	}
	return nil
}

// build resolves which columns to use and renders each row.
func build(obj interface{}, tableOpts printer.TableOptions, printerOpts printer.Options) ([]string, [][]string, error) {
	// Column priority: user-specified > command defaults > reflection fallback
	var colNames []string
	switch {
	case len(tableOpts.Columns) > 0:
		colNames = tableOpts.Columns
	case len(printerOpts.Columns) > 0:
		colNames = printerOpts.Columns
	}

	rObj := reflect.ValueOf(obj)

	switch rObj.Kind() {
	case reflect.Ptr:
		elt := rObj.Elem()
		if printerOpts.Transform != nil {
			m := printerOpts.Transform(obj)
			headers, row := fromTransformMap(m, colNames)
			return headers, [][]string{row}, nil
		}
		headers := resolveHeaders(elt, colNames)
		return headers, [][]string{entryFromStruct(elt, headers)}, nil

	case reflect.Slice, reflect.Array:
		if rObj.Len() == 0 {
			return nil, nil, nil
		}
		// Derive headers from first element
		first := rObj.Index(0)
		if first.Kind() == reflect.Interface {
			first = first.Elem()
		}
		if first.Kind() == reflect.Ptr {
			first = first.Elem()
		}
		var headers []string
		if printerOpts.Transform != nil {
			// Transform the first item to discover column names
			m := printerOpts.Transform(rObj.Index(0).Interface())
			headers, _ = fromTransformMap(m, colNames)
		} else {
			headers = resolveHeaders(first, colNames)
		}

		rows := make([][]string, 0, rObj.Len())
		for i := 0; i < rObj.Len(); i++ {
			item := rObj.Index(i)
			var row []string
			if printerOpts.Transform != nil {
				m := printerOpts.Transform(item.Interface())
				_, row = fromTransformMap(m, headers)
			} else {
				elt := item
				if elt.Kind() == reflect.Interface {
					elt = elt.Elem()
				}
				if elt.Kind() == reflect.Ptr {
					elt = elt.Elem()
				}
				row = entryFromStruct(elt, headers)
			}
			rows = append(rows, row)
		}
		return headers, rows, nil

	default:
		return nil, nil, fmt.Errorf("unable to render table for type: %v", rObj.Kind())
	}
}

// resolveHeaders picks the display column names for a struct value.
// With explicit colNames: case-insensitive match + dot notation support.
// Without: all exported scalar fields (reflection fallback).
func resolveHeaders(v reflect.Value, colNames []string) []string {
	if len(colNames) == 0 {
		return headersFromStruct(v)
	}
	// Build a lowercase→actual-name map from the struct fields
	fieldMap := make(map[string]string)
	collectFields(v, "", fieldMap)

	headers := make([]string, 0, len(colNames))
	for _, name := range colNames {
		lower := strings.ToLower(name)
		if actual, ok := fieldMap[lower]; ok {
			headers = append(headers, actual)
		} else {
			// Accept the name as-is (for computed columns via Transform)
			headers = append(headers, name)
		}
	}
	return headers
}

// collectFields recursively populates a lowercase→dotPath map for a struct value.
func collectFields(v reflect.Value, prefix string, out map[string]string) {
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.Type().NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}
		path := f.Name
		if prefix != "" {
			path = prefix + "." + f.Name
		}
		lower := strings.ToLower(path)
		out[lower] = path
		// Also map top-level short name without prefix
		if prefix != "" {
			out[strings.ToLower(f.Name)] = path
		}
	}
}

// fromTransformMap extracts headers and a row from a Transform output map.
func fromTransformMap(m map[string]string, colNames []string) ([]string, []string) {
	if len(colNames) == 0 {
		// Use map keys in deterministic order (sorted by original insertion isn't possible
		// with maps, so fall back to the column definition order if available)
		// For display, collect all keys alphabetically
		headers := make([]string, 0, len(m))
		for k := range m {
			headers = append(headers, k)
		}
		row := make([]string, len(headers))
		for i, h := range headers {
			row[i] = m[h]
		}
		return headers, row
	}
	// Use provided colNames as headers
	lower := make(map[string]string, len(m))
	for k, v := range m {
		lower[strings.ToLower(k)] = v
	}
	row := make([]string, len(colNames))
	for i, col := range colNames {
		row[i] = lower[strings.ToLower(col)]
	}
	return colNames, row
}

// headersFromStruct extracts exported, displayable field names via reflection.
// Skips pointer-to-struct (nested objects) and unexported fields.
// Returns nil for non-struct values (strings, ints, etc.).
func headersFromStruct(v reflect.Value) []string {
	if v.Kind() != reflect.Struct {
		return nil
	}
	headers := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		if !f.IsExported() {
			continue
		}
		fv := v.Field(i)
		// Skip pointer-to-struct (nested objects)
		if fv.Kind() == reflect.Ptr {
			if elt := fv.Elem(); elt.IsValid() && elt.Kind() == reflect.Struct {
				continue
			}
		}
		// Skip direct struct fields
		if fv.Kind() == reflect.Struct {
			continue
		}
		headers = append(headers, f.Name)
	}
	return headers
}

// entryFromStruct renders a struct value to a string row for the given headers.
// Supports dot notation in header names (e.g., "Owner.Username").
func entryFromStruct(obj reflect.Value, headers []string) []string {
	row := make([]string, len(headers))
	for i, h := range headers {
		row[i] = fieldValueStr(obj, h)
	}
	return row
}

// fieldValueStr extracts a field value from a struct, supporting dot notation.
func fieldValueStr(v reflect.Value, path string) string {
	parts := strings.SplitN(path, ".", 2)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return ""
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ""
	}

	fv := v.FieldByName(parts[0])
	if !fv.IsValid() {
		return ""
	}
	if len(parts) == 1 {
		return renderValue(fv)
	}
	return fieldValueStr(fv, parts[1])
}

// renderValue converts a reflect.Value to a display string.
func renderValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.Slice:
		if v.Type().Elem().Kind() == reflect.String {
			parts := make([]string, v.Len())
			for i := 0; i < v.Len(); i++ {
				parts[i] = v.Index(i).String()
			}
			return strings.Join(parts, ", ")
		}
		return strconv.Itoa(v.Len())
	case reflect.Ptr:
		elt := v.Elem()
		if !elt.IsValid() {
			return ""
		}
		switch elt.Kind() {
		case reflect.String:
			return elt.String()
		case reflect.Uint32, reflect.Uint64:
			return strconv.FormatUint(elt.Uint(), 10)
		case reflect.Int64:
			t := time.Unix(elt.Int(), 0)
			return t.Format(time.RFC3339)
		case reflect.Bool:
			return fmt.Sprintf("%t", elt.Bool())
		default:
			return elt.Kind().String()
		}
	default:
		return v.Kind().String()
	}
}

// fitToWidth progressively drops rightmost columns when output exceeds terminal width.
// The identifier column is never dropped. If user explicitly chose columns, skip fitting.
func fitToWidth(headers []string, rows [][]string, identifier string, termWidth int, userChoseColumns bool) ([]string, [][]string) {
	if termWidth <= 0 || userChoseColumns || len(headers) == 0 {
		return headers, rows
	}

	identIdx := -1
	for i, h := range headers {
		if strings.EqualFold(h, identifier) {
			identIdx = i
			break
		}
	}

	for len(headers) > 1 {
		if estimateWidth(headers, rows) <= termWidth {
			break
		}
		// Drop the last column, but never the identifier
		last := len(headers) - 1
		if last == identIdx {
			break
		}
		headers = headers[:last]
		for i := range rows {
			if len(rows[i]) > last {
				rows[i] = rows[i][:last]
			}
		}
	}
	return headers, rows
}

// estimateWidth estimates the table's rendered width.
func estimateWidth(headers []string, rows [][]string) int {
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(camelToTitle(h))
	}
	for _, row := range rows {
		for i := range widths {
			if i < len(row) && len(row[i]) > widths[i] {
				widths[i] = len(row[i])
			}
		}
	}
	total := 0
	for _, w := range widths {
		total += w + 2 // +2 for padding
	}
	return total
}

// terminalWidth returns the current terminal width, or 0 if not a terminal.
func terminalWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 0
	}
	return w
}

// camelToTitle converts CamelCase to Title Case with spaces.
// e.g. "CreatedAt" → "Created At", "ConfigRepositoryCanonical" → "Config Repository Canonical"
var camelSplitter = regexp.MustCompile(`([a-z])([A-Z])|([A-Z]+)([A-Z][a-z])`)

func camelToTitle(s string) string {
	// Insert space between camel humps
	result := camelSplitter.ReplaceAllString(s, "$1$3 $2$4")
	// Capitalise first letter of each word
	words := strings.Fields(result)
	for i, w := range words {
		if len(w) > 0 {
			runes := []rune(w)
			runes[0] = unicode.ToUpper(runes[0])
			words[i] = string(runes)
		}
	}
	return strings.Join(words, " ")
}
