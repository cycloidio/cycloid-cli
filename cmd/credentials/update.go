package credentials

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update [ssh|basic_auth|custom|aws|azure|azure_storage|gcp|elasticsearch|swift]",
		Args:  cobra.NoArgs,
		Short: "update a credential",
		Example: `
	# update a credential for basic authentication
	cy --org my-org credential update basic_auth --name foo --username my-username --password my-password

	# update a credential for SSH
	cy --org my-org credential update ssh --name foo --ssh-key /path/to/private/key

	# update a credential for AWS
	cy --org my-org credential update aws --name foo --access-key foo --secret-key bar

	# update a credential for Azure
	cy --org my-org credential update azure --name foo --client-id myid --client-secret mysec --subscription-id mysub --tenant-id mytenant

	# update a credential for Azure storage
	cy --org my-org credential update azure_storage --name foo --account-name myaccount --access-key mykey

	# update a credential for GCP
	cy --org my-org credential update gcp --name foo --json-key /path/to/json/key

	# update a credential for Elasticsearch
	cy --org my-org credential update elasticsearch --name foo --username foo --password bar --ca-cert /path/to/cert

	# update a credential for Swift
	cy --org my-org credential update swift --name foo --username foo --password bar --tenant-id mytenant --auth-url url --domain-id mydomain
`,
	}

	cyargs.AddCredentialNamePersistentFlag(cmd)
	cyargs.AddCredentialDescriptionPersistentFlag(cmd)
	cyargs.AddCredentialCanonicalPersistentFlag(cmd)
	cyargs.AddCredentialPathPersistentFlag(cmd)
	cmd.PersistentFlags().Bool(cyargs.UpdateFlag, false, "update this credential if it already exists.")
	if err := cmd.PersistentFlags().MarkHidden(cyargs.UpdateFlag); err != nil {
		panic(fmt.Sprintf("We should be able to mark this flag hidden: %s", err.Error()))
	}

	// SSH
	sshCmd := &cobra.Command{
		Use:  "ssh",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for SSH
	cy --org my-org credential update ssh --name foo --ssh-key /path/to/private/key
`,
	}
	sshCmd.MarkFlagRequired(cyargs.AddCredentialSSHKeyFlag(sshCmd))

	// Basic auth
	basicAuthCmd := &cobra.Command{
		Use:  "basic_auth",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for basic authentication
	cy --org my-org credential update basic_auth --name foo --username my-username --password my-password
`,
	}
	cyargs.AddCredentialUsernameFlag(basicAuthCmd)
	cyargs.AddCredentialPasswordFlag(basicAuthCmd)

	// Custom
	customCmd := &cobra.Command{
		Use:  "custom",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for custom type
	cy --org my-org credential update custom --name foo --field my-key=my-value --field my-key2=my-value2 --field-file my-key3=/file/path
`,
	}

	cyargs.AddCredentialFieldFlag(customCmd)
	cyargs.AddCredentialFieldFileFlag(customCmd)

	// AWS
	awsCmd := &cobra.Command{
		Use:  "aws",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for AWS
	cy --org my-org credential update aws --name foo --access-key foo --secret-key bar
`,
	}
	awsCmd.MarkFlagRequired(cyargs.AddCredentialAccessKeyFlag(awsCmd))
	awsCmd.MarkFlagRequired(cyargs.AddCredentialSecretKeyFlag(awsCmd))

	// Azure
	azureCmd := &cobra.Command{
		Use:  "azure",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for Azure
	cy --org my-org credential update azure --name foo --client-id myid --client-secret mysec --subscription-id mysub --tenant-id mytenant
`,
	}
	azureCmd.MarkFlagRequired(cyargs.AddCredentialClientIDFlag(azureCmd))
	azureCmd.MarkFlagRequired(cyargs.AddCredentialClientSecretFlag(azureCmd))
	azureCmd.MarkFlagRequired(cyargs.AddCredentialSubscriptionIDFlag(azureCmd))
	azureCmd.MarkFlagRequired(cyargs.AddCredentialTenantIDFlag(azureCmd))

	// Azure Storage
	azureStorageCmd := &cobra.Command{
		Use:  "azure_storage",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for Azure storage
	cy --org my-org credential update azure_storage --name foo --account-name myaccount --access-key mykey
`,
	}
	azureStorageCmd.MarkFlagRequired(cyargs.AddCredentialAccountNameFlag(azureStorageCmd))
	azureStorageCmd.MarkFlagRequired(cyargs.AddCredentialAccessKeyFlag(azureStorageCmd))

	// GCP
	gcpCmd := &cobra.Command{
		Use:  "gcp",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for GCP
	cy --org my-org credential update gcp --name foo --json-key /path/to/json/key
`,
	}
	gcpCmd.MarkFlagRequired(cyargs.AddCredentialJSONKeyFlag(gcpCmd))

	// Swift
	swiftCmd := &cobra.Command{
		Use:  "swift",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for Swift
	cy --org my-org credential update swift --name foo --username foo --password bar --tenant-id mytenant --auth-url url --domain-id mydomain
`,
	}
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialUsernameFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialPasswordFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialTenantIDFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialAuthURLFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialDomainIDFlag(swiftCmd))

	// Elasticsearch
	elasticsearchCmd := &cobra.Command{
		Use:  "elasticsearch",
		Args: cobra.NoArgs,
		RunE: update,
		Example: `
	# update a credential for Elasticsearch
	cy --org my-org credential update elasticsearch --name foo --username foo --password bar --ca-cert /path/to/cert
`,
	}

	cyargs.AddCredentialUsernameFlag(elasticsearchCmd)
	cyargs.AddCredentialPasswordFlag(elasticsearchCmd)
	elasticsearchCmd.MarkFlagRequired(cyargs.AddCredentialCaCertFlag(elasticsearchCmd))

	// Command
	cmd.AddCommand(
		customCmd,
		basicAuthCmd,
		sshCmd,
		awsCmd,
		azureCmd,
		azureStorageCmd,
		gcpCmd,
		swiftCmd,
		elasticsearchCmd,
	)

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
func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	credT := cmd.CalledAs()
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	credential, err := cyargs.GetCredentialCanonical(cmd)
	if err != nil {
		return err
	}

	name, _ := cyargs.GetCredentialName(cmd)
	if name == "" {
		name = credential
	}

	credentialPath, _ := cyargs.GetCredentialPath(cmd)
	if credentialPath == "" {
		credentialPath = pathFromCanonical(credential)
	}

	description, err := cyargs.GetCredentialDescription(cmd)
	if err != nil {
		return err
	}

	rawCred, err := BuildCredentialRaw(cmd, credT)
	if err != nil {
		return err
	}

	outCred, _, err := m.UpdateCredential(org, name, credT, rawCred, credentialPath, credential, description)
	return cyout.PrintWithOptions(cmd, outCred, err, "unable to update credential", printer.Options{})
}
