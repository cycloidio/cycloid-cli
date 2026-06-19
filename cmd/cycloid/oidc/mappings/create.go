package mappings

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

// NewCreateCommand returns the `mappings create` command.
func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "Create an OIDC group-to-team mapping",
		Example: `
  # Map OIDC group "devs" to team "dev-team" in my-org
  cy --org my-org oidc mappings create --group-name devs --team dev-team
`,
		RunE: createMapping,
	}

	cmd.MarkFlagRequired(cyargs.AddOIDCGroupNameFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddTeamFlag(cmd))

	cyout.RegisterModel(cmd, middleware.OIDCGroupMapping{})
	return cmd
}

func createMapping(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	groupName, err := cyargs.GetOIDCGroupName(cmd)
	if err != nil {
		return err
	}

	team, err := cyargs.GetTeam(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.CreateOIDCGroupMapping(org, groupName, team)
	return cyout.PrintWithOptions(cmd, result, err, "unable to create OIDC group mapping", mappingTableOptions)
}
