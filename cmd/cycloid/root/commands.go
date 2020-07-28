package root

import (
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/catalog-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/config-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/creds"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/environments"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/events"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/organizations"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/pipelines"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/projects"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/root/stacks"
	"github.com/spf13/cobra"
)

func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		// Root
		NewLoginCmd(),
		NewStatusCmd(),
		NewVersionCmd(),
		NewDeprecatedExampleCmd(),

		organizations.NewCommands(),
		catalogRepositories.NewCommands(),
		configRepositories.NewCommands(),
		creds.NewCommands(),
		environments.NewCommands(),
		events.NewCommands(),
		pipelines.NewCommands(),
		projects.NewCommands(),
		stacks.NewCommands(),
	)
}
