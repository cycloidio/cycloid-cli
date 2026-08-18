package teams

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/utils"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func NewUpdateTeamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing team.",
		Args:  cobra.NoArgs,
		RunE:  updateTeam,
	}

	cmd.MarkFlagsOneRequired(
		cyargs.AddTeamNameFlag(cmd),
		cyargs.AddTeamFlag(cmd),
	)
	cmd.MarkFlagRequired(cyargs.AddTeamRolesFlag(cmd))
	cyargs.AddTeamOwnerFlag(cmd)
	return cmd
}

func updateTeam(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	team, err := cyargs.GetTeam(cmd)
	if err != nil {
		return err
	}

	teamName, err := cyargs.GetTeamName(cmd)
	if err != nil {
		return err
	}

	teamOwner, err := cyargs.GetTeamOwner(cmd)
	if err != nil {
		return err
	}

	roles, err := cyargs.GetTeamRoles(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	_, canonical, err := apiclient.NameOrCanonical(ptr.Ptr(teamName), ptr.Ptr(team))
	if err != nil {
		return fmt.Errorf("failed to infer canonical: %w", err)
	}

	currentTeam, _, err := m.GetTeam(org, canonical)
	if err != nil {
		return fmt.Errorf("failed to Get the team to update with canonical %q: %w", canonical, err)
	}

	newTeam, _, err := m.UpdateTeam(
		org, ptr.Ptr(utils.CoalesceNonZero(teamName, ptr.Value(currentTeam.Name))),
		currentTeam.Canonical, ptr.Ptr(utils.CoalesceNonZero(teamOwner, ptr.Value(ptr.Value(currentTeam.Owner).Username))), roles,
	)
	return cyout.PrintWithOptions(cmd, newTeam, err, "failed to UpdateTeam", printer.Options{})
}
