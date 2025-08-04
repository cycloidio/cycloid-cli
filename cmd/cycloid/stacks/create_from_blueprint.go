package stacks

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateFromBlueprintCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create-from-blueprint",
		Short:   "create a new stack from a blueprint",
		Example: `cy --org my-org stack create-from-blueprint --blueprint-ref repo:blueprint-canonical --name "My Stack" --canonical my-stack --service-catalog-source-canonical my-catalog --use-case production`,
		RunE:    createFromBlueprint,
		Args:    cobra.NoArgs,
	}

	cmd.MarkFlagRequired(cyargs.AddBlueprintRefFlag(cmd))

	cyargs.AddNameFlag(cmd)
	cmd.MarkFlagRequired("name")

	cmd.MarkFlagRequired(cyargs.AddCanonicalFlag(cmd))

	cmd.MarkFlagRequired(cyargs.AddServiceCatalogSourceCanonicalFlag(cmd))

	cmd.MarkFlagRequired(cyargs.AddUseCaseFlag(cmd))

	return cmd
}

func createFromBlueprint(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return err
	}

	// Get required flags using cyargs functions
	blueprintRef, err := cyargs.GetBlueprintRef(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetCanonical(cmd)
	if err != nil {
		return err
	}

	serviceCatalogSourceCanonical, err := cyargs.GetServiceCatalogSourceCanonical(cmd)
	if err != nil {
		return err
	}

	useCasePtr, err := cyargs.GetUseCase(cmd)
	if err != nil {
		return err
	}
	useCase := *useCasePtr

	// Initialize middleware after all arguments are collected
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Create the stack from blueprint
	stack, err := m.CreateStackFromBlueprint(org, blueprintRef, name, canonical, serviceCatalogSourceCanonical, useCase)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to create stack from blueprint", printer.Options{}, cmd.OutOrStderr())
	}

	// For JSON output, return all fields; for table output, return only specific fields
	if output == "json" {
		return printer.SmartPrint(p, stack, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	// For table output, format to show only the requested fields
	type createFromBlueprintOutput struct {
		Canonical   string `json:"canonical"`
		Name        string `json:"name"`
		Ref         string `json:"ref"`
		Description string `json:"description"`
	}

	formattedOutput := &createFromBlueprintOutput{
		Canonical:   ptr.Value(stack.Canonical),
		Name:        ptr.Value(stack.Name),
		Ref:         ptr.Value(stack.Ref),
		Description: stack.Description,
	}

	return printer.SmartPrint(p, formattedOutput, nil, "", printer.Options{}, cmd.OutOrStdout())
}
