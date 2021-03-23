package login

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/config"
)

// NewCommands returns the cobra command holding
// the login command / subcommands
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login against the Cycloid console",
		Example: `,
	# Login in an org using api-key
	cy login --org my-org --api-key eyJhbGciOiJI...
`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE: func(cmd *cobra.Command, args []string) error {

			org, err := cmd.Flags().GetString("org")
			if err != nil {
				return errors.Wrap(err, "unable to get org flag")
			}
			apiKey, err := cmd.Flags().GetString("api-key")
			if err != nil {
				return errors.Wrap(err, "unable to get child flag")
			}

			return login(org, apiKey)
		},
	}

	common.RequiredFlag(WithFlagAPIKey, cmd)
	common.RequiredFlag(WithFlagOrg, cmd)

	cmd.AddCommand(
		NewListCommand(),
	)

	return cmd
}

func login(org, key string) error {

	// we save the new token and remove the previous one
	conf, _ := config.ReadConfig()

	conf.Organizations[org] = config.Organization{
		Token: key,
	}

	if err := config.WriteConfig(conf); err != nil {
		return errors.Wrap(err, "unable to save config")
	}

	return nil
}
