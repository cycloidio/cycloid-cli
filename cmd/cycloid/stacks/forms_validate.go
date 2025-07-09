package stacks

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewFormsValidateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:               "validate",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.ValidateForms,
		Short:             "validate a .forms.yml file",
		Example:           `cy stacks validate-form --org my-org .forms.yml`,
		PreRunE:           internal.CheckAPIAndCLIVersion,
		RunE:              validateForm,
	}

	return cmd
}

func validateForm(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	formsPath := args[0]
	rawForm, err := os.ReadFile(formsPath)
	if err != nil {
		return errors.Wrapf(err, "unable to read the form file at path '%s'", formsPath)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	validation, err := m.ValidateForm(org, rawForm)
	if err != nil {
		return printer.SmartPrint(p, validation, err, "form validation failed", printer.Options{}, cmd.OutOrStdout())
	}

	if len(validation.Errors) == 0 {
		// do this to not output just `ERROR` with table output
		return printer.SmartPrint(p, nil, nil, "ok", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, validation, nil, "", printer.Options{}, cmd.OutOrStdout())
}
