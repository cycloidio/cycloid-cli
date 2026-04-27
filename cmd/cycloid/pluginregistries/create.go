package pluginregistries

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create a plugin registry",
		Example: `
	# Create a plugin registry
	cy --org my-org plugin-registry create --name my-registry --url https://registry.example.com
`,
		RunE: createRegistry,
	}

	cmd.Flags().String("name", "", "name of the plugin registry")
	cmd.MarkFlagRequired("name")
	cmd.Flags().String("url", "", "URL of the plugin registry")
	cmd.MarkFlagRequired("url")

	return cmd
}

func createRegistry(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}
	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	result, _, err := m.CreatePluginRegistry(org, name, url)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to create plugin registry", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
