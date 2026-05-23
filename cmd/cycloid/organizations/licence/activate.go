package licence

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewActivateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate",
		Short: "Activate or replace the Cycloid licence on this organization",
		Long: `Activate (or replace) the Cycloid licence on this organization.

The API overwrites any existing licence in place; running this command twice
is safe and the second run replaces the first.

Examples:
  cy organization licence activate --org root-org --key eyJhbG...
  cy organization licence activate --org root-org --key-file ./licence.jwt
  cat licence.jwt | cy organization licence activate --org root-org`,
		Args: cobra.NoArgs,
		RunE: activate,
	}

	cyargs.AddLicenceKeyFlag(cmd)
	cyargs.AddLicenceKeyFileFlag(cmd)
	cmd.MarkFlagsMutuallyExclusive("key", "key-file")

	return cmd
}

func activate(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	key, err := cyargs.GetLicenceKey(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	_, err = m.ActivateLicence(org, key)
	return cyout.PrintWithOptions(cmd, nil, err, "failed to activate licence", printer.Options{})
}
