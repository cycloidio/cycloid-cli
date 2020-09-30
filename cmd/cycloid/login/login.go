package login

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/internal"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/youdeploy-cli/config"
)

var (
	example = `
	# Login into my-org using email / password as flags
	cy login --org my-org --email my-email --password my-password

	# Login without organization (can be used to access endpoint without organization)
	cy login --email my-email --password my-password
`
	short    = "Login against the Cycloid console"
	long     = short
	org      string
	email    string
	password string
	LoginCmd = &cobra.Command{
		Use:     "login",
		Short:   short,
		Long:    long,
		Example: example,
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

			api := common.NewAPI()
			m := middleware.NewMiddleware(api)

			// we first need to authenticate the user against cycloid console
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

			if len(org) != 0 {
				orgSession, err := m.LoginOrg(org, email, password)
				if err != nil {
					return errors.Wrapf(err, "unable to log user: %s", email)
				}

				// we save the new generated token and remove the previous one
				conf, err := config.ReadConfig()
				if err != nil {
					return errors.Wrap(err, "unable to read config: %s")
				}
				conf.Organizations[org] = config.Organization{
					Token: *orgSession.Token,
				}
				if err := config.WriteConfig(conf); err != nil {
					return errors.Wrap(err, "unable to save config")
				}
			}
			return nil
		},
	}
)

func init() {
	LoginCmd.PersistentFlags().StringVar(&org, "org", "", "organization")
	LoginCmd.PersistentFlags().StringVar(&email, "email", "", "email")
	LoginCmd.PersistentFlags().StringVar(&password, "password", "", "password")

	LoginCmd.AddCommand(list)
}
