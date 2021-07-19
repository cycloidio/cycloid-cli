package table

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/printer"
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
			values = append(values, value.String())
		case reflect.Uint32:
			values = append(values, strconv.FormatInt(int64(value.Uint()), 10))
		case reflect.Slice:
			// Render Slice of strings
			typ := value.Type().Elem()
			if typ.Kind() == reflect.String {
				stringSlice := make([]string, value.Len())
				for i := 0; i < value.Len(); i++ {
					stringSlice[i] = value.Index(i).String()
				}
				values = append(values, strings.Join(stringSlice[:], "\n"))
			} else {
				values = append(values, strconv.Itoa(value.Len()))
			}
		case reflect.Ptr:
			elt := value.Elem()
			switch elt.Kind() {
			case reflect.String:
				values = append(values, elt.String())
			case reflect.Uint32:
				values = append(values, strconv.FormatInt(int64(elt.Uint()), 10))
			case reflect.Int64:
				t := time.Unix(elt.Int(), 0)
				values = append(values, t.Format(time.RFC3339))
			default:
				// in the case we don't support the type, we print it
				// for further integration
				values = append(values, elt.Kind().String())
			}
		default:
			// in the case we don't support the type, we print it
			// for further integration
			values = append(values, value.Kind().String())
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
		// the following lines help to avoid
		// adding pointer to a struct into the table
		v := obj.FieldByName(f.Name)
		if v.Kind() == reflect.Ptr {
			elt := v.Elem()
			if elt.Kind() == reflect.Struct {
				continue
			}
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

	// Check if the obj is an API error with payload containing errors details
	apiErr, ok := obj.(interface {
		GetPayload() *models.ErrorPayload
	})
	if ok {
		errP := apiErr.GetPayload()
		if reflect.TypeOf(errP) == reflect.TypeOf(&models.ErrorPayload{}) {
			obj = errP.Errors
		}
	}

	// Print our obj
	headers, entries, err := generate(obj, opts)
	if err != nil {
		return err
	}
	table.SetHeader(headers)
	table.AppendBulk(entries)

	table.Render()
	return nil
}
