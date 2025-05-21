package events

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"

	"time"
)

func NewListCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		Short:   "list events",
		Example: `# Get events since last week
cy --org my-org event list --begin 0 --end "$(date --date "last week" +"%s")" --severity info,warn,err,crit --type Cycloid,AWS,Monitoring,Custom
`,
		RunE:    list,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	now := time.Now().Unix()
	yesterday := time.Now().AddDate(0, 0, -1).Unix()
	cmd.Flags().Uint64P("begin", "b", uint64(yesterday), "the starting date to get events from, using a unix timestamp, default to yesterday")
	cmd.Flags().Uint64P("end", "e", uint64(now), "the starting date to get events from, using a unix timestamp, default to now")

	cmd.Flags().StringSliceP("severity", "s", []string{"info", "err", "warn", "crit"}, "filter events by severity, you can set more than one, default '-s info,err,warn,crit'")
	cmd.Flags().StringSliceP("type", "t", []string{"Cycloid", "Custom", "AWS", "Monitoring"}, "filter events by type, you can set more than one, default '-t Cycloid,Custom,AWS,Monitoring'")
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	begin, err := cmd.Flags().GetUint64("begin")
	if err != nil {
		return err
	}

	end, err := cmd.Flags().GetUint64("end")
	if err != nil {
		return err
	}

	eventSeverity, err := cmd.Flags().GetStringSlice("severity")
	if err != nil {
		return err
	}
	eventType, err := cmd.Flags().GetStringSlice("type")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	events, err := m.ListEvents(org, eventType, eventSeverity, begin*1000, end*1000)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to list events", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, events, err, "", printer.Options{}, cmd.OutOrStdout())
}
