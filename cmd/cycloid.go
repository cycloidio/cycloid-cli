package main

import (
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	catalogRepositories "github.com/cycloidio/youdeploy-cli/cmd/cycloid/catalog-repositories"
	configRepositories "github.com/cycloidio/youdeploy-cli/cmd/cycloid/config-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/creds"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/events"
	externalBackends "github.com/cycloidio/youdeploy-cli/cmd/cycloid/external-backends"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/members"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/organizations"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/projects"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/stacks"
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
		root.LoginCmd,
		root.NewStatusCmd(),
		root.NewVersionCmd(),
		root.NewDeprecatedExampleCmd(),

		organizations.NewCommands(),
		catalogRepositories.NewCommands(),
		configRepositories.NewCommands(),
		creds.NewCommands(),
		externalBackends.NewCommands(),
		events.NewCommands(),
		pipelines.NewCommands(),
		projects.NewCommands(),
		stacks.NewCommands(),
		members.NewCommands(),
	)
}
