package login

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/config"
	"github.com/cycloidio/cycloid-cli/internal/cy_args"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

// NewCommands returns the cobra command holding
// the login command / subcommands
func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login against the Cycloid console",
		Example: `# Login in an org using api-key
export CY_API_KEY=xxxx
cy login --org my-org`,
		PreRunE: internal.CheckAPIAndCLIVersion,
		RunE:    login,
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

	org, err := cy_args.GetOrg(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get org flag")
	}

	p, err := factory.GetPrinter(viper.GetString("output"))
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Get api key via env var or cli flag
	apiKey := viper.GetString("api-key")
	if apiKey == "" {
		return printer.SmartPrint(p, nil, nil, "CY_API_KEY is not set or invalid", printer.Options{}, cmd.OutOrStderr())
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
		return printer.SmartPrint(p, nil, err, "unable to write config file", printer.Options{}, cmd.OutOrStderr())
	}

	return nil
}
