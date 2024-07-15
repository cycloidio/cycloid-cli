package printer

import (
	"io"

	"github.com/pkg/errors"
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
			return errors.Wrap(perr, "unable to print result")
		}
		return errors.Wrap(err, errStr)
	}

	// print the result on the standard output
	if perr := p.Print(obj, opt, w); perr != nil {
		return errors.Wrap(perr, "unable to print result")
	}
	return nil
}
