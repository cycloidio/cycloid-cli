package root

import (
	"io/ioutil"
	"os"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var formsFlag string

func NewValidateFormCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "validate-form",
		Short: "validate a .forms.yml file",
		Example: `
		# validate a stackforms file
		cy validate-form --org my-org --forms .forms.yml
		`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    validateForm,
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	cmd.Flags().StringVar(&formsFlag, "forms", ".forms.yml", "Path to your stackform file, default .forms.yml")

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

	vf, err := m.ValidateForm(org, rawForms)
	if err != nil {
		return errors.Wrap(err, "unable validate form")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// print the result on the standard output
	if err := p.Print(vf, printer.Options{}, os.Stdout); err != nil {
		return errors.Wrap(err, "unable to print result")
	}
	return nil
}
