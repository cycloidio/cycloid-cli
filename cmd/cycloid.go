package cmd

import (
	root "github.com/cycloidio/cycloid-cli/cmd/cycloid"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/apikey"
	catalogRepositories "github.com/cycloidio/cycloid-cli/cmd/cycloid/catalog-repositories"
	configRepositories "github.com/cycloidio/cycloid-cli/cmd/cycloid/config-repositories"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/creds"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/events"
	externalBackends "github.com/cycloidio/cycloid-cli/cmd/cycloid/external-backends"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/infrapolicies"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/login"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/members"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/organizations"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/projects"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/roles"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/stacks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("CY")
	// Disabled, we decided to choose which arg will be available as env var
	// viper.AutomaticEnv()

}

func AttachCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		// Root
		root.NewVersionCmd(),
		root.NewValidateFormCmd(),
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
		roles.NewCommands(),
		stacks.NewCommands(),
		login.NewCommands(),
	)
}
