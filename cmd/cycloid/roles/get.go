package roles

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var (
		example = `cy --org my-org roles get my-role`
		short   = "Get a role specification."
		long    = short
	)

	var cmd = &cobra.Command{
		Use: "get",
		Args: cobra.MatchAll(
			cobra.MaximumNArgs(1),
		),
		Example:           example,
		Short:             short,
		Long:              long,
		RunE:              getRole,
		ValidArgsFunction: cyargs.CompleteRoleCanonicals,
	}

	cyargs.AddRoleCanonicalFlag(cmd)

	// keep legacy flag just in case
	// TODO: deprecate in next update
	cmd.Flags().String("canonical", "", "the role canonical")
	cmd.Flags().MarkDeprecated("canonical", "use --role or pass the canonical as argument directly")

	return cmd
}

func getRole(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	role, err := cyargs.GetRoleCanonical(cmd)
	if err != nil {
		return err
	}

	// flag has precedence
	if role == "" && len(args) == 1 {
		role = args[0]
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	mb, err := m.GetRole(org, role)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get role", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, mb, nil, "", printer.Options{}, cmd.OutOrStdout())
}
