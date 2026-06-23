package bootstrapfirstorg

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

var bootstrapTableOptions = printer.Options{
	Columns: []string{"Org", "Username", "Email", "Token", "APIKey", "CredentialCanonical"},
	Transform: func(obj interface{}) map[string]string {
		data, ok := obj.(*apiclient.FirstOrgData)
		if !ok || data == nil {
			return nil
		}

		row := map[string]string{
			"Org":      data.Org,
			"Username": data.Username,
			"Email":    data.Email,
			"Token":    maskSecret(data.Token),
			"Password": maskSecret(data.Password),
		}
		if data.APIKey != nil {
			row["APIKey"] = maskSecret(*data.APIKey)
		}
		if data.CredentialCanonical != nil {
			row["CredentialCanonical"] = *data.CredentialCanonical
		}
		return row
	},
}

func bootstrap(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	username, err := cmd.Flags().GetString("username")
	if err != nil {
		return err
	}
	fullName, err := cmd.Flags().GetString("full-name")
	if err != nil {
		return err
	}
	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return err
	}
	password, err := cyargs.GetBootstrapPassword(cmd)
	if err != nil {
		return err
	}
	licence, err := cyargs.GetLicence(cmd)
	if err != nil {
		return err
	}
	apiKeyCanonical, err := cyargs.GetBootstrapAPIKeyCanonical(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	result, _, err := m.InitFirstOrg(org, username, fullName, email, password, licence, apiKeyCanonical)
	return cyout.PrintWithOptions(cmd, result, err, "failed to bootstrap first organization", bootstrapTableOptions)
}

func maskSecret(value string) string {
	if len(value) <= 5 {
		return strings.Repeat("*", len(value))
	}
	return "***" + value[len(value)-5:]
}
