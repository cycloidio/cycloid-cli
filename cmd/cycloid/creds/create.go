package creds

import (
	"errors"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [ssh|basic_auth|custom]",
		Short: "create a credential",
		Example: `
	# create a credential for basic authentication
	cy --my-org credential create basic_auth --username my-username --password my-password

	# create a credential for SSH
	cy --my-org credential create ssh --ssh-key /path/to/private/key
`,
	}

	WithPersistentFlagDescription(cmd)
	common.RequiredPersistentFlag(WithPersistentFlagName, cmd)
	WithPersistentFlagPath(cmd)

	var ssh = &cobra.Command{
		Use:  "ssh",
		RunE: create,
		Example: `
	# create a credential for SSH
	cy --my-org credential create ssh --ssh-key /path/to/private/key
`,
	}
	common.RequiredFlag(WithFlagSSHKey, ssh)

	var basicAuth = &cobra.Command{
		Use:  "basic_auth",
		RunE: create,
		Example: `
	# create a credential for basic authentication
	cy --my-org credential create basic_auth --username my-username --password my-password
`,
	}
	common.RequiredFlag(WithFlagUsername, basicAuth)
	common.RequiredFlag(WithFlagPassword, basicAuth)

	var custom = &cobra.Command{
		Use:  "custom",
		RunE: create,
		Example: `
	# create a credential for custom type
	cy --my-org credential create custom --my-key=my-value
`,
	}
	common.RequiredFlag(WithFlagField, custom)

	cmd.AddCommand(custom, basicAuth, ssh)

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

	err = m.CreateCredential(org, name, credT, rawCred, path, description)

	return err
}
