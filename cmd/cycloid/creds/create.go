package creds

import (
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [ssh|basic_auth|custom|aws|azure|azure_storage|gcp|elasticsearch|swift]",
		Short: "create a credential",
		Example: `
	# create a credential for basic authentication
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

	WithPersistentFlagDescription(cmd)
	common.RequiredPersistentFlag(WithPersistentFlagName, cmd)
	WithPersistentFlagPath(cmd)

	// SSH
	var ssh = &cobra.Command{
		Use:     "ssh",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for SSH
	cy --org my-org credential create ssh --name foo --ssh-key /path/to/private/key
`,
	}
	common.RequiredFlag(WithFlagSSHKey, ssh)

	// Basic auth
	var basicAuth = &cobra.Command{
		Use:     "basic_auth",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for basic authentication
	cy --org my-org credential create basic_auth --name foo --username my-username --password my-password
`,
	}
	common.RequiredFlag(WithFlagUsername, basicAuth)
	common.RequiredFlag(WithFlagPassword, basicAuth)

	// Custom
	var custom = &cobra.Command{
		Use:     "custom",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for custom type
	cy --org my-org credential create custom --name foo --field my-key=my-value --field my-key2=my-value2
`,
	}
	common.RequiredFlag(WithFlagField, custom)

	// AWS
	var aws = &cobra.Command{
		Use:     "aws",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for AWS
	cy --org my-org credential create aws --name foo --access-key foo --secret-key bar
`,
	}
	common.RequiredFlag(WithFlagAccessKey, aws)
	common.RequiredFlag(WithFlagSecretKey, aws)

	// Azure
	var azure = &cobra.Command{
		Use:     "azure",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for Azure
	cy --org my-org credential create azure --name foo --client-id myid --client-secret mysec --subscription-id mysub --tenant-id mytenant
`,
	}
	common.RequiredFlag(WithFlagClientID, azure)
	common.RequiredFlag(WithFlagClientSecret, azure)
	common.RequiredFlag(WithFlagSubscriptionID, azure)
	common.RequiredFlag(WithFlagTenantID, azure)

	// Azure Storage
	var azureStorage = &cobra.Command{
		Use:     "azure_storage",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for Azure storage
	cy --org my-org credential create azure_storage --name foo --account-name myaccount --access-key mykey
`,
	}
	common.RequiredFlag(WithFlagAccountName, azureStorage)
	common.RequiredFlag(WithFlagAccessKey, azureStorage)

	// GCP
	var gcp = &cobra.Command{
		Use:     "gcp",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for GCP
	cy --org my-org credential create gcp --name foo --json-key /path/to/json/key
`,
	}
	common.RequiredFlag(WithFlagJsonKey, gcp)

	// Swift
	var swift = &cobra.Command{
		Use:     "swift",
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for Swift
	cy --org my-org credential create swift --name foo --username foo --password bar --tenant-id mytenant --auth-url url --domain-id mydomain
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
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
		Example: `
	# create a credential for Elasticsearch
	cy --org my-org credential create elasticsearch --name foo --username foo --password bar --ca-cert /path/to/cert
`,
	}
	WithFlagUsername(elasticsearch)
	WithFlagPassword(elasticsearch)
	common.RequiredFlag(WithFlagCaCert, elasticsearch)

	// Command
	cmd.AddCommand(custom, basicAuth, ssh, aws, azure, azureStorage, gcp, swift, elasticsearch)

	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error
	var rawCred *models.CredentialRaw

	credT := cmd.CalledAs()
	org, err := cmd.Flags().GetString("org")
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

		sshKey, err := ioutil.ReadFile(sshKeyPath)
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

	err = m.CreateCredential(org, name, credT, rawCred, path, description)
	return printer.SmartPrint(p, nil, err, "unable to create credential", printer.Options{}, cmd.OutOrStdout())
}
