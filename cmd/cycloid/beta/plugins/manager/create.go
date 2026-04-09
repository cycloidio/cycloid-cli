package manager

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
	cmd := &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "[beta] Invite a plugin manager to the organization",
		Example: `
  cy beta plugin manager create --name my-manager --url https://pm.example.com
`,
		RunE: createPluginManager,
	}

	cmd.MarkFlagRequired(cyargs.AddNameFlag(cmd))
	cmd.Flags().String("url", "", "URL of the plugin manager instance (required)")
	cmd.MarkFlagRequired("url")
	return cmd
}

func createPluginManager(cmd *cobra.Command, args []string) error {
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

	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return errors.Wrap(err, "unable to get --url flag")
	}

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	result, _, err := m.CreatePluginManager(org, name, url)
	return printer.SmartPrint(p, result, err, "unable to create plugin manager", printer.Options{}, cmd.OutOrStdout())
}
