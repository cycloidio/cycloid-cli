package printer

import (
	"fmt"
	"io"
)

// Printer is an interface that knows how to print runtime objects
type Printer interface {
	// Print will write the object in the writer using the given options
	Print(obj interface{}, opt Options, w io.Writer) error
}

// SmartPrint prints the object or print the error if not nil using a Printer
func SmartPrint(p Printer, obj interface{}, err error, errStr string, opt Options, w io.Writer) error {
	if err == nil && obj == nil {
		return nil
	}

	if err != nil {
		// print the result on the standard output
		if perr := p.Print(err, opt, w); perr != nil {
			return fmt.Errorf("unable to print result: %w", perr)
		}
		return fmt.Errorf("%s: %w", errStr, err)
	}

	// print the result on the standard output
	if perr := p.Print(obj, opt, w); perr != nil {
		return fmt.Errorf("unable to print result: %w", perr)
	}
	return nil
}
