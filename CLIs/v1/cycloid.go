package main

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/catalog-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/config-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/creds"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/environments"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/events"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/external-backends"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/organizations"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/pipelines"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/projects"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/stacks"
	"github.com/spf13/cobra"
)

func AttachCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		// Root
		root.NewLoginCmd(),
		root.NewStatusCmd(),
		root.NewVersionCmd(),
		root.NewDeprecatedExampleCmd(),

		organizations.NewCommands(),
		catalogRepositories.NewCommands(),
		configRepositories.NewCommands(),
		creds.NewCommands(),
		externalBackends.NewCommands(),
		environments.NewCommands(),
		events.NewCommands(),
		pipelines.NewCommands(),
		projects.NewCommands(),
		stacks.NewCommands(),
	)
}
