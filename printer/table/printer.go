package table

import (
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/cycloidio/youdeploy-cli/printer"

	"github.com/olekukonko/tablewriter"
)

type Table struct{}

// entryFromStruct is a helper to get struct field name
// which represents the column titles
func entryFromStruct(obj reflect.Value, h []string) []string {
	// we can't use obj.NumField() because it will
	// create a too big slice since it will include
	// the unexported field
	values := make([]string, 0)
	for _, header := range h {
		value := obj.FieldByName(header)
		switch value.Kind() {
		case reflect.String:
			values = append(values, value.Interface().(string))
		case reflect.Uint32:
			values = append(values, strconv.FormatUint(uint64(value.Interface().(uint32)), 10))
		case reflect.Ptr:
			elt := value.Elem()
			values = append(values, elt.String())
		}
	}
	return values
}

// headersFromStruct is a helper to get struct field name
// which represents the column titles
func headersFromStruct(obj reflect.Value, opts printer.Options) []string {
	// we don't set a size to the slice since `NumField`
	// return the number of struct fields including the
	// unexported fields
	headers := make([]string, 0)
	for i := 0; i < obj.NumField(); i++ {
		f := obj.Type().Field(i)
		// remove unexported fields
		if len(f.PkgPath) != 0 {
			continue
		}
		headers = append(headers, f.Name)
	}
	return headers
}

func generate(obj interface{}, opts printer.Options) ([]string, [][]string, error) {
	var (
		headers []string
		entries [][]string
		err     error
	)

	// obj can be a list of pointer or a pointer to struct
	// we need to handle both cases
	rObj := reflect.ValueOf(obj)

	switch rObj.Kind() {

	// the object is a pointer to a struct:
	// example: *models.ExternalBackend
	case reflect.Ptr:
		// we need to get the Value targetted by this pointer
		elt := rObj.Elem()
		headers = headersFromStruct(elt, opts)
		entries = make([][]string, 1)
		entry := entryFromStruct(elt, headers)
		entries = append(entries, entry)

	// the object is a slice of pointers to a struct
	// example: []*models.ExternalBackend
	case reflect.Slice, reflect.Array:
		if rObj.Len() > 0 {
			// it's supposed to be an uniform slice
			// create headers from the first element is enough
			elt := rObj.Index(0).Elem()
			headers = headersFromStruct(elt, opts)
		}
		entries = make([][]string, rObj.Len())
		for i := 0; i < rObj.Len(); i++ {
			elt := rObj.Index(i).Elem()
			entry := entryFromStruct(elt, headers)
			entries = append(entries, entry)
		}

	// default we return an error to help for further object types
	default:
		err = fmt.Errorf("unable to get headers for object type: %v", rObj.Kind())
	}

	return headers, entries, err
}

func (t Table) Print(obj interface{}, opts printer.Options, w io.Writer) error {

	// TODO: init the array using the opts
	// given by the user
	table := tablewriter.NewWriter(w)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)

	headers, entries, err := generate(obj, opts)
	if err != nil {
		return err
	}
	table.SetHeader(headers)
	table.AppendBulk(entries)

	table.Render()
	return nil
}
