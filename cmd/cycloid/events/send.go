package events

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

var (
	colorFlag       string
	iconFlag        string
	messageFlag     string
	messageFileFlag string
	severityFlag    string
	titleFlag       string
	typeFlag        string
	tagsFlag        map[string]string
)

func NewSendCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Args:  cobra.NoArgs,
		Short: "create an event",
		Example: `
	# create a custom event
	cy --org my-org event create --tag env=staging --title success --message "successful deployment"
`,
		RunE: send,
	}

	cmd.Flags().StringToStringVar(&tagsFlag, "tag", nil, "tags of the event (key=value)")
	cmd.MarkFlagRequired("tag")
	cmd.Flags().StringVar(&titleFlag, "title", "", "title of the event")
	cmd.MarkFlagRequired("title")

	cmd.Flags().StringVar(&colorFlag, "color", "", "color of the event")
	cmd.Flags().StringVar(&messageFlag, "message", "", "message of the event")
	cmd.Flags().StringVar(&messageFileFlag, "message-file", "", "path to the file holding the message")
	cmd.Flags().StringVar(&iconFlag, "icon", "", "icon of the event")
	cmd.Flags().StringVar(&severityFlag, "severity", "info", "severity of the event")
	cmd.Flags().StringVar(&typeFlag, "type", "Custom", "type of the event")

	return cmd
}

func send(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	color, err := cmd.Flags().GetString("color")
	if err != nil {
		return err
	}
	severity, err := cmd.Flags().GetString("severity")
	if err != nil {
		return err
	}
	eType, err := cmd.Flags().GetString("type")
	if err != nil {
		return err
	}
	title, err := cmd.Flags().GetString("title")
	if err != nil {
		return err
	}
	message, err := cmd.Flags().GetString("message")
	if err != nil {
		return err
	}
	messageFile, err := cmd.Flags().GetString("message-file")
	if err != nil {
		return err
	}
	tags, err := cmd.Flags().GetStringToString("tag")
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

	var msg string
	if messageFile != "" {
		rawMsg, err := os.ReadFile(messageFile)
		if err != nil {
			return errors.Wrap(err, "unable to read message file")
		}
		msg = string(rawMsg)
	} else if message != "" {
		msg = message
	} else {
		return fmt.Errorf("required flag(s) \"message\" or \"message-file\" not set")
	}

	err = m.SendEvent(org, eType, title, msg, severity, tags, color)
	return printer.SmartPrint(p, nil, err, "unable to send event", printer.Options{}, cmd.OutOrStdout())
}
