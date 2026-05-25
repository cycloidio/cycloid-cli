package cloudaccounts

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

func createInlineCredential(cmd *cobra.Command, m middleware.Middleware, org, credType, name string) (*models.Credential, error) {
	raw, err := buildCredentialRaw(cmd, credType)
	if err != nil {
		return nil, err
	}

	description, _ := cyargs.GetDescription(cmd)
	path, _ := cyargs.GetCredentialPath(cmd)
	canonical, _ := cyargs.GetCredentialCanonical(cmd)

	cred, _, err := m.CreateCredential(org, name, credType, raw, path, canonical, description)
	if err != nil {
		return nil, err
	}
	return cred, nil
}

func buildAccessCredential(cmd *cobra.Command, credType, name, canonical string) (*models.NewCloudAccountCredential, error) {
	raw, err := buildCredentialRaw(cmd, credType)
	if err != nil {
		return nil, err
	}

	description, err := cyargs.GetDescription(cmd)
	if err != nil {
		return nil, err
	}

	path, _ := cyargs.GetCredentialPath(cmd)
	credential := &models.NewCloudAccountCredential{
		Name:        ptr.Ptr(name),
		Type:        ptr.Ptr(credType),
		Raw:         raw,
		Description: description,
	}
	if canonical != "" {
		credential.Canonical = canonical
	}
	if path != "" {
		credential.Path = path
	}
	return credential, nil
}

func buildCredentialRaw(cmd *cobra.Command, credType string) (*models.CredentialRaw, error) {
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
			ClientID: clientID, ClientSecret: clientSecret,
			SubscriptionID: subscriptionID, TenantID: tenantID,
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
			Username: username, Password: password, CaCert: string(caCert),
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
			Username: username, Password: password, TenantID: tenantID,
			AuthURL: authURL, DomainID: domainID,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported credential type %q", credType)
	}
}

func addCredentialFlags(cmd *cobra.Command) {
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

func cloudAccountName(cmd *cobra.Command) (string, error) {
	name, err := cyargs.GetName(cmd)
	if err != nil {
		return "", err
	}
	if name == "" {
		canonical, err := cyargs.GetCloudAccount(cmd)
		if err != nil {
			return "", err
		}
		name = canonical
	}
	return name, nil
}

func ptrString(value string) *string {
	if value == "" {
		return nil
	}
	return ptr.Ptr(value)
}
