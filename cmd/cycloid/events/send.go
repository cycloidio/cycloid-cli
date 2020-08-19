package events

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/client/client/organizations"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

var colorFlag string
var iconFlag string
var messageFlag string
var messageFileFlag string
var severityFlag string
var titleFlag string
var typeFlag string
var tagsFlag map[string]string

func NewSendCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "send",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  send,
	}

	cmd.Flags().StringToStringVar(&tagsFlag, "tag", nil, "key=value")
	cmd.MarkFlagRequired("tag")
	cmd.Flags().StringVar(&titleFlag, "title", "", "...")
	cmd.MarkFlagRequired("title")

	cmd.Flags().StringVar(&colorFlag, "color", "", "...")
	cmd.Flags().StringVar(&messageFlag, "message", "", "...")
	cmd.Flags().StringVar(&messageFileFlag, "message-file", "", "...")
	cmd.Flags().StringVar(&iconFlag, "icon", "", "...")
	cmd.Flags().StringVar(&severityFlag, "severity", "info", "...")
	cmd.Flags().StringVar(&typeFlag, "type", "Custom", "...")

	return cmd
}

func send(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
	fmt.Println(args)

	var err error
	var body *models.NewEvent

	org, err := cmd.Flags().GetString("org")
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

	var msg string
	if messageFile != "" {
		rawMsg, err := ioutil.ReadFile(messageFile)
		if err != nil {
			return fmt.Errorf("Message file reading error : %s", err.Error())
		}
		msg = string(rawMsg)
	} else if message != "" {
		msg = message
	} else {
		return errors.New("required flag(s) \"message\" or \"message-file\" not set")
	}

	eventP := organizations.NewSendOrgEventParams()
	eventP.SetOrganizationCanonical(org)

	var ts []*models.Tag

	for k, v := range tags {
		tag := &models.Tag{
			Key:   &k,
			Value: &v,
		}
		err = tag.Validate(strfmt.Default)
		if err != nil {
			continue
		}

		ts = append(ts, tag)
	}

	body = &models.NewEvent{
		Tags:     ts,
		Type:     &eType,
		Title:    &title,
		Color:    color,
		Severity: &severity,
		Message:  &msg,
	}

	eventP.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.Organizations.SendOrgEvent(eventP, root.ClientCredentials())
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}

// '/organizations/{organization_canonical}/events':
// post: sendOrgEvent
// Send a event on the organization to be registered.
