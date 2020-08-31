package printer

import "io"

// Printer is an interface that knows how to print runtime objects
type Printer interface {
	// Print will write the object in the writer using the given options
	Print(obj interface{}, opt Options, w io.Writer) error
}
