package teams

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
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

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	outMembers := make([]*models.MemberTeam, len(args))
	currentMembers, _, err := m.ListTeamMembers(org, team)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, fmt.Errorf("failed to list members of team %q: %w", team, err), "", printer.Options{})
	}

	for i, memberArg := range args {
		if index := slices.IndexFunc(currentMembers, func(m *models.MemberTeam) bool {
			emailStr := ""
			if m.Email != nil {
				emailStr = m.Email.String()
			}
			return memberArg == strconv.Itoa(int(ptr.Value(m.ID))) ||
				memberArg == m.Username ||
				memberArg == emailStr
		}); index != -1 {
			outMembers[i] = currentMembers[i]
		}
	}

	if len(args) == 1 {
		return cyout.PrintWithOptions(cmd, outMembers[0], nil, "", printer.Options{})
	}
	return cyout.PrintWithOptions(cmd, outMembers, nil, "", printer.Options{})
}
