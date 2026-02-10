package cmd

import (
	"fmt"
	"strings"

	root "github.com/cycloidio/cycloid-cli/cmd/cycloid"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/apikey"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/beta"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/catalogrepositories"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/components"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/configrepositories"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/credentials"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/environments"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/events"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/externalbackends"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/infrapolicies"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/kpis"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/login"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/members"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/organizations"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/projects"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/roles"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/stacks"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/teams"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/terracost"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/uri"
	"github.com/cycloidio/cycloid-cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Args:          cobra.NoArgs,
		Use:           "cy",
		Short:         "Cycloid CLI",
		Long: `--- CLI tool to interact with Cycloid API. ---
Documentation at https://docs.cycloid.io/reference/cli/

-- Environment variables --
Some environment variables can be set to ease context setting in Cycloid.
Those variables will be overridden by related flags.

Name         |  Description
-------------|-----------------
CY_API_URL   | Specify the HTTP url of Cycloid API to use, default https://http-api.cycloid.io
CY_ORG       | Set the current organization
CY_PROJECT   | Set the current project
CY_ENV       | Set the current environment
CY_COMPONENT | Set the current component
CY_API_KEY   | Set the current API Key to use
CY_VERBOSITY | Set the verbosity level (debug, info, warning, error), default warning.
HTTP_PROXY   | Set the http proxy with host[:port] format for http request
HTTPS_PROXY  | Set the https proxy with host[:port] format for https request
NO_PROXY     | List of hosts that must bypass proxy configuration
`,
	}

	rootCmd.PersistentFlags().StringVarP(&userOutput, "output", "o", "table", "The formatting style for command output: json|yaml|table")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	rootCmd.RegisterFlagCompletionFunc("output", func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		return []cobra.Completion{"json", "table", "yaml"}, cobra.ShellCompDirectiveDefault
	})

	rootCmd.PersistentFlags().StringP("verbosity", "v", "warning", "Override the default verbosity for this command. VERBOSITY must be one of: debug, info, warning, error, critical, none.")
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	viper.SetDefault("verbosity", "warning")

	rootCmd.PersistentFlags().String("api-url", "https://http-api.cycloid.io", "Specify the HTTP url of Cycloid API to use eg https://http-api.cycloid.io. This can also be given by CY_API_URL environment variable.")
	viper.BindPFlag("api-url", rootCmd.PersistentFlags().Lookup("api-url"))

	rootCmd.PersistentFlags().Bool("insecure", false, "Decide to skip or not TLS verification")
	viper.BindPFlag("insecure", rootCmd.PersistentFlags().Lookup("insecure"))

	rootCmd.PersistentFlags().String("org", "", "Specify the org to use. override CY_ORG env var. Required for all Org scoped endpoint.")
	viper.BindPFlag("org", rootCmd.PersistentFlags().Lookup("org"))

	// Remove usage on error, this is annoying in scripting
	rootCmd.SilenceUsage = true

	// Disable file completion fallback by default
	rootCmd.CompletionOptions.SetDefaultShellCompDirective(cobra.ShellCompDirectiveNoFileComp)

	AttachCommands(rootCmd)

	return rootCmd
}

func AttachCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		root.NewVersionCmd(),
		root.NewStatusCmd(),
		root.NewCompletionCmd(),
		root.NewGetCommand(),
		apikey.NewCommands(),
		catalogrepositories.NewCommands(),
		configrepositories.NewCommands(),
		credentials.NewCommands(),
		events.NewCommands(),
		externalbackends.NewCommands(),
		infrapolicies.NewCommands(),
		members.NewCommands(),
		organizations.NewCommands(),
		pipelines.NewCommands(),
		projects.NewCommands(),
		environments.NewCommands(),
		components.NewCommands(),
		kpis.NewCommands(),
		roles.NewCommands(),
		stacks.NewCommands(),
		login.NewCommands(),
		terracost.NewCommands(),
		beta.NewCommands(),
		uri.NewURICommands(),
		teams.NewTeamsCommands(),
	)
}
