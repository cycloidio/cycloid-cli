package teams

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/utils"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
)

func NewListTeamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list",
		Short:             "List teams in current organization.",
		RunE:              listTeam,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cyargs.CompleteTeam,
	}

	cyargs.AddTeamNameFlag(cmd)
	cyargs.AddTeamCreatedAt(cmd)
	cyargs.AddTeamOwnerFlag(cmd)
	cyargs.AddTeamMemberIDFlag(cmd)
	cyargs.AddTeamOrderBy(cmd)
	return cmd
}

func listTeam(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	teamName, err := cyargs.GetTeamName(cmd)
	if err != nil {
		return err
	}

	teamCreatedAt, err := cyargs.GetTeamCreatedAt(cmd)
	if err != nil {
		return err
	}

	teamMemberID, err := cyargs.GetTeamMemberID(cmd)
	if err != nil {
		return err
	}

	teamOrderBy, err := cyargs.GetTeamOrderBy(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("failed to list printer for output type %q: %s", output, err.Error())
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	teams, err := m.ListTeams(
		org, utils.CoalesceNonZeroPtr(teamName),
		teamCreatedAt, utils.CoalesceNonZeroPtr(teamMemberID), teamOrderBy,
	)
	if err != nil {
		return printer.SmartPrint(p, nil, err, fmt.Sprintf("failed to list team: %s", err.Error()), printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, teams, nil, "", printer.Options{}, cmd.OutOrStdout())
}
