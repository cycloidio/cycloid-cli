package teams

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
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

	var outUnassigned = make([]*uint32, len(members))
	for i, member := range members {
		err = m.UnAssignMemberFromTeam(org, team, ptr.Value(member))
		if err != nil {
			return printer.SmartPrint(p, outUnassigned, fmt.Errorf("failed to unassign member with ID %d in team %q: %w", member, team, err), "", printer.Options{}, cmd.OutOrStderr())
		}

		outUnassigned[i] = member
	}

	return printer.SmartPrint(p, outUnassigned, nil, "", printer.Options{}, cmd.OutOrStdout())
}
