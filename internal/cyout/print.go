// Package cyout provides a one-liner helper to replace the 3-line
// GetOutput/GetPrinter/SmartPrint boilerplate used across all commands.
package cyout

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// Print outputs obj (or err) using the command's --output flag.
// Replaces the GetOutput/GetPrinter/SmartPrint 3-liner.
func Print(cmd *cobra.Command, obj interface{}, err error, errMsg string) error {
	return PrintWithOptions(cmd, obj, err, errMsg, printer.Options{})
}

// PrintWithOptions outputs with explicit column/identifier options for table mode.
// Commands with curated views call this to specify default columns and identifier.
//
// On error: writes a formatted error block to stderr (API errors include status,
// method, path, payload details, request ID; local errors show command context),
// then returns the wrapped error so main.go can print the final "Error:" summary line.
//
// On success: writes the object to stdout using the selected printer.
func PrintWithOptions(cmd *cobra.Command, obj interface{}, err error, errMsg string, opts printer.Options) error {
	if err != nil {
		printError(cmd, err)
		if errMsg != "" {
			return fmt.Errorf("%s: %w", errMsg, err)
		}
		return err
	}

	output, oerr := cyargs.GetOutput(cmd)
	if oerr != nil {
		return oerr
	}
	p, perr := factory.GetPrinter(output)
	if perr != nil {
		return perr
	}
	return printer.SmartPrint(p, obj, nil, "", opts, cmd.OutOrStdout())
}
