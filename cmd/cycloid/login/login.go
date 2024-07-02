package login

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
		RunE:    login,
	}

	WithFlagAPIKey(cmd)

	cmd.AddCommand(
		NewListCommand(),
	)

	return cmd
}

func login(cmd *cobra.Command, args []string) error {

	conf, _ := config.Read()
	// If err != nil, the file does not exist, we create it anyway

	org, err := common.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get org flag")
	}

	apiKey := viper.GetString("api-key")
	if apiKey == "" {
		return errors.Wrap(err, "apiKey not set or invalid")
	}

	// Check for a nil map.
	// This can be the case if the config file is empty
	if conf.Organizations == nil {
		conf.Organizations = make(map[string]config.Organization)
	}

	conf.Organizations[org] = config.Organization{
		Token: apiKey,
	}

	if err := config.Write(conf); err != nil {
		return errors.Wrap(err, "unable to save config")
	}

	return nil
}
