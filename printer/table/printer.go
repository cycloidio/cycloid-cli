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
	// Guard: nil interface or typed nil pointer — nothing to render.
	if obj == nil {
		return nil
	}
	if v := reflect.ValueOf(obj); (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && v.IsNil() {
		return nil
	}

	// Unwrap API error payloads into their error list for display
	if apiErr, ok := obj.(interface {
		GetPayload() *models.ErrorPayload
	}); ok {
		if p := apiErr.GetPayload(); p != nil && reflect.TypeOf(p) == reflect.TypeOf(&models.ErrorPayload{}) && len(p.Errors) > 0 {
			obj = p.Errors
		}
	}

	termWidth := terminalWidth()

	// Dynamic column expansion: on wide terminals show all struct fields;
	// curated columns are protected from being dropped by fitToWidth.
	userChoseColumns := len(t.opts.Columns) > 0
	protectedCount := 0
	buildOpts := t.opts

	if len(opts.Columns) > 0 && termWidth > 0 && !userChoseColumns && opts.Transform == nil {
		expanded := expandColumns(obj, opts.Columns)
		if len(expanded) > len(opts.Columns) {
			buildOpts.Columns = expanded
			protectedCount = len(opts.Columns)
		}
	}

	headers, rows, err := build(obj, buildOpts, opts)
	if err != nil {
		return err
	}
	if len(headers) == 0 {
		return nil
	}

	headers, rows = fitToWidth(headers, rows, opts.Identifier, termWidth, protectedCount, t.opts.Border, userChoseColumns)

	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	var tbl *libtable.Table
	if t.opts.Border {
		tbl = libtable.New().
			Border(lipgloss.RoundedBorder()).
			BorderRow(false).
			BorderStyle(borderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				s := lipgloss.NewStyle().Padding(0, 1)
				if row == libtable.HeaderRow {
					return s.Bold(true)
				}
				return s
			})
	} else {
		tbl = libtable.New().
			Border(lipgloss.NormalBorder()).
			BorderTop(false).
			BorderBottom(false).
			BorderLeft(false).
			BorderRight(false).
			BorderColumn(false).
			BorderHeader(true).
			BorderStyle(borderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				s := lipgloss.NewStyle().Padding(0, 1)
				if row == libtable.HeaderRow {
					return s.Bold(true)
				}
				return s
			})
	}

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

	if termWidth > 0 {
		tbl = tbl.Width(termWidth)
	}

	output := tbl.Render()
	fmt.Fprint(w, output)
	if !strings.HasSuffix(output, "\n") {
		fmt.Fprintln(w)
	}
	return nil
}

// expandColumns returns curatedCols followed by any additional exported scalar
// fields from obj that are not already in curatedCols.
func expandColumns(obj interface{}, curatedCols []string) []string {
	v := reflect.ValueOf(obj)
	for (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) && !v.IsNil() {
		v = v.Elem()
	}
	if !v.IsValid() || v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return curatedCols
	}
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		if v.Len() == 0 {
			return curatedCols
		}
		first := v.Index(0)
		for (first.Kind() == reflect.Interface || first.Kind() == reflect.Ptr) && !first.IsNil() {
			first = first.Elem()
		}
		if !first.IsValid() || first.Kind() == reflect.Ptr || first.Kind() == reflect.Interface {
			return curatedCols
		}
		v = first
	}

	allFields := headersFromStruct(v)
	if len(allFields) == 0 {
		return curatedCols
	}

	curatedSet := make(map[string]bool, len(curatedCols))
	for _, c := range curatedCols {
		curatedSet[strings.ToLower(c)] = true
	}

	result := make([]string, len(curatedCols), len(curatedCols)+len(allFields))
	copy(result, curatedCols)
	for _, f := range allFields {
		if !curatedSet[strings.ToLower(f)] {
			result = append(result, f)
		}
	}
	return result
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
		if rObj.IsNil() {
			return nil, nil, nil
		}
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
		// Derive headers from first non-nil element
		first := firstNonNilElem(rObj)
		if !first.IsValid() {
			return nil, nil, nil
		}

		// String slices: render as a single-column table.
		// Column name comes from Columns[0] if set, otherwise "Value".
		if first.Kind() == reflect.String {
			colName := "Value"
			if len(colNames) > 0 {
				colName = colNames[0]
			}
			rows := make([][]string, rObj.Len())
			for i := 0; i < rObj.Len(); i++ {
				item := rObj.Index(i)
				if item.Kind() == reflect.Interface {
					item = item.Elem()
				}
				rows[i] = []string{item.String()}
			}
			return []string{colName}, rows, nil
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
				elt := derefValue(item)
				if !elt.IsValid() {
					// nil element — emit empty row
					row = make([]string, len(headers))
				} else {
					row = entryFromStruct(elt, headers)
				}
			}
			rows = append(rows, row)
		}
		return headers, rows, nil

	default:
		return nil, nil, fmt.Errorf("unable to render table for type: %v", rObj.Kind())
	}
}

// firstNonNilElem returns the first non-nil, dereferenced element from a slice/array.
// Returns an invalid reflect.Value if all elements are nil.
func firstNonNilElem(rObj reflect.Value) reflect.Value {
	for i := 0; i < rObj.Len(); i++ {
		v := derefValue(rObj.Index(i))
		if v.IsValid() {
			return v
		}
	}
	return reflect.Value{}
}

// derefValue unwraps interface and pointer wrappers, returning the underlying value.
// Returns an invalid reflect.Value if the chain terminates in nil.
func derefValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return reflect.Value{}
		}
		v = v.Elem()
	}
	return v
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
	if !v.IsValid() || v.Kind() != reflect.Struct {
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
	if !obj.IsValid() {
		return row
	}
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
	if !v.IsValid() {
		return ""
	}
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Int, reflect.Int32:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Bool())
	case reflect.Struct:
		n := countExportedFields(v.Type())
		return fmt.Sprintf("{record %d fields}", n)
	case reflect.Slice:
		if v.IsNil() {
			return ""
		}
		if v.Type().Elem().Kind() == reflect.String {
			parts := make([]string, v.Len())
			for i := 0; i < v.Len(); i++ {
				parts[i] = v.Index(i).String()
			}
			return strings.Join(parts, ", ")
		}
		return strconv.Itoa(v.Len())
	case reflect.Map:
		if v.IsNil() || v.Len() == 0 {
			return ""
		}
		return fmt.Sprintf("{%d entries}", v.Len())
	case reflect.Interface:
		if v.IsNil() {
			return ""
		}
		return renderValue(v.Elem())
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
		case reflect.Int, reflect.Int32:
			return strconv.FormatInt(elt.Int(), 10)
		case reflect.Float32, reflect.Float64:
			return strconv.FormatFloat(elt.Float(), 'f', -1, 64)
		case reflect.Bool:
			return fmt.Sprintf("%t", elt.Bool())
		case reflect.Struct:
			n := countExportedFields(elt.Type())
			return fmt.Sprintf("{record %d fields}", n)
		default:
			return elt.Kind().String()
		}
	default:
		return v.Kind().String()
	}
}

// countExportedFields returns the number of exported fields in a struct type.
func countExportedFields(t reflect.Type) int {
	n := 0
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).IsExported() {
			n++
		}
	}
	return n
}

// fitToWidth progressively drops rightmost columns when output exceeds terminal width.
// protectedCount > 0: protect the first protectedCount columns from being dropped.
// protectedCount == 0: protect the identifier column by name (legacy behavior).
// If user explicitly chose columns, skip fitting entirely.
func fitToWidth(headers []string, rows [][]string, identifier string, termWidth, protectedCount int, hasBorder, userChoseColumns bool) ([]string, [][]string) {
	if termWidth <= 0 || userChoseColumns || len(headers) == 0 {
		return headers, rows
	}

	// For legacy mode (no curated cols), find identifier index by name
	identIdx := -1
	if protectedCount == 0 {
		for i, h := range headers {
			if strings.EqualFold(h, identifier) {
				identIdx = i
				break
			}
		}
	}

	for len(headers) > 1 {
		if estimateWidth(headers, rows, hasBorder) <= termWidth {
			break
		}
		last := len(headers) - 1
		if protectedCount > 0 {
			if last < protectedCount {
				break
			}
		} else {
			if last == identIdx {
				break
			}
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
func estimateWidth(headers []string, rows [][]string, hasBorder bool) int {
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
		total += w + 2 // +2 for padding (1 left + 1 right)
	}
	if hasBorder {
		total += len(widths) + 1 // left border + column separators + right border
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
