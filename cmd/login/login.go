package login

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/internal"
	"github.com/cycloidio/cycloid-cli/config"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

// NewCommands returns the cobra command holding
// the login command / subcommands
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Args:  cobra.NoArgs,
		Short: "Login against the Cycloid console",
		Example: `# Login in an org using api-key
export CY_API_KEY=xxxx
cy login --org my-org`,
		RunE: login,
	}

	cmd.Flags().String("api-key", "", "[deprecated] set the API key, use CY_API_KEY env var instead.")
	viper.BindPFlag("api-key", cmd.Flags().Lookup("api-key"))

	cmd.AddCommand(
		NewListCommand(),
	)

	return cmd
}

func login(cmd *cobra.Command, args []string) error {
	conf, _ := config.Read()
	// If err != nil, the file does not exist, we create it anyway

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return fmt.Errorf("unable to get org flag: %w", err)
	}

	// Get api key via env var or cli flag
	apiKey := viper.GetString("api-key")
	if apiKey == "" {
		return cyout.PrintWithOptions(cmd, nil, nil, "CY_API_KEY is not set or invalid", printer.Options{})
	}

	// Warn user about deprecation
	if cmd.Flags().Lookup("api-key").Changed {
		internal.Warning(cmd.ErrOrStderr(), "--api-key is deprecated, use CY_API_KEY env var instead")
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
		return cyout.PrintWithOptions(cmd, nil, err, "unable to write config file", printer.Options{})
	}

	return nil
}
