package roles

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetCommand() *cobra.Command {
	var (
		example = `
	# get a single role
	cy --org my-org roles get my-role

	# get multiple roles
	cy --org my-org roles get role-a role-b
`
		short = "Get a role specification."
		long  = short
	)

	var cmd = &cobra.Command{
		Use:               "get [canonical...]",
		Args:              cyargs.RequireArgsOrFlag("role"),
		Example:           example,
		Short:             short,
		Long:              long,
		RunE:              getRole,
		ValidArgsFunction: cyargs.CompleteRoleCanonical,
	}

	cyargs.AddRoleCanonicalFlag(cmd)

	// keep legacy flag just in case
	// TODO: deprecate in next update
	cmd.Flags().String("canonical", "", "the role canonical")
	cmd.Flags().MarkHidden("canonical")

	cyout.RegisterModel(cmd, models.Role{})
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

	if role != "" {
		found := false
		for _, a := range args {
			if a == role {
				found = true
				break
			}
		}
		if !found {
			args = append(args, role)
		}
	}

	if len(args) == 1 {
		mb, _, err := m.GetRole(org, args[0])
		return cyout.PrintWithOptions(cmd, mb, err, "unable to get role", roleTableOptions)
	}

	results := make([]*models.Role, 0, len(args))
	for _, canonical := range args {
		mb, _, err := m.GetRole(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to get role "+canonical, roleTableOptions)
		}
		results = append(results, mb)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", roleTableOptions)
}
