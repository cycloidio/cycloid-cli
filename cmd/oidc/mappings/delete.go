package mappings

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

// NewDeleteCommand returns the `mappings delete` command.
func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"rm"},
		Args:    cobra.NoArgs,
		Short:   "Delete an OIDC group-to-team mapping by ID",
		Example: `
  # Delete OIDC group mapping with ID 42 in my-org
  cy --org my-org oidc mappings delete --mapping-id 42
`,
		RunE: deleteMapping,
	}

	cmd.MarkFlagRequired(cyargs.AddOIDCMappingIDFlag(cmd))
	return cmd
}

func deleteMapping(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	id, err := cyargs.GetOIDCMappingID(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	_, err = m.DeleteOIDCGroupMapping(org, id)
	return cyout.Print(cmd, nil, err, "unable to delete OIDC group mapping")
}
