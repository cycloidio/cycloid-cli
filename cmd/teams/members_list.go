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

func NewTeamMemberListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list [team_canonical]",
		Short:             "List members of a team",
		Example:           "cy team member list --team my-team",
		ValidArgsFunction: cyargs.CompleteTeam,
		RunE:              listTeamMember,
	}

	teamFlag := cyargs.AddTeamFlag(cmd)
	cmd.Args = cobra.MatchAll(cobra.MaximumNArgs(1), cyargs.RequireArgsOrFlag(teamFlag))
	return cmd
}

func listTeamMember(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	team, err := cyargs.GetTeam(cmd)
	if err != nil {
		return err
	}

	// Args guarantees at least one of the two is present.
	if team != "" && len(args) == 1 {
		return fmt.Errorf("team given twice: --team %q and argument %q", team, args[0])
	}
	if team == "" {
		team = args[0]
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	members, _, err := m.ListTeamMembers(org, team)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, fmt.Errorf("failed to list members of team %q: %w", team, err), "", printer.Options{})
	}

	return cyout.PrintWithOptions(cmd, members, nil, "", printer.Options{})
}
