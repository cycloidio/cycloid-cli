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

func NewConfigGetCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "get <ref?> <use-case?> [flags]",
		Short:   "output a V2 stack default configuration in JSON (require stackforms)",
		Example: `cy --org my-org stacks get-config my:stack-ref stack-usecase`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    getConfig,
		Args:    cobra.RangeArgs(0, 2),
	}

	cyargs.AddStackRefFlag(cmd)
	cyargs.AddUseCaseFlag(cmd)
	return cmd
}

func getConfig(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	stackRef, _ := cyargs.GetStackRef(cmd)
	if len(args) >= 1 && stackRef == "" {
		stackRef = args[0]
	} else if stackRef == "" {
		return fmt.Errorf("missing ref argument")
	}

	useCase, _ := cyargs.GetUseCase(cmd)
	if len(args) == 2 && *useCase == "" {
		useCase = &args[1]
	} else if *useCase == "" {
		return fmt.Errorf("missing use-case argument")
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// Default output is in JSON
	if output == "table" {
		output = "json"
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	stackConfigs, err := m.GetStackConfig(org, stackRef)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get the stack configuration", printer.Options{}, cmd.OutOrStderr())
	}

	useCaseConfig, err := common.FormUseCaseToFormVars(stackConfigs, *useCase)
	if err != nil {
		return fmt.Errorf("failed to parse default value for stack '%s' with use-case '%s': %s", stackRef, *useCase, err)
	}

	return printer.SmartPrint(p, useCaseConfig, nil, "", printer.Options{}, cmd.OutOrStdout())
}
