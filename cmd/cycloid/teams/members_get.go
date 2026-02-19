package teams

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/spf13/cobra"
)

func NewTeamMemberGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get [member_id]",
		Short:             "Get a team member data.",
		Example:           "cy team member get --team my-team 17281 member1 member@example.org",
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: cyargs.CompleteTeamMemberAny,
		RunE:              getTeamMember,
	}

	return cmd
}

func getTeamMember(cmd *cobra.Command, args []string) error {
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
	currentMembers, err := m.ListTeamMembers(org, team)
	if err != nil {
		return printer.SmartPrint(p, nil, fmt.Errorf("failed to list members of team %q: %w", team, err), "", printer.Options{}, cmd.OutOrStderr())
	}

	for i, memberArg := range args {
		if index := slices.IndexFunc(currentMembers, func(m *models.MemberTeam) bool {
			return memberArg == strconv.Itoa(int(ptr.Value(m.ID))) ||
				memberArg == m.Username ||
				memberArg == m.Email.String()
		}); index != -1 {
			outMembers[i] = currentMembers[i]
		}
	}

	if len(args) == 1 {
		return printer.SmartPrint(p, outMembers[0], nil, "", printer.Options{}, cmd.OutOrStdout())
	} else {
		return printer.SmartPrint(p, outMembers, nil, "", printer.Options{}, cmd.OutOrStdout())
	}
}
