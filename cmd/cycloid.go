package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	root "github.com/cycloidio/cycloid-cli/cmd/cycloid"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/apikey"
	catalogRepositories "github.com/cycloidio/cycloid-cli/cmd/cycloid/catalog-repositories"
	configRepositories "github.com/cycloidio/cycloid-cli/cmd/cycloid/config-repositories"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/creds"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/events"
	externalBackends "github.com/cycloidio/cycloid-cli/cmd/cycloid/external-backends"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/infrapolicies"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/kpis"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/login"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/members"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/organizations"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/projects"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/roles"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/stacks"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/terracost"
	"github.com/cycloidio/cycloid-cli/internal/version"
)

var (
	versionString = fmt.Sprintf("%s, revision %s, branch %s, date %s; go %s", version.Version, version.Revision, version.Branch, version.BuildDate, version.GoVersion)

	// Used for flags.
	userOutput string
)

func init() {
	viper.SetEnvPrefix("CY")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Version:       versionString,
		SilenceErrors: true,
		SilenceUsage:  false,
		Use:           "cy",
		Short:         "Cycloid CLI",
		Long:          `Cy is a CLI for Cycloid framework. Learn more at https://www.cycloid.io/.`,
	}

	rootCmd.PersistentFlags().StringVarP(&userOutput, "output", "o", "table", "The formatting style for command output: json|yaml|table")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	viper.SetDefault("verbosity", "warning")

	rootCmd.PersistentFlags().String("api-url", "https://http-api.cycloid.io", "Specify the HTTP url of Cycloid API to use eg https://http-api.cycloid.io. This can also be given by CY_API_URL environment variable.")
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))

	rootCmd.PersistentFlags().Bool("insecure", false, "Decide to skip or not TLS verification")
	viper.BindPFlag("insecure", rootCmd.PersistentFlags().Lookup("insecure"))

	// Remove usage on error, this is annoying in scripting
	rootCmd.SilenceUsage = true

	AttachCommands(rootCmd)

	return rootCmd
}

func AttachCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		// Root
		root.NewVersionCmd(),
		root.NewStatusCmd(),
		root.NewCompletionCmd(),
		apikey.NewCommands(),
		catalogRepositories.NewCommands(),
		configRepositories.NewCommands(),
		creds.NewCommands(),
		events.NewCommands(),
		externalBackends.NewCommands(),
		infrapolicies.NewCommands(),
		members.NewCommands(),
		organizations.NewCommands(),
		pipelines.NewCommands(),
		projects.NewCommands(),
		kpis.NewCommands(),
		roles.NewCommands(),
		stacks.NewCommands(),
		login.NewCommands(),
		terracost.NewCommands(),
	)
}
