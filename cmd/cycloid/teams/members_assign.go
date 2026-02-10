package teams

import (
	"fmt"
	"strings"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
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

	var outMembers = make([]*models.MemberTeam, len(args))
	for i, id := range args {
		var username, email *string = nil, nil
		if strings.Contains(id, "@") {
			email = &id
		} else {
			username = &id
		}

		outMembers[i], err = m.AssignMemberToTeam(org, team, username, email)
		if err != nil {
			return printer.SmartPrint(p, outMembers, fmt.Errorf("failed to assign member %q in team %q: %w", id, team, err), "", printer.Options{}, cmd.OutOrStderr())
		}
	}

	return printer.SmartPrint(p, outMembers, nil, "", printer.Options{}, cmd.OutOrStdout())
}
