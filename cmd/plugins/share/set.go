package share

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "set <id-or-name>",
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: cyargs.CompletePluginInstallID,
		Short:             "Set the sharing configuration of a plugin",
		Long: `Set the sharing configuration of a plugin install.

Visibility can be 'local' (only this org) or 'shared' (visible to child orgs).
Mode can be 'include' (only listed orgs) or 'exclude' (all except listed orgs).
Organizations is a comma-separated list of child org canonicals.`,
		Example: `
  # Share with all child organizations
  cy plugin share set my-plugin --visibility shared

  # Share only with specific child organizations
  cy plugin share set my-plugin --visibility shared --mode include --orgs child-org-1,child-org-2

  # Share with all except specific organizations
  cy plugin share set my-plugin --visibility shared --mode exclude --orgs excluded-org

  # Make plugin local (not shared)
  cy plugin share set my-plugin --visibility local
`,
		RunE: setSharing,
	}

	cmd.Flags().String("visibility", "", "Visibility scope: 'local' or 'shared' (required)")
	cmd.Flags().String("mode", "include", "Sharing mode: 'include' or 'exclude'")
	cmd.Flags().String("orgs", "", "Comma-separated list of child organization canonicals")
	_ = cmd.MarkFlagRequired("visibility")
	return cmd
}

func setSharing(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	visibility, err := cmd.Flags().GetString("visibility")
	if err != nil {
		return err
	}

	mode, err := cmd.Flags().GetString("mode")
	if err != nil {
		return err
	}

	orgsStr, err := cmd.Flags().GetString("orgs")
	if err != nil {
		return err
	}

	var organizations []string
	if orgsStr != "" {
		organizations = strings.Split(orgsStr, ",")
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	id, err := cyargs.ResolvePluginInstallID(org, args[0], m)
	if err != nil {
		return err
	}

	_, err = m.SetPluginInstallSharing(org, id, visibility, mode, organizations)
	return cyout.PrintWithOptions(cmd, nil, err, "unable to set plugin sharing", printer.Options{})
}
