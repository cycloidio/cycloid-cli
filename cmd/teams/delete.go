package teams

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewDeleteTeamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "delete [team_canonicals...]",
		Short:             "Delete a team and output the deleted teams canonicals",
		RunE:              deleteTeam,
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeam,
	}

	return cmd
}

func deleteTeam(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	deleted := make([]string, len(args))
	for i, team := range args {
		_, err = m.DeleteTeam(org, team)
		if err != nil {
			return cyout.PrintWithOptions(cmd, deleted, err, fmt.Sprintf("failed to delete team %q: %s", team, err.Error()), printer.Options{})
		}

		deleted[i] = team
	}

	return cyout.PrintWithOptions(cmd, deleted, nil, "", printer.Options{Columns: []string{"Canonical"}})
}
