package stacks

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateFromBlueprintCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create",
		Short:   "create a new stack from a blueprint",
		Example: `cy stack create --name "My Stack" --stack my-stack --catalog-repository my-catalog --blueprint-ref org:blueprint --use-case production`,
		RunE:    createFromBlueprint,
		Args:    cobra.NoArgs,
	}

	cmd.MarkFlagRequired(cyargs.AddStackNameFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddBlueprintRefFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddStackFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddCatalogRepositoryFlag(cmd))
	cmd.MarkFlagRequired(cyargs.AddUseCaseFlag(cmd))
	cmd.Flags().Bool("update", false, "if the stack canonical already exists, return it without failing (idempotent create)")
	return cmd
}

func createFromBlueprint(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	blueprintRef, err := cyargs.GetBlueprintRef(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	stack, err := cyargs.GetStack(cmd)
	if err != nil {
		return err
	}

	displayName, stackCanonical, err := middleware.NameOrCanonical(&name, &stack)
	if err != nil {
		return err
	}

	catalogRepository, err := cyargs.GetCatalogRepository(cmd)
	if err != nil {
		return err
	}

	useCasePtr, err := cyargs.GetUseCase(cmd)
	if err != nil {
		return err
	}
	useCase := useCasePtr

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	stacksList, _, err := m.ListStacks(org)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "failed to list stacks to check if stack exists", printer.Options{})
	}

	idx := slices.IndexFunc(stacksList, func(s *models.ServiceCatalog) bool {
		return s.Canonical != nil && *s.Canonical == stackCanonical
	})
	if idx != -1 {
		ref := ptr.Value(stacksList[idx].Ref)
		if ref == "" {
			return cyout.PrintWithOptions(cmd, nil,
				fmt.Errorf("stack with canonical %q has no ref", stackCanonical),
				"failed to create stack from blueprint", printer.Options{})
		}
		if update {
			outStack, _, err := m.GetStack(org, ref)
			return cyout.PrintWithOptions(cmd, outStack, err, "failed to get existing stack", printer.Options{})
		}
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("stack %q already exists; use --update for an idempotent create", stackCanonical),
			"failed to create stack from blueprint", printer.Options{})
	}

	createdStack, _, err := m.CreateStackFromBlueprint(org, blueprintRef, displayName, stackCanonical, catalogRepository, useCase)
	return cyout.PrintWithOptions(cmd, createdStack, err, "failed to create stack from blueprint", printer.Options{})
}
