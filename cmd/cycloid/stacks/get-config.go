package stacks

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewGetConfigCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "get-config <ref?> <use-case?> [flags]",
		Short: "output a V2 stack default configuration in JSON (require stackforms)",
		Example: `
cy --org my-org stacks get-config my:stack-ref stack-usecase
`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    getConfig,
		Args:    cobra.RangeArgs(0, 2),
	}
	cmd.Flags().StringP("ref", "r", "", "referential of the stack")
	cmd.Flags().StringP("use-case", "u", "", "usecase you want")

	return cmd
}

func getConfig(cmd *cobra.Command, args []string) error {
	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return err
	}

	ref, _ := cmd.Flags().GetString("ref")
	if len(args) >= 1 && ref == "" {
		ref = args[0]
	} else if ref == "" {
		return fmt.Errorf("missing ref argument.")
	}

	useCase, _ := cmd.Flags().GetString("use-case")
	if len(args) == 2 && useCase == "" {
		useCase = args[1]
	} else if useCase == "" {
		return fmt.Errorf("missing use-case argument.")
	}

	internal.Debug("ref: ", ref, "usecase:", useCase)

	output, err := cmd.Flags().GetString("output")
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

	data, err := m.GetStackConfig(org, ref)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get the stack configuration", printer.Options{}, cmd.OutOrStdout())
	}

	var mappedData map[string]struct {
		Forms common.UseCase `json:"forms"`
	}

	// Type casting is not working but marshall/unmashall works
	// TODO: Clean this or ask backend to return a typed response
	jsonData, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "failed to parse config from API.")
	}

	err = json.Unmarshal(jsonData, &mappedData)
	if err != nil {
		return errors.Wrap(err, "failed to parse forms usecase from API.")
	}

	formConfig := common.UseCaseToFormInput(mappedData[useCase].Forms, true)
	return printer.SmartPrint(p, formConfig, err, "unable to get stack config", printer.Options{}, cmd.OutOrStdout())
}
