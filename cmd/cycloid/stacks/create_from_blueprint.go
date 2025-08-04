package stacks

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// CreateFromBlueprintOutput represents the formatted output for create-from-blueprint command
type CreateFromBlueprintOutput struct {
	Canonical   string `json:"canonical"`
	Name        string `json:"name"`
	Ref         string `json:"ref"`
	Description string `json:"description"`
}

func NewCreateFromBlueprintCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "create-from-blueprint",
		Short:   "create a new stack from a blueprint",
		Example: `cy --org my-org stack create-from-blueprint --blueprint-ref repo:blueprint-canonical --name "My Stack" --canonical my-stack --service-catalog-source-canonical my-catalog --use-case production`,
		RunE:    createFromBlueprint,
		Args:    cobra.NoArgs,
	}

	cmd.Flags().String("blueprint-ref", "", "Blueprint reference to use as template (required)")
	cmd.MarkFlagRequired("blueprint-ref")

	cmd.Flags().String("name", "", "Name of the new stack (required)")
	cmd.MarkFlagRequired("name")

	cmd.Flags().String("canonical", "", "Canonical name (slug) of the new stack (required)")
	cmd.MarkFlagRequired("canonical")

	cmd.Flags().String("service-catalog-source-canonical", "", "Service catalog source canonical (required)")
	cmd.MarkFlagRequired("service-catalog-source-canonical")

	cmd.Flags().String("use-case", "", "Use case canonical to apply from the blueprint (required)")
	cmd.MarkFlagRequired("use-case")

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

	// Get required flags
	blueprintRef, err := cmd.Flags().GetString("blueprint-ref")
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	serviceCatalogSourceCanonical, err := cmd.Flags().GetString("service-catalog-source-canonical")
	if err != nil {
		return err
	}

	useCase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return err
	}

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
		return printer.SmartPrint(p, nil, err, "failed to create stack from blueprint", printer.Options{}, cmd.OutOrStdout())
	}

	// For JSON output, return all fields; for table output, return only specific fields
	if output == "json" {
		return printer.SmartPrint(p, stack, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	// For table output, format to show only the requested fields
	formattedOutput := &CreateFromBlueprintOutput{
		Canonical:   getStringValue(stack.Canonical),
		Name:        getStringValue(stack.Name),
		Ref:         getStringValue(stack.Ref),
		Description: stack.Description,
	}

	return printer.SmartPrint(p, formattedOutput, nil, "", printer.Options{}, cmd.OutOrStdout())
}
