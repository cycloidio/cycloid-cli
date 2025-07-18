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

func NewGetComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get -p project -e env -c component",
		Short: "Get the state of a current component.",
		RunE:  getComponent,
	}
	cy_args.AddCyContext(cmd)
	return cmd
}

func getComponent(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cy_args.GetCyContext(cmd)
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	componentState, err := m.GetComponent(org, project, env, component)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to fetch state of component '"+component+"'", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, componentState, nil, "", printer.Options{}, cmd.OutOrStdout())
}
