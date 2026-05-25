package cloudaccounts

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-account",
		Aliases: []string{"cloud-accounts", "cloudaccount", "ca"},
		Short:   "Manage cloud accounts",
	}

	cmd.AddCommand(
		NewCreateCommand(),
		NewUpdateCommand(),
		NewGetCommand(),
		NewDeleteCommand(),
		NewListCommand(),
	)
	return cmd
}

var cloudAccountTableOptions = printer.Options{
	Columns:    []string{"Canonical", "Name", "CloudProvider"},
	Identifier: "Canonical",
	Transform: func(obj interface{}) map[string]string {
		account, ok := obj.(*models.CloudAccountDetail)
		if !ok {
			if simple, ok := obj.(*models.CloudAccount); ok {
				account = &models.CloudAccountDetail{CloudAccount: *simple}
			} else {
				return map[string]string{}
			}
		}
		canonical, name, provider := "", "", ""
		if account.Canonical != nil {
			canonical = *account.Canonical
		}
		if account.Name != nil {
			name = *account.Name
		}
		if account.CloudProvider != nil {
			provider = *account.CloudProvider
		}
		return map[string]string{
			"Canonical":     canonical,
			"Name":          name,
			"CloudProvider": provider,
		}
	},
}
