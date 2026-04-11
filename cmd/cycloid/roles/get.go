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

	// flag has precedence
	if role == "" && len(args) == 1 {
		role = args[0]
	}

	mb, _, err := m.GetRole(org, role)
	return cyout.PrintWithOptions(cmd, mb, err, "unable to get role", roleTableOptions)
}
