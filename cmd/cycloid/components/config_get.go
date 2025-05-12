package components

import (
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewGetComponentConfigCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get -p project -e env -c component",
		Short: "Fetch the current Stackforms variables of a component as a JSON.",
		RunE:  getComponentConfig,
	}
	cy_args.AddCyContext(cmd)
	return cmd
}

func getComponentConfig(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	// This endpoint doesn't make sense in table mode
	if output == "table" {
		output = "json"
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	config, err := m.GetComponentConfig(org, project, env, component)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to fetch config of component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, config, nil, "", printer.Options{}, cmd.OutOrStdout())
}
