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
	conf, err := config.Read()
	if err != nil {
		return errors.Wrap(err, "unable to read config")
	}

	// Check for a nil map.
	// This can be the case if the config file is empty
	if conf.Organizations == nil {
		conf.Organizations = make(map[string]config.Organization)
	}

	conf.Organizations[org] = config.Organization{
		Token: key,
	}

	if err := config.Write(conf); err != nil {
		return errors.Wrap(err, "unable to save config")
	}

	return nil
}
