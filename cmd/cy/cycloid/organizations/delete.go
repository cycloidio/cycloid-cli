package organizations

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cy/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewDeleteCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete",
		Args:  cobra.NoArgs,
		Short: "delete an organization (require root API_KEY)",
		Example: `
	# delete an organization with canonical name my-org
	# The API_KEY must be obtained from the root organization
	cy organization delete --org my-org
`,
		RunE:    del,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	return cmd
}

func del(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable get org flag")
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	err = m.DeleteOrganization(org)
	return printer.SmartPrint(p, nil, err, "unable to delete organization", printer.Options{}, cmd.OutOrStdout())
}
