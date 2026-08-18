package credentials

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

// AddAllRawFlags registers every flag that any credential type may consume.
// Use this when a single command needs to accept any credential type inline
// (e.g. `cy cloud-account create --new-credential aws --access-key …`).
//
// Per-type subcommands (used by `cy credential create <type>`) should keep
// registering only the flags they need to preserve a tight UX.
func AddAllRawFlags(cmd *cobra.Command) {
	cyargs.AddCredentialCanonicalFlag(cmd)
	cyargs.AddCredentialPathFlag(cmd)
	cyargs.AddCredentialSSHKeyFlag(cmd)
	cyargs.AddCredentialUsernameFlag(cmd)
	cyargs.AddCredentialPasswordFlag(cmd)
	cyargs.AddCredentialFieldFlag(cmd)
	cyargs.AddCredentialFieldFileFlag(cmd)
	cyargs.AddCredentialAccessKeyFlag(cmd)
	cyargs.AddCredentialSecretKeyFlag(cmd)
	cyargs.AddCredentialClientIDFlag(cmd)
	cyargs.AddCredentialClientSecretFlag(cmd)
	cyargs.AddCredentialSubscriptionIDFlag(cmd)
	cyargs.AddCredentialTenantIDFlag(cmd)
	cyargs.AddCredentialAccountNameFlag(cmd)
	cyargs.AddCredentialJSONKeyFlag(cmd)
	cyargs.AddCredentialCaCertFlag(cmd)
	cyargs.AddCredentialAuthURLFlag(cmd)
	cyargs.AddCredentialDomainIDFlag(cmd)
}

// BuildCredentialRaw extracts the raw credential payload from the command flags
// based on credType. Shared by `cy credential` per-type subcommands and
// `cy cloud-account` inline credential creation so both stay in sync.
func BuildCredentialRaw(cmd *cobra.Command, credType string) (*models.CredentialRaw, error) {
	switch credType {
	case "ssh":
		sshKeyPath, err := cyargs.GetCredentialSSHKey(cmd)
		if err != nil {
			return nil, err
		}
		sshKey, err := os.ReadFile(sshKeyPath)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read SSH key")
		}
		return &models.CredentialRaw{SSHKey: string(sshKey)}, nil

	case "basic_auth":
		username, err := cyargs.GetCredentialUsername(cmd)
		if err != nil {
			return nil, err
		}
		password, err := cyargs.GetCredentialPassword(cmd)
		if err != nil {
			return nil, err
		}
		return &models.CredentialRaw{Username: username, Password: password}, nil

	case "custom":
		fields, err := cyargs.GetCredentialField(cmd)
		if err != nil {
			return nil, err
		}
		fileFields, err := cyargs.GetCredentialFieldFile(cmd)
		if err != nil {
			return nil, err
		}
		if len(fields) == 0 && len(fileFields) == 0 {
			return nil, fmt.Errorf("at least one --field or --field-file has to be specified")
		}
		for field, path := range fileFields {
			content, err := os.ReadFile(path)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("unable to read file at path %q", path))
			}
			fields[field] = strings.TrimSuffix(string(content), "\n")
		}
		return &models.CredentialRaw{Raw: fields}, nil

	case "aws":
		accessKey, err := cyargs.GetCredentialAccessKey(cmd)
		if err != nil {
			return nil, err
		}
		secretKey, err := cyargs.GetCredentialSecretKey(cmd)
		if err != nil {
			return nil, err
		}
		return &models.CredentialRaw{AccessKey: accessKey, SecretKey: secretKey}, nil

	case "azure":
		clientID, err := cyargs.GetCredentialClientID(cmd)
		if err != nil {
			return nil, err
		}
		clientSecret, err := cyargs.GetCredentialClientSecret(cmd)
		if err != nil {
			return nil, err
		}
		subscriptionID, err := cyargs.GetCredentialSubscriptionID(cmd)
		if err != nil {
			return nil, err
		}
		tenantID, err := cyargs.GetCredentialTenantID(cmd)
		if err != nil {
			return nil, err
		}
		return &models.CredentialRaw{
			ClientID:       clientID,
			ClientSecret:   clientSecret,
			SubscriptionID: subscriptionID,
			TenantID:       tenantID,
		}, nil

	case "azure_storage":
		accessKey, err := cyargs.GetCredentialAccessKey(cmd)
		if err != nil {
			return nil, err
		}
		accountName, err := cyargs.GetCredentialAccountName(cmd)
		if err != nil {
			return nil, err
		}
		return &models.CredentialRaw{AccessKey: accessKey, AccountName: accountName}, nil

	case "gcp":
		jsonKeyPath, err := cyargs.GetCredentialJSONKey(cmd)
		if err != nil {
			return nil, err
		}
		jsonKey, err := os.ReadFile(jsonKeyPath)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read JSON key")
		}
		return &models.CredentialRaw{JSONKey: string(jsonKey)}, nil

	case "elasticsearch":
		username, err := cyargs.GetCredentialUsername(cmd)
		if err != nil {
			return nil, err
		}
		password, err := cyargs.GetCredentialPassword(cmd)
		if err != nil {
			return nil, err
		}
		caCertPath, err := cyargs.GetCredentialCaCert(cmd)
		if err != nil {
			return nil, err
		}
		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read CA cert")
		}
		return &models.CredentialRaw{
			Username: username,
			Password: password,
			CaCert:   string(caCert),
		}, nil

	case "swift":
		username, err := cyargs.GetCredentialUsername(cmd)
		if err != nil {
			return nil, err
		}
		password, err := cyargs.GetCredentialPassword(cmd)
		if err != nil {
			return nil, err
		}
		tenantID, err := cyargs.GetCredentialTenantID(cmd)
		if err != nil {
			return nil, err
		}
		authURL, err := cyargs.GetCredentialAuthURL(cmd)
		if err != nil {
			return nil, err
		}
		domainID, err := cyargs.GetCredentialDomainID(cmd)
		if err != nil {
			return nil, err
		}
		return &models.CredentialRaw{
			Username: username,
			Password: password,
			TenantID: tenantID,
			AuthURL:  authURL,
			DomainID: domainID,
		}, nil

	default:
		return nil, fmt.Errorf("unsupported credential type %q", credType)
	}
}
