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

func NewTeamMemberListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "list [team_canonical]",
		Short:             "List members of a team",
		Example:           "cy team member list my-team",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompleteTeam,
		RunE:              listTeamMember,
	}

	return cmd
}

func listTeamMember(cmd *cobra.Command, args []string) error {
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
		return fmt.Errorf("failed to list printer for output type %q: %w", &output, err)
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	members, err := m.ListTeamMembers(org, args[0])
	if err != nil {
		return printer.SmartPrint(p, nil, fmt.Errorf("failed to list members of team %q: %w", args[0], err), "", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, members, nil, "", printer.Options{}, cmd.OutOrStdout())
}
