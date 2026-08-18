package roles

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteCommand() *cobra.Command {
	var (
		example = `cy --org my-org roles delete [role_canonicals...]`
		short   = "Delete one or more roles from the organization"
		long    = short
	)

	cmd := &cobra.Command{
		Use:               "delete",
		Aliases:           []string{"rm"},
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteRoleCanonical,
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

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	deleted := make([]string, 0, len(args))
	for _, role := range args {
		_, err := m.DeleteRole(org, role)
		if err != nil {
			return cyout.PrintWithOptions(cmd, deleted, err, fmt.Sprintf("failed to delete role %q", role), printer.Options{})
		}

		deleted = append(deleted, role)
	}

	return cyout.PrintWithOptions(cmd, deleted, nil, "", printer.Options{Columns: []string{"Canonical"}})
}
