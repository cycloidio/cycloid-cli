package components

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewComponentConfigGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Args:    cobra.NoArgs,
		Short:   "Fetch the current Stackforms variables of a component in JSON format.",
		RunE:    getComponentConfig,
		Example: "cy config get -p project -e env -c component",
	}
	cyargs.AddCyContext(cmd)
	return cmd
}

func getComponentConfig(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	output, err := cyargs.GetOutput(cmd)
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
