package teams

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}
	if output == "table" {
		output = "json"
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("failed to get printer for output type %q: %w", output, err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var deleted = make([]string, len(args))
	for i, team := range args {
		err = m.DeleteTeam(org, team)
		if err != nil {
			return printer.SmartPrint(p, deleted, err, fmt.Sprintf("failed to delete team %q: %s", team, err.Error()), printer.Options{}, cmd.OutOrStderr())
		}

		deleted[i] = team
	}

	return printer.SmartPrint(p, deleted, nil, "", printer.Options{}, cmd.OutOrStdout())
}
