package credentials

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
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
	cmd.PersistentFlags().Bool("update", false, "update this credential if it already exists.")

	// SSH
	var sshCmd = &cobra.Command{
		Use:  "ssh",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for SSH
	cy --org my-org credential create ssh --name foo --ssh-key /path/to/private/key
`,
	}
	sshCmd.MarkFlagRequired(cyargs.AddCredentialSSHKeyFlag(sshCmd))

	// Basic auth
	var basicAuthCmd = &cobra.Command{
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

	// Custom
	var customCmd = &cobra.Command{
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

	// AWS
	var awsCmd = &cobra.Command{
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

	// Azure
	var azureCmd = &cobra.Command{
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

	// Azure Storage
	var azureStorageCmd = &cobra.Command{
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

	// GCP
	var gcpCmd = &cobra.Command{
		Use:  "gcp",
		Args: cobra.NoArgs,
		RunE: create,
		Example: `
	# create a credential for GCP
	cy --org my-org credential create gcp --name foo --json-key /path/to/json/key
`,
	}
	gcpCmd.MarkFlagRequired(cyargs.AddCredentialJSONKeyFlag(gcpCmd))

	// Swift
	var swiftCmd = &cobra.Command{
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

	// Elasticsearch
	var elasticsearchCmd = &cobra.Command{
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

func create(cmd *cobra.Command, args []string) error {
	var err error
	var rawCred *models.CredentialRaw

	credentialTypes := cmd.CalledAs()
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	credential, err := cyargs.GetCredentialCanonical(cmd)
	if err != nil {
		return err
	}

	credentialPath, _ := cyargs.GetCredentialPath(cmd)
	if credentialPath == "" {
		credentialPath = pathFromCanonical(credential)
	}

	name, _ := cyargs.GetCredentialName(cmd)
	if name == "" {
		name = credential
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

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	if allowUpdate, _ := cmd.Flags().GetBool("update"); allowUpdate {
		_, err := m.GetCredential(org, credential)
		if err == nil {
			// if the cred exists, forward the call to the update func.
			return update(cmd, args)
		}
	}

	switch credentialTypes {
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
			for field, path := range fileFields {
				fc, err := os.ReadFile(path)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("unable to read file at path '%s'", path))
				}

				fields[field] = strings.TrimSuffix(string(fc), "\n")
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
		return fmt.Errorf("unsupported credential type: %s", credentialTypes)
	}

	outCredential, err := m.CreateCredential(org, name, credentialTypes, rawCred, credentialPath, credential, description)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "unable to create credential", printer.Options{}, cmd.OutOrStderr())
	}

	return printer.SmartPrint(p, outCredential, nil, "", printer.Options{}, cmd.OutOrStdout())
}
