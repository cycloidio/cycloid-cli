package teams

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
)

func NewGetTeamCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get [team_canonicals...]",
		Short:             "Get one or more teams information",
		RunE:              getTeam,
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeam,
	}

	return cmd
}

func getTeam(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
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

	var teams = make([]*models.Team, len(args))
	for i, canonical := range args {
		team, err := m.GetTeam(org, canonical)
		if err != nil {
			return printer.SmartPrint(p, nil, err, fmt.Sprintf("failed to get team %q: %s", canonical, err.Error()), printer.Options{}, cmd.OutOrStderr())
		}

		if len(args) == 1 {
			return printer.SmartPrint(p, team, nil, "", printer.Options{}, cmd.OutOrStdout())
		}

		teams[i] = team
	}

	return printer.SmartPrint(p, teams, nil, "", printer.Options{}, cmd.OutOrStdout())
}
