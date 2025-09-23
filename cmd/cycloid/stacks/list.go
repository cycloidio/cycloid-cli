package stacks

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/custommodels"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	blueprint, err := cmd.Flags().GetBool("blueprint")
	if err != nil {
		return err
	}

	// Initialize middleware after all arguments are collected
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("unable to get printer: %w", err)
	}

	// list standard stacks
	if !blueprint {
		stacks, err := m.ListStacks(org)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "failed to list stacks from API", printer.Options{}, cmd.OutOrStderr())
		}

		return printer.SmartPrint(p, stacks, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	// list --blueprints
	// table output is broken with composed structs so we force jso
	if output == "table" {
		output = "json"
	}

	// Re-init the printer
	p, err = factory.GetPrinter(output)
	if err != nil {
		return fmt.Errorf("unable to get printer: %w", err)
	}

	stacks, err := m.ListBlueprints(org)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to list blueprints from API", printer.Options{}, cmd.OutOrStderr())
	}

	var blueprints = make([]*custommodels.Blueprint, len(stacks))
	for index, stack := range stacks {
		blueprints[index] = &custommodels.Blueprint{
			ServiceCatalog: *stack,
			UseCases:       &[]string{},
		}

		if stack.Ref != nil {
			stackUseCases, err := m.ListStackUseCases(org, *stack.Ref)
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

	return printer.SmartPrint(p, blueprints, nil, "", printer.Options{}, cmd.OutOrStdout())
}
