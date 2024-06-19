package stacks

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
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
	cmd.Flags().String("org", "", "organization")
	cmd.MarkFlagRequired("org")
	cmd.Flags().StringP("ref", "r", "", "referential of the stack")
	cmd.Flags().StringP("use-case", "u", "", "usecase you want")

	return cmd
}

func ExtractFormsFromStackConfig(data interface{}, useCase string) (*models.FormUseCase, error) {
	formData, ok := data.(map[string]interface{})
	if !ok {
		return nil, errors.New("failed to cast forms data")
	}

	useCaseData, ok := formData[useCase].(map[string]interface{})
	if !ok {
		return nil, errors.New("failed to extract selected use-case: " + useCase)
	}

	// Type casting is not working but marshall/unmashall works
	// TODO: Clean this or ask backend to return a typed response
	jsonData, err := json.Marshal(useCaseData["forms"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall config")
	}

	var d *models.FormUseCase
	err = json.Unmarshal(jsonData, &d)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse forms usecase, struct invalid.")
	}

	return d, nil
}

func getConfig(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}

	ref, err := cmd.Flags().GetString("ref")
	if len(args) >= 1 && ref == "" {
		ref = args[0]
	} else if ref == "" {
		return fmt.Errorf("missing ref argument.")
	}

	useCase, err := cmd.Flags().GetString("use-case")
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

	s, err := m.GetStackConfig(org, ref)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to get the stack configuration", printer.Options{}, cmd.OutOrStdout())
	}

	data, err := ExtractFormsFromStackConfig(s, useCase)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to parse config from API response", printer.Options{}, cmd.OutOrStdout())
	}

	formConfig, err := common.ParseFormsConfig(data, false)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to parse stack config", printer.Options{}, cmd.OutOrStdout())
	}

	return printer.SmartPrint(p, formConfig, err, "unable to get stack", printer.Options{}, cmd.OutOrStdout())
}
