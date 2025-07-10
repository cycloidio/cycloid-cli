package credentials

import (
	"fmt"
	"io/ioutil"
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

	WithPersistentFlagDescription(cmd)
	WithPersistentFlagName(cmd)
	common.RequiredPersistentFlag(common.WithFlagCan, cmd)
	common.WithPersistentFlagCan(cmd)
	WithPersistentFlagPath(cmd)

	// SSH
	var ssh = &cobra.Command{
		Use:     "ssh",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for SSH
	cy --org my-org credential update ssh --name foo --ssh-key /path/to/private/key
`,
	}
	common.RequiredFlag(WithFlagSSHKey, ssh)

	// Basic auth
	var basicAuth = &cobra.Command{
		Use:     "basic_auth",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for basic authentication
	cy --org my-org credential update basic_auth --name foo --username my-username --password my-password
`,
	}
	common.RequiredFlag(WithFlagUsername, basicAuth)
	common.RequiredFlag(WithFlagPassword, basicAuth)

	// Custom
	var custom = &cobra.Command{
		Use:     "custom",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for custom type
	cy --org my-org credential update custom --name foo --field my-key=my-value --field my-key2=my-value2 --field-file my-key3=/file/path
`,
	}
	WithFlagField(custom)
	WithFlagFieldFile(custom)

	// AWS
	var aws = &cobra.Command{
		Use:     "aws",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for AWS
	cy --org my-org credential update aws --name foo --access-key foo --secret-key bar
`,
	}
	common.RequiredFlag(WithFlagAccessKey, aws)
	common.RequiredFlag(WithFlagSecretKey, aws)

	// Azure
	var azure = &cobra.Command{
		Use:     "azure",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for Azure
	cy --org my-org credential update azure --name foo --client-id myid --client-secret mysec --subscription-id mysub --tenant-id mytenant
`,
	}
	common.RequiredFlag(WithFlagClientID, azure)
	common.RequiredFlag(WithFlagClientSecret, azure)
	common.RequiredFlag(WithFlagSubscriptionID, azure)
	common.RequiredFlag(WithFlagTenantID, azure)

	// Azure Storage
	var azureStorage = &cobra.Command{
		Use:     "azure_storage",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for Azure storage
	cy --org my-org credential update azure_storage --name foo --account-name myaccount --access-key mykey
`,
	}
	common.RequiredFlag(WithFlagAccountName, azureStorage)
	common.RequiredFlag(WithFlagAccessKey, azureStorage)

	// GCP
	var gcp = &cobra.Command{
		Use:     "gcp",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for GCP
	cy --org my-org credential update gcp --name foo --json-key /path/to/json/key
`,
	}
	common.RequiredFlag(WithFlagJsonKey, gcp)

	// Swift
	var swift = &cobra.Command{
		Use:     "swift",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for Swift
	cy --org my-org credential update swift --name foo --username foo --password bar --tenant-id mytenant --auth-url url --domain-id mydomain
`,
	}
	common.RequiredFlag(WithFlagUsername, swift)
	common.RequiredFlag(WithFlagPassword, swift)
	common.RequiredFlag(WithFlagTenantID, swift)
	common.RequiredFlag(WithFlagAuthUrl, swift)
	common.RequiredFlag(WithFlagDomainID, swift)

	// Elasticsearch
	var elasticsearch = &cobra.Command{
		Use:     "elasticsearch",
		RunE:    update,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# update a credential for Elasticsearch
	cy --org my-org credential update elasticsearch --name foo --username foo --password bar --ca-cert /path/to/cert
`,
	}
	WithFlagUsername(elasticsearch)
	WithFlagPassword(elasticsearch)
	common.RequiredFlag(WithFlagCaCert, elasticsearch)

	// Command
	cmd.AddCommand(custom, basicAuth, ssh, aws, azure, azureStorage, gcp, swift, elasticsearch)

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
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	path, err := cmd.Flags().GetString("path")
	if err != nil {
		return err
	}
	can, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
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
		sshKeyPath, err := cmd.Flags().GetString("ssh-key")
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
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			Username: username,
			Password: password,
		}
	case "custom":
		fields, err := cmd.Flags().GetStringToString("field")
		if err != nil {
			return err
		}
		fileFields, err := cmd.Flags().GetStringToString("field-file")
		if err != nil {
			return err
		}

		if len(fields) == 0 && len(fileFields) == 0 {
			return fmt.Errorf("at least one --field or --field-file has to be specified")
		}

		// Read file fields
		if len(fileFields) > 0 {
			for f, p := range fileFields {
				fc, err := ioutil.ReadFile(p)
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
		accessKey, err := cmd.Flags().GetString("access-key")
		if err != nil {
			return err
		}
		secretKey, err := cmd.Flags().GetString("secret-key")
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			AccessKey: accessKey,
			SecretKey: secretKey,
		}
	case "azure":
		clientID, err := cmd.Flags().GetString("client-id")
		if err != nil {
			return err
		}
		clientSecret, err := cmd.Flags().GetString("client-secret")
		if err != nil {
			return err
		}
		subscriptionID, err := cmd.Flags().GetString("subscription-id")
		if err != nil {
			return err
		}
		tenantID, err := cmd.Flags().GetString("tenant-id")
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
		accessKey, err := cmd.Flags().GetString("access-key")
		if err != nil {
			return err
		}
		accountName, err := cmd.Flags().GetString("account-name")
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			AccessKey:   accessKey,
			AccountName: accountName,
		}
	case "gcp":
		jsonKeyPath, err := cmd.Flags().GetString("json-key")
		if err != nil {
			return err
		}

		jsonKey, err := ioutil.ReadFile(jsonKeyPath)
		if err != nil {
			return errors.Wrap(err, "unable to read JSON key")
		}

		rawCred = &models.CredentialRaw{
			JSONKey: string(jsonKey),
		}
	case "swift":
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}
		authURL, err := cmd.Flags().GetString("auth-url")
		if err != nil {
			return err
		}
		domainID, err := cmd.Flags().GetString("domain-id")
		if err != nil {
			return err
		}
		tenantID, err := cmd.Flags().GetString("tenant-id")
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
		caCertPath, err := cmd.Flags().GetString("ca-cert")
		if err != nil {
			return err
		}

		caCert, err := ioutil.ReadFile(caCertPath)
		if err != nil {
			return errors.Wrap(err, "unable to read CA cert file")
		}
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}
		password, err := cmd.Flags().GetString("password")
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

	_, err = m.UpdateCredential(org, name, credT, rawCred, path, can, description)
	return printer.SmartPrint(p, nil, err, "unable to update credential", printer.Options{}, cmd.OutOrStdout())
}
