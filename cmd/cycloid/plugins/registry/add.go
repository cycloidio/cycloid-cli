package registry

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Args:  cobra.NoArgs,
		Short: "Add a plugin registry",
		Example: `
  cy plugin registry add --name my-registry --url https://registry.example.com
`,
		RunE: addPluginRegistry,
	}

	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cyargs.AddURLFlag(cmd, "URL of the plugin registry (required)")
	cmd.MarkFlagRequired("url")
	return cmd
}

func addPluginRegistry(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetName(cmd)
	if err != nil {
		return err
	}

	url, err := cyargs.GetURL(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get --url flag")
	}

	result, _, err := m.CreatePluginRegistry(org, name, url)
	return cyout.PrintWithOptions(cmd, result, err, "unable to add plugin registry", printer.Options{})
}
