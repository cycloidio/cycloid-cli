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

	stack, err := cyargs.GetStack(cmd)
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

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	createdStack, err := m.CreateStackFromBlueprint(org, blueprintRef, name, stack, catalogRepository, useCase)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to create stack from blueprint", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, createdStack, nil, "", printer.Options{}, cmd.OutOrStdout())
}
