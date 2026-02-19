package teams

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/internal/utils"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("failed to get printer for output type %q: %w", output, err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	_, canonical, err := middleware.NameOrCanonical(ptr.Ptr(teamName), ptr.Ptr(team))
	if err != nil {
		return fmt.Errorf("failed to infer canonical: %w", err)
	}

	currentTeam, err := m.GetTeam(org, canonical)
	if err != nil {
		return fmt.Errorf("failed to Get the team to update with canonical %q: %w", canonical, err)
	}

	newTeam, err := m.UpdateTeam(
		org, ptr.Ptr(utils.CoalesceNonZero(teamName, ptr.Value(currentTeam.Name))),
		currentTeam.Canonical, ptr.Ptr(utils.CoalesceNonZero(teamOwner, ptr.Value(ptr.Value(currentTeam.Owner).Username))), roles,
	)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to UpdateTeam", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, newTeam, nil, "", printer.Options{}, cmd.OutOrStdout())
}
