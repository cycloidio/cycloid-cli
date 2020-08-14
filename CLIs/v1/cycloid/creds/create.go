package creds

import (
	"errors"
	"fmt"
	"io/ioutil"

	"regexp"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/spf13/cobra"

	strfmt "github.com/go-openapi/strfmt"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [type]",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
	}

	// - aws
	// - azure
	// - azure_storage
	// - gcp
	// - elasticsearch
	// - swift

	WithPersistentFlagDescription(cmd)
	common.RequiredPersistentFlag(WithPersistentFlagName, cmd)
	WithPersistentFlagPath(cmd)

	var ssh = &cobra.Command{
		Use:  "ssh",
		RunE: create,
	}
	common.RequiredFlag(WithFlagSSHKey, ssh)

	var basicAuth = &cobra.Command{
		Use:  "basic_auth",
		RunE: create,
	}
	common.RequiredFlag(WithFlagUsername, basicAuth)
	common.RequiredFlag(WithFlagPassword, basicAuth)

	var custom = &cobra.Command{
		Use:  "custom",
		RunE: create,
	}
	common.RequiredFlag(WithFlagField, custom)

	cmd.AddCommand(custom, basicAuth, ssh)

	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	var err error
	var body *models.CreateCredential
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

	credP := organization_credentials.NewCreateCredentialParams()
	credP.SetOrganizationCanonical(org)

	// rawCred = CredentialRaw{
	// 	AccessKey string `json:"access_key,omitempty"`
	// 	AccountName string `json:"account_name,omitempty"`
	// 	AuthURL string `json:"auth_url,omitempty"`
	// 	CaCert string `json:"ca_cert,omitempty"`
	// 	ClientID string `json:"client_id,omitempty"`
	// 	ClientSecret string `json:"client_secret,omitempty"`
	// 	DomainID string `json:"domain_id,omitempty"`
	// 	JSONKey string `json:"json_key,omitempty"`
	// 	SecretKey string `json:"secret_key,omitempty"`
	// 	SSHKey string `json:"ssh_key,omitempty"`
	// 	SubscriptionID string `json:"subscription_id,omitempty"`
	// 	TenantID string `json:"tenant_id,omitempty"`

	// }

	if credT == "ssh" {
		sshKeyPath, err := cmd.Flags().GetString("ssh-key")
		if err != nil {
			return err
		}

		sshKey, err := ioutil.ReadFile(sshKeyPath)
		if err != nil {
			return errors.New("File reading error")

		}

		rawCred = &models.CredentialRaw{
			SSHKey: string(sshKey),
		}

	} else if credT == "basic_auth" {
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
	} else if credT == "custom" {
		fields, err := cmd.Flags().GetStringToString("field")
		if err != nil {
			return err
		}
		rawCred = &models.CredentialRaw{
			Raw: fields,
		}
	} else {
		return errors.New("Unexpected type")
	}

	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		safePath := re.ReplaceAllString(name, "-")
		path = fmt.Sprintf("%s_%s", credT, safePath)
	}
	fmt.Println(path)

	body = &models.CreateCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &credT,
	}

	credP.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.OrganizationCredentials.CreateCredential(credP, root.ClientCredentials())
	// TODO create a error handeling function to format our error with a better display
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}

// /organizations/{organization_canonical}/credentials
// post: createCredential
// Create a new Credential, based on the type you will have to pass different parameters within the body:
// * ssh: ssh_key
// * aws: access_key, secret_key
// * gcp: json_key
// * azure: client_id, client_secret, subscription_id, tenant_id
// * azure_storage: account_name, access_key
// * basic_auth: username, password
// * elasticsearch: username, password, ca_cert
// * swift: auth_url, username, password, domain_id, tenant_id
