package credentials

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/client/organization_credentials"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get",
		Args:  cobra.NoArgs,
		Short: "get a credential",
		Example: `
	# get a credential by its canonical
	cy --org my-org credential get --canonical my-cred
`,
		RunE:    get,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredFlag(common.WithFlagCan, cmd)
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	can, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	params := organization_credentials.NewGetCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(can)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	c, err := m.GetCredential(org, can)
	return printer.SmartPrint(p, c, err, "unable to get credential", printer.Options{}, cmd.OutOrStdout())
}
