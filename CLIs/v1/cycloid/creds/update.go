package creds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "update",
		Short:  "...",
		Hidden: true,
		Long:   `........ . . .... .. .. ....`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("...")
		},
	}

	return cmd
}

// /organizations/{organization_canonical}/credentials/{credential_id}
// put: updateCredential
// Update an existing Credential, based on the type you will have to pass different parameters within the body:
// * ssh: ssh_key
// * aws: access_key, secret_key
// * gcp: json_key
// * azure: client_id, client_secret, subscription_id, tenant_id
// * azure_storage: account_name, access_key
// * basic_auth: username, password
// * elasticsearch: username, password, ca_cert
// * swift: auth_url, username, password, domain_id, tenant_id
