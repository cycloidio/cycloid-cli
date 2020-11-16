package login

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/config"
)

// NewCommands returns the cobra command holding
// the login command / subcommands
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login against the Cycloid console",
		Example: `,
	# Login into my-org using email / password as flags
	cy login --org my-org --email my-email --password my-password

	# Login in a org child of an organization
	cy login --org my-org --child child-org  --email my-email --password my-password

	# Login in an org using an API token
	cy login --org my-org --api-key eyJhbGciOiJI...

	# Login without organization (can be used to access endpoint without organization)
	cy login --email my-email --password my-password

`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE: func(cmd *cobra.Command, args []string) error {

			org, err := cmd.Flags().GetString("org")
			if err != nil {
				return errors.Wrap(err, "unable to get org flag")
			}
			email, err := cmd.Flags().GetString("email")
			if err != nil {
				return errors.Wrap(err, "unable to get email flag")
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				return errors.Wrap(err, "unable to get password flag")
			}
			child, err := cmd.Flags().GetString("child")
			if err != nil {
				return errors.Wrap(err, "unable to get child flag")
			}
			apiKey, err := cmd.Flags().GetString("api-key")
			if err != nil {
				return errors.Wrap(err, "unable to get child flag")
			}

			return login(org, child, email, password, apiKey)
		},
	}

	common.RequiredPersistentFlag(common.WithFlagOrg, cmd)

	WithFlagEmail(cmd)
	WithFlagPassword(cmd)
	WithFlagChild(cmd)
	WithFlagAPIKey(cmd)

	cmd.AddCommand(
		NewListCommand(),
	)

	return cmd
}

func login(org, child, email, password, key string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// if the email and the password are set, it means
	// we need to use the user login
	if len(email) != 0 && len(password) != 0 {
		session, err := m.Login(email, password)
		if err != nil {
			return errors.Wrapf(err, "unable to log user: %s", email)
		}

		// fetch any existing config
		// we skip the error in case it's the first usage and the config
		// file does not exist
		conf, _ := config.ReadConfig()

		// we save the new generated token into the config file
		conf.Token = *session.Token
		if err := config.WriteConfig(conf); err != nil {
			return errors.Wrap(err, "unable to save config")
		}
	}

	if len(org) != 0 {

		// if the API key is not provided, we fallback on the
		// default org authentication using user the token
		if len(key) == 0 {
			orgSession, err := m.LoginOrg(org, child)
			if err != nil {
				return errors.Wrapf(err, "unable to log user: %s", email)
			}
			key = *orgSession.Token
		}
		// we save the new generated token and remove the previous one
		conf, _ := config.ReadConfig()

		// there is no distinction between child or "root" org, we just
		// need to save it once
		// if the org and child are given, we save only the child token
		// else we save the org token
		if len(child) != 0 {
			conf.Organizations[child] = config.Organization{
				Token: key,
			}
		} else {
			conf.Organizations[org] = config.Organization{
				Token: key,
			}
		}
		if err := config.WriteConfig(conf); err != nil {
			return errors.Wrap(err, "unable to save config")
		}
	}
	return nil
}
