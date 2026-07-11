package teams

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewTeamMemberAssignCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "assign [usernames|emails...]",
		Short:             "Assign an organization member to a team by username or email.",
		Example:           "cy team member assign --team my-team member1 member@example.org",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeamMemberEmailOrUsername,
		RunE:              assignTeamMember,
	}

	cmd.MarkFlagRequired(cyargs.AddTeamFlag(cmd))
	return cmd
}

func assignTeamMember(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	team, err := cyargs.GetTeam(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	outMembers := make([]*models.MemberTeam, len(args))
	for i, id := range args {
		var username, email *string = nil, nil
		if strings.Contains(id, "@") {
			email = &id
		} else {
			username = &id
		}

		outMembers[i], _, err = m.AssignMemberToTeam(org, team, username, email)
		if err != nil {
			return cyout.PrintWithOptions(cmd, outMembers, fmt.Errorf("failed to assign member %q in team %q: %w", id, team, err), "", printer.Options{})
		}
	}

	return cyout.PrintWithOptions(cmd, outMembers, nil, "", printer.Options{})
}
