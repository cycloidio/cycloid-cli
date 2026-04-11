package teams

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewTeamMemberListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list [team_canonical]",
		Short:             "List members of a team",
		Example:           "cy team member list --team my-team",
		Args:              cobra.MaximumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeam,
		RunE:              listTeamMember,
	}

	cyargs.AddTeamFlag(cmd)
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

	if team == "" && len(args) == 1 {
		team = args[0]
	} else {
		return fmt.Errorf("missing team canonical parameter, give it by argument or flag")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	members, _, err := m.ListTeamMembers(org, team)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, fmt.Errorf("failed to list members of team %q: %w", team, err), "", printer.Options{})
	}

	return cyout.PrintWithOptions(cmd, members, nil, "", printer.Options{})
}
