package orgplugins

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewWidgetListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "list plugin widgets",
		Example: `
	# List all plugin widgets
	cy --org my-org plugin widget list

	# List plugin widgets filtered by placement
	cy --org my-org plugin widget list --placement dashboard
`,
		RunE: listWidgets,
	}

	cmd.Flags().String("placement", "", "filter widgets by placement")
	cmd.MarkFlagRequired("placement")

	return cmd
}

func listWidgets(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	placement, err := cmd.Flags().GetString("placement")
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

	result, _, err := m.ListPluginWidgets(org, placement)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to list plugin widgets", printer.Options{}, cmd.OutOrStderr())
	}
	return printer.SmartPrint(p, result, nil, "", printer.Options{}, cmd.OutOrStdout())
}
