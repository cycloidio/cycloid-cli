package stacks

import (
	"io/ioutil"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewValidateFormCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "validate-form",
		Short: "validate a .forms.yml file",
		Example: `
		# validate a stackforms file
		cy stacks validate-form --org my-org --forms .forms.yml
		`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    validateForm,
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)
	common.RequiredPersistentFlag(WithFlagForms, cmd)

	return cmd
}

func validateForm(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	formPath, err := cmd.Flags().GetString("forms")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	rawForms, err := ioutil.ReadFile(formPath)
	if err != nil {
		return errors.Wrap(err, "unable to read the form file")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	vf, err := m.ValidateForm(org, rawForms)
	return printer.SmartPrint(p, vf, err, "unable validate form", printer.Options{}, cmd.OutOrStdout())
}
