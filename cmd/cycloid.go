package main

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/creds"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/external-backends"
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
		root.NewLoginCmd(),
		root.NewStatusCmd(),
		root.NewVersionCmd(),
		root.NewDeprecatedExampleCmd(),

		// organizations.NewCommands(),
		// catalogRepositories.NewCommands(),
		// configRepositories.NewCommands(),
		creds.NewCommands(),
		externalBackends.NewCommands(),
		// environments.NewCommands(),
		// events.NewCommands(),
		// pipelines.NewCommands(),
		// projects.NewCommands(),
		// stacks.NewCommands(),
	)
}
