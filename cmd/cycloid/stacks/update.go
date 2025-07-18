package stacks

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Args:  cobra.NoArgs,
		Short: "update the visibility of a stack",
		Example: `cy stacks update --stack-ref org:myStack --visibility shared

# Full args example
cy stacks update \
	--stack-ref my:stack-ref \
	--visibility "hidden" \
	--team "teamCanonical"
`,
		RunE: update,
	}

	cmd.MarkFlagRequired(cyargs.AddStackRefFlag(cmd))
	cyargs.AddVisibilityFlag(cmd)
	cyargs.AddTeamFlag(cmd)
	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	stackRef, err := cyargs.GetStackRef(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Fetch the current stack state
	stack, err := m.GetStack(org, stackRef)
	if err != nil {
		return printer.SmartPrint(p, nil, err, fmt.Sprintf("failed to retrieve the stack with stack ref: %s", stackRef), printer.Options{}, cmd.OutOrStderr())
	}

	// Manage optional parameters
	flagSet := cmd.Flags()

	var visibility *string
	var team string
	if flagSet.Changed("visibility") {
		visibilityStr, err := cyargs.GetVisibility(cmd)
		if err != nil {
			return err
		}
		visibility = &visibilityStr
	} else {
		visibility = stack.Visibility
	}

	if flagSet.Changed("team") {
		team, err = cyargs.GetTeam(cmd)
		if err != nil {
			return err
		}
	} else {
		if stack.Team != nil {
			team = *stack.Team.Canonical
		}
	}

	// Send request
	s, err := m.UpdateStack(org, stackRef, team, visibility)
	if err != nil {
		return printer.SmartPrint(p, nil, err, fmt.Sprintf("fail to update stack with ref: %s", stackRef), printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, s, nil, "", printer.Options{}, cmd.OutOrStdout())
}
