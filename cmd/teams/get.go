package teams

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewGetTeamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get [team_canonicals...]",
		Short:             "Get one or more teams information",
		RunE:              getTeam,
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeam,
	}
	cyout.RegisterModel(cmd, models.Team{})
	return cmd
}

func getTeam(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	teams := make([]*models.Team, len(args))
	for i, canonical := range args {
		team, _, err := m.GetTeam(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err,
				fmt.Sprintf("failed to get team %q", canonical), teamTableOptions)
		}

		if len(args) == 1 {
			return cyout.PrintWithOptions(cmd, team, nil, "", teamTableOptions)
		}

		teams[i] = team
	}

	return cyout.PrintWithOptions(cmd, teams, nil, "", teamTableOptions)
}
