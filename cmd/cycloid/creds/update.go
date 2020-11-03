package creds

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "update",
		Hidden:  true,
		Short:   "not implemented yet",
		Long:    `not implemented yet`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("not implemented yet")
			return nil
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
