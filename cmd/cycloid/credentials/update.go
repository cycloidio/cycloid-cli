package credentials

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
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
	cmd.PersistentFlags().Bool("update", false, "update this credential if it already exists.")
	err := cmd.PersistentFlags().MarkHidden("update")
	if err != nil {
		panic(fmt.Sprintf("We should be able to mark this flag hidden: %s", err.Error()))
	}

	// SSH
	var sshCmd = &cobra.Command{
		Use:     "ssh",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for SSH
	cy --org my-org credential update ssh --name foo --ssh-key /path/to/private/key
`,
	}
	sshCmd.MarkFlagRequired(cyargs.AddCredentialSSHKeyFlag(sshCmd))

	// Basic auth
	var basicAuthCmd = &cobra.Command{
		Use:     "basic_auth",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for basic authentication
	cy --org my-org credential update basic_auth --name foo --username my-username --password my-password
`,
	}
	cyargs.AddCredentialUsernameFlag(basicAuthCmd)
	cyargs.AddCredentialPasswordFlag(basicAuthCmd)

	// Custom
	var customCmd = &cobra.Command{
		Use:     "custom",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for custom type
	cy --org my-org credential update custom --name foo --field my-key=my-value --field my-key2=my-value2 --field-file my-key3=/file/path
`,
	}

	cyargs.AddCredentialFieldFlag(customCmd)
	cyargs.AddCredentialFieldFileFlag(customCmd)

	// AWS
	var awsCmd = &cobra.Command{
		Use:     "aws",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for AWS
	cy --org my-org credential update aws --name foo --access-key foo --secret-key bar
`,
	}
	awsCmd.MarkFlagRequired(cyargs.AddCredentialAccessKeyFlag(awsCmd))
	awsCmd.MarkFlagRequired(cyargs.AddCredentialSecretKeyFlag(awsCmd))

	// Azure
	var azureCmd = &cobra.Command{
		Use:     "azure",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
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
	var azureStorageCmd = &cobra.Command{
		Use:     "azure_storage",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for Azure storage
	cy --org my-org credential update azure_storage --name foo --account-name myaccount --access-key mykey
`,
	}
	azureStorageCmd.MarkFlagRequired(cyargs.AddCredentialAccountNameFlag(azureStorageCmd))
	azureStorageCmd.MarkFlagRequired(cyargs.AddCredentialAccessKeyFlag(azureStorageCmd))

	// GCP
	var gcpCmd = &cobra.Command{
		Use:     "gcp",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for GCP
	cy --org my-org credential update gcp --name foo --json-key /path/to/json/key
`,
	}
	gcpCmd.MarkFlagRequired(cyargs.AddCredentialJSONKeyFlag(gcpCmd))

	// Swift
	var swiftCmd = &cobra.Command{
		Use:     "swift",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
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
	var elasticsearchCmd = &cobra.Command{
		Use:     "elasticsearch",
		Args:    cobra.NoArgs,
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
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
	m := middleware.NewMiddleware(api)

	var err error
	var rawCred *models.CredentialRaw

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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	switch credT {
	case "ssh":
		sshKeyPath, err := cyargs.GetCredentialSSHKey(cmd)
		if err != nil {
			return err
		}

		sshKey, err := os.ReadFile(sshKeyPath)
		if err != nil {
			return errors.Wrap(err, "unable to read SSH key")
		}

		rawCred = &models.CredentialRaw{
			SSHKey: string(sshKey),
		}
	case "basic_auth":
		username, err := cyargs.GetCredentialUsername(cmd)
		if err != nil {
			return err
		}
		password, err := cyargs.GetCredentialPassword(cmd)
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			Username: username,
			Password: password,
		}
	case "custom":
		fields, err := cyargs.GetCredentialField(cmd)
		if err != nil {
			return err
		}
		fileFields, err := cyargs.GetCredentialFieldFile(cmd)
		if err != nil {
			return err
		}

		if len(fields) == 0 && len(fileFields) == 0 {
			return fmt.Errorf("at least one --field or --field-file has to be specified")
		}

		// Read file fields
		if len(fileFields) > 0 {
			for f, p := range fileFields {
				fc, err := os.ReadFile(p)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("unable to read file path %s", p))
				}

				fields[f] = strings.TrimSuffix(string(fc), "\n")
			}
		}

		rawCred = &models.CredentialRaw{
			Raw: fields,
		}
	case "aws":
		accessKey, err := cyargs.GetCredentialAccessKey(cmd)
		if err != nil {
			return err
		}
		secretKey, err := cyargs.GetCredentialSecretKey(cmd)
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}
	case "azure":
		clientID, err := cyargs.GetCredentialClientID(cmd)
		if err != nil {
			return err
		}
		clientSecret, err := cyargs.GetCredentialClientSecret(cmd)
		if err != nil {
			return err
		}
		subscriptionID, err := cyargs.GetCredentialSubscriptionID(cmd)
		if err != nil {
			return err
		}
		tenantID, err := cyargs.GetCredentialTenantID(cmd)
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			SubscriptionID: subscriptionID,
			TenantID:       tenantID,
		}
	case "azure_storage":
		accessKey, err := cyargs.GetCredentialAccessKey(cmd)
		if err != nil {
			return err
		}
		accountName, err := cyargs.GetCredentialAccountName(cmd)
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			AccessKey:   accessKey,
			AccountName: accountName,
		}
	case "gcp":
		jsonKeyPath, err := cyargs.GetCredentialJSONKey(cmd)
		if err != nil {
			return err
		}

		jsonKey, err := os.ReadFile(jsonKeyPath)
		if err != nil {
			return errors.Wrap(err, "unable to read JSON key")
		}

		rawCred = &models.CredentialRaw{
			JSONKey: string(jsonKey),
		}
	case "swift":
		username, err := cyargs.GetCredentialUsername(cmd)
		if err != nil {
			return err
		}

		password, err := cyargs.GetCredentialPassword(cmd)
		if err != nil {
			return err
		}

		authURL, err := cyargs.GetCredentialAuthURL(cmd)
		if err != nil {
			return err
		}

		domainID, err := cyargs.GetCredentialDomainID(cmd)
		if err != nil {
			return err
		}

		tenantID, err := cyargs.GetCredentialTenantID(cmd)
		if err != nil {
			return err
		}

		rawCred = &models.CredentialRaw{
			Username: username,
			Password: password,
			AuthURL:  authURL,
			DomainID: domainID,
			TenantID: tenantID,
		}

	case "elasticsearch":
		caCertPath, err := cyargs.GetCredentialCaCert(cmd)
		if err != nil {
			return err
		}

		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			return errors.Wrap(err, "unable to read CA cert file")
		}

		username, err := cyargs.GetCredentialUsername(cmd)
		if err != nil {
			return err
		}

		password, err := cyargs.GetCredentialPassword(cmd)
		if err != nil {
			return err
		}

		rawCred = &models.CredentialRaw{
			Username: username,
			Password: password,
			CaCert:   string(caCert),
		}

	default:
		return fmt.Errorf("unsupported credential type: %s", credT)
	}

	outCred, err := m.UpdateCredential(org, name, credT, rawCred, credentialPath, credential, description)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to update credential", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, outCred, nil, "", printer.Options{}, cmd.OutOrStdout())
}
