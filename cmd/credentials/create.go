package credentials

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [ssh|basic_auth|custom|aws|azure|azure_storage|gcp|elasticsearch|swift]",
		Short: "create a credential",
		Example: `# create a credential for basic authentication
	cy --org my-org credential create basic_auth --name foo --username my-username --password my-password

	# create a credential for SSH
	cy --org my-org credential create ssh --name foo --ssh-key /path/to/private/key

	# create a credential for AWS
	cy --org my-org credential create aws --name foo --access-key foo --secret-key bar

	# create a credential for Azure
	cy --org my-org credential create azure --name foo --client-id myid --client-secret mysec --subscription-id mysub --tenant-id mytenant

	# create a credential for Azure storage
	cy --org my-org credential create azure_storage --name foo --account-name myaccount --access-key mykey

	# create a credential for GCP
	cy --org my-org credential create gcp --name foo --json-key /path/to/json/key

	# create a credential for Elasticsearch
	cy --org my-org credential create elasticsearch --name foo --username foo --password bar --ca-cert /path/to/cert

	# create a credential for Swift
	cy --org my-org credential create swift --name foo --username foo --password bar --tenant-id mytenant --auth-url url --domain-id mydomain
`,
	}
	cyargs.AddCredentialNamePersistentFlag(cmd)
	cyargs.AddCredentialDescriptionPersistentFlag(cmd)
	cmd.MarkFlagRequired(cyargs.AddCredentialCanonicalPersistentFlag(cmd))
	cyargs.AddCredentialPathPersistentFlag(cmd)
	cmd.PersistentFlags().Bool(cyargs.UpdateFlag, false, "update this credential if it already exists.")

	sshCmd := &cobra.Command{
		Use:  "ssh",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for SSH
	cy --org my-org credential create ssh --name foo --ssh-key /path/to/private/key
`,
	}
	sshCmd.MarkFlagRequired(cyargs.AddCredentialSSHKeyFlag(sshCmd))

	basicAuthCmd := &cobra.Command{
		Use:  "basic_auth",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for basic authentication
	cy --org my-org credential create basic_auth --name foo --username my-username --password my-password
`,
	}
	cyargs.AddCredentialUsernameFlag(basicAuthCmd)
	cyargs.AddCredentialPasswordFlag(basicAuthCmd)

	customCmd := &cobra.Command{
		Use:  "custom",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for custom type
	cy --org my-org credential create custom --name foo --field my-key=my-value --field my-key2=my-value2 --field-file my-key3=/file/path
`,
	}
	cyargs.AddCredentialFieldFlag(customCmd)
	cyargs.AddCredentialFieldFileFlag(customCmd)

	awsCmd := &cobra.Command{
		Use:  "aws",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for AWS
	cy --org my-org credential create aws --name foo --access-key foo --secret-key bar
`,
	}
	awsCmd.MarkFlagRequired(cyargs.AddCredentialAccessKeyFlag(awsCmd))
	awsCmd.MarkFlagRequired(cyargs.AddCredentialSecretKeyFlag(awsCmd))

	azureCmd := &cobra.Command{
		Use:  "azure",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for Azure
	cy --org my-org credential create azure --name foo --client-id myid --client-secret mysec --subscription-id mysub --tenant-id mytenant
`,
	}
	azureCmd.MarkFlagRequired(cyargs.AddCredentialClientIDFlag(azureCmd))
	azureCmd.MarkFlagRequired(cyargs.AddCredentialClientSecretFlag(azureCmd))
	azureCmd.MarkFlagRequired(cyargs.AddCredentialSubscriptionIDFlag(azureCmd))
	azureCmd.MarkFlagRequired(cyargs.AddCredentialTenantIDFlag(azureCmd))

	azureStorageCmd := &cobra.Command{
		Use:  "azure_storage",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for Azure storage
	cy --org my-org credential create azure_storage --name foo --account-name myaccount --access-key mykey
`,
	}
	azureStorageCmd.MarkFlagRequired(cyargs.AddCredentialAccountNameFlag(azureStorageCmd))
	azureStorageCmd.MarkFlagRequired(cyargs.AddCredentialAccessKeyFlag(azureStorageCmd))

	gcpCmd := &cobra.Command{
		Use:  "gcp",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for GCP
	cy --org my-org credential create gcp --name foo --json-key /path/to/json/key
`,
	}
	gcpCmd.MarkFlagRequired(cyargs.AddCredentialJSONKeyFlag(gcpCmd))

	swiftCmd := &cobra.Command{
		Use:  "swift",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for Swift
	cy --org my-org credential create swift --name foo --username foo --password bar --tenant-id mytenant --auth-url url --domain-id mydomain
`,
	}
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialUsernameFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialPasswordFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialTenantIDFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialAuthURLFlag(swiftCmd))
	swiftCmd.MarkFlagRequired(cyargs.AddCredentialDomainIDFlag(swiftCmd))

	elasticsearchCmd := &cobra.Command{
		Use:  "elasticsearch",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for Elasticsearch
	cy --org my-org credential create elasticsearch --name foo --username foo --password bar --ca-cert /path/to/cert
`,
	}
	cyargs.AddCredentialUsernameFlag(elasticsearchCmd)
	cyargs.AddCredentialPasswordFlag(elasticsearchCmd)
	elasticsearchCmd.MarkFlagRequired(cyargs.AddCredentialCaCertFlag(elasticsearchCmd))

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

func create(cmd *cobra.Command, args []string) error {
	credentialType := cmd.CalledAs()
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
	credentialPath = defaultCredentialPath(credentialPath, credential, name)

	description, err := cyargs.GetCredentialDescription(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	if cyargs.GetUpdate(cmd) {
		creds, _, err := m.ListCredentials(org, credentialType)
		if err != nil {
			return fmt.Errorf("failed to create --update credential, cannot check for existing credential (canonical=%q path=%q name=%q): %w",
				credential, credentialPath, name, err)
		}

		existingCredential := findCredentialForUpdate(creds, credential, credentialPath, name)
		if existingCredential != nil {
			// update requires the canonical in the route: infer it when the user identified
			// the credential by path/name (the common create --name --update flow).
			if credential == "" && existingCredential.Canonical != nil {
				credential = *existingCredential.Canonical
				if err := cmd.Flags().Set("canonical", credential); err != nil {
					return fmt.Errorf("failed to set credential canonical before update: %w", err)
				}
			}

			return update(cmd, args)
		}
	}

	rawCred, err := BuildCredentialRaw(cmd, credentialType)
	if err != nil {
		return err
	}

	outCredential, _, err := m.CreateCredential(org, name, credentialType, rawCred, credentialPath, credential, description)
	return cyout.PrintWithOptions(cmd, outCredential, err, "unable to create credential", printer.Options{})
}

func defaultCredentialPath(path, canonical, name string) string {
	if path != "" {
		return path
	}
	if canonical != "" {
		return pathFromCanonical(canonical)
	}
	if name != "" {
		return pathFromCanonical(name)
	}
	return ""
}

func findCredentialForUpdate(creds []*models.CredentialSimple, canonical, path, name string) *models.CredentialSimple {
	if canonical != "" {
		for _, credential := range creds {
			if credential != nil && credential.Canonical != nil && *credential.Canonical == canonical {
				return credential
			}
		}
	}

	if path != "" {
		for _, credential := range creds {
			if credential != nil && credential.Path != nil && *credential.Path == path {
				return credential
			}
		}
	}

	if name != "" {
		for _, credential := range creds {
			if credential != nil && credential.Name != nil && *credential.Name == name {
				return credential
			}
		}
	}

	return nil
}
