package stacks

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewFormsValidateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "validate [file...]",
		Args:              cobra.ArbitraryArgs,
		ValidArgsFunction: cyargs.ValidateForms,
		Short:             "validate one or more .forms.yml files",
		Example: `
	# validate a specific forms file
	cy stack forms validate .forms.yml

	# validate multiple files
	cy stack forms validate .forms.yml other/.forms.yaml

	# validate all .forms.y*ml files in the current directory (default)
	cy stack forms validate`,
		RunE: validateForm,
	}

	return cmd
}

func validateForm(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Default: glob .forms.y*ml in cwd
	if len(args) == 0 {
		matches, err := filepath.Glob(".forms.yml")
		if err == nil && len(matches) > 0 {
			args = append(args, matches...)
		}
		matches, err = filepath.Glob(".forms.yaml")
		if err == nil && len(matches) > 0 {
			args = append(args, matches...)
		}
		if len(args) == 0 {
			return fmt.Errorf("no .forms.yml or .forms.yaml found in current directory; pass a file path explicitly")
		}
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var hasErrors bool
	for _, formsPath := range args {
		rawForm, err := os.ReadFile(formsPath)
		if err != nil {
			return errors.Wrapf(err, "unable to read the form file at path '%s'", formsPath)
		}

		validation, _, err := m.ValidateForm(org, rawForm)
		if err != nil {
			return printer.SmartPrint(p, validation, err, fmt.Sprintf("form validation failed for %s", formsPath), printer.Options{}, cmd.OutOrStderr())
		}

		if len(validation.Errors) == 0 {
			if err := printer.SmartPrint(p, nil, nil, fmt.Sprintf("%s: ok", formsPath), printer.Options{}, cmd.OutOrStderr()); err != nil {
				return err
			}
			continue
		}

		hasErrors = true
		if err := printer.SmartPrint(p, validation, nil, "", printer.Options{}, cmd.OutOrStdout()); err != nil {
			return err
		}
	}

	if hasErrors {
		return fmt.Errorf("one or more forms files failed validation")
	}
	return nil
}
