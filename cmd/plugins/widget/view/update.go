package view

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <widget-view-id>",
		Args:  cobra.ExactArgs(1),
		Short: "Update a widget view (enable/disable and set URL slug)",
		Long: `Update a plugin widget view's enabled state and URL slug.

For inherited widgets (from a parent org), this creates or updates
a local override without modifying the parent's configuration.`,
		Example: `
  # Enable a widget view with a custom URL slug
  cy plugin widget view update 42 --enabled --url-slug my-custom-page

  # Disable a widget view
  cy plugin widget view update 42 --enabled=false --url-slug my-custom-page
`,
		RunE: updateWidgetView,
	}

	cmd.Flags().Bool("enabled", true, "Whether the widget view is enabled")
	cmd.Flags().String("url-slug", "", "URL slug for the widget view (required)")
	_ = cmd.MarkFlagRequired("url-slug")
	return cmd
}

func updateWidgetView(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	n, err := strconv.ParseUint(args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid widget-view-id %q: must be a positive integer", args[0])
	}
	widgetViewID := uint32(n)

	enabled, err := cmd.Flags().GetBool("enabled")
	if err != nil {
		return err
	}

	urlSlug, err := cmd.Flags().GetString("url-slug")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	_, err = m.UpdatePluginWidgetView(org, widgetViewID, enabled, urlSlug)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to update plugin widget view", printer.Options{})
}
