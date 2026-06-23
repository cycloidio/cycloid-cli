package credentials

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var credentialTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "Type", "Path", "Keys"},
	Identifier: "Canonical",
}

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Short:   "list the credentials",
		Example: `cy --org my-org credentials list -o json`,
		RunE:    list,
	}

	cyargs.AddCredentialTypeFlag(cmd)
	cyout.RegisterModel(cmd, models.CredentialSimple{})
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	credT, err := cyargs.GetCredentialType(cmd)
	if err != nil {
		return err
	}

	creds, _, err := m.ListCredentials(org, credT)
	return cyout.PrintWithOptions(cmd, creds, err, "unable to list credentials", credentialTableOptions)
}
