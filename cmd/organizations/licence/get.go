package licence

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

var licenceTableOptions = printer.Options{
	Columns:    []string{"CompanyName", "EmailAddress", "ExpiresAt", "MembersCount", "Version", "Key"},
	Identifier: "CompanyName",
}

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "Get the currently active licence",
		Example: `cy organization licence get --org root-org -o yaml`,
		Args:    cobra.NoArgs,
		RunE:    get,
	}

	cyout.RegisterModel(cmd, models.Licence{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	licence, _, err := m.GetLicence(org)
	return cyout.PrintWithOptions(cmd, licence, err, "failed to get licence", licenceTableOptions)
}
