package stacks

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/custommodels"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Aliases: []string{
			"ls",
		},
		Args:    cobra.NoArgs,
		Short:   "list stacks available in the current org.",
		Example: `cy my-org stack list`,
		RunE:    list,
	}

	cmd.Flags().Bool("blueprint", false, "list blueprint stacks")
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	blueprint, err := cmd.Flags().GetBool("blueprint")
	if err != nil {
		return err
	}

	// Initialize apiclient after all arguments are collected
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	// list standard stacks
	if !blueprint {
		stacks, _, err := m.ListStacks(org)
		return cyout.PrintWithOptions(cmd, stacks, err, "failed to list stacks from API", printer.Options{})
	}

	// list --blueprints
	stacks, _, err := m.ListBlueprints(org)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "failed to list blueprints from API", printer.Options{})
	}

	blueprints := make([]*custommodels.Blueprint, len(stacks))
	for index, stack := range stacks {
		blueprints[index] = &custommodels.Blueprint{
			ServiceCatalog: *stack,
			UseCases:       &[]string{},
		}

		if stack.Ref != nil {
			// Pass empty version flags to use default/latest version
			stackUseCases, _, err := m.ListStackUseCases(org, *stack.Ref, "", "", "")
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), "error: failed to fetch use cases for blueprint '%s': %s\n", *stack.Ref, err.Error())
				continue
			}

			var useCases []string
			for _, useCase := range stackUseCases {
				useCases = append(useCases, *useCase.UseCase)
			}

			blueprints[index].UseCases = &useCases
		}
	}

	return cyout.PrintWithOptions(cmd, blueprints, nil, "failed to list blueprints from API", printer.Options{})
}
