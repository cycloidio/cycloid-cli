package stacks

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
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
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	cmd.Flags().String("stack-ref", "", "stack reference, format 'org:stack-canonical'")
	cmd.MarkFlagRequired("stack-ref")
	cmd.Flags().String("visibility", "", "update stack visibility")
	cmd.Flags().String("team", "", "update the maintainer team canonical")

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

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
		visibilityStr, err := flagSet.GetString("visibility")
		if err != nil {
			return err
		}
		visibility = &visibilityStr
	} else {
		visibility = stack.Visibility
	}

	if flagSet.Changed("team") {
		team, err = flagSet.GetString("team")
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
	return printer.SmartPrint(p, s, err, fmt.Sprintf("fail to update stack with ref: %s", stackRef), printer.Options{}, cmd.OutOrStdout())
}
