package teams

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func NewTeamMemberUnAssignCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "unassign [usernames|emails...]",
		Short:             "Unassign a member from a team.",
		Example:           "cy team member unassign --team my-team member1 member@example.org",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeamMemberAny,
		RunE:              unassignTeamMember,
	}

	cmd.MarkFlagRequired(cyargs.AddTeamFlag(cmd))
	return cmd
}

func unassignTeamMember(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	team, err := cyargs.GetTeam(cmd)
	if err != nil {
		return err
	}

	members, err := anyMembersToID(org, team, args)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	outUnassigned := make([]*uint32, len(members))
	for i, member := range members {
		_, err = m.UnAssignMemberFromTeam(org, team, ptr.Value(member))
		if err != nil {
			return cyout.PrintWithOptions(cmd, outUnassigned, fmt.Errorf("failed to unassign member with ID %d in team %q: %w", member, team, err), "", printer.Options{})
		}

		outUnassigned[i] = member
	}

	return cyout.PrintWithOptions(cmd, outUnassigned, nil, "", printer.Options{})
}
