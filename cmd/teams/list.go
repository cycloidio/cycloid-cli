package teams

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils"
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
	cyout.RegisterModel(cmd, models.Team{})
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

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	teams, _, err := m.ListTeams(
		org, utils.CoalesceNonZeroPtr(teamName),
		teamCreatedAt, utils.CoalesceNonZeroPtr(teamMemberID), teamOrderBy,
	)
	errMsg := fmt.Sprintf("failed to list teams: %v", err)
	return cyout.PrintWithOptions(cmd, teams, err, errMsg, teamTableOptions)
}
