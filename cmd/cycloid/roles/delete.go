package roles

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	var (
		example = `cy --org my-org roles delete [role_canonicals...]`
		short   = "Remove a user from the organization"
		long    = short
	)

	var cmd = &cobra.Command{
		Use:               "delete",
		Aliases:           []string{"rm"},
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteRoleCanonicals,
		Example:           example,
		Short:             short,
		Long:              long,
		RunE:              deleteRole,
	}

	return cmd
}

func deleteRole(cmd *cobra.Command, args []string) error {
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

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	for _, role := range args {
		err := m.DeleteRole(org, role)
		if err != nil {
			return printer.SmartPrint(p, nil, err, fmt.Sprintf("failed to delete role %q", role), printer.Options{}, cmd.OutOrStderr())
		}
	}

	return printer.SmartPrint(p, nil, nil, "", printer.Options{}, cmd.OutOrStdout())
}
