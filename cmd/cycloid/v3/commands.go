package root

import (
	v1 "github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/catalog-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/config-repositories"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/creds"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/environments"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/events"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/organizations"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/pipelines"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/projects"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v1/stacks"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/v3/external-backends"
	"github.com/spf13/cobra"
)

func AttachCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		// Root
		v1.NewLoginCmd(),
		v1.NewStatusCmd(),
		v1.NewVersionCmd(),
		v1.NewDeprecatedExampleCmd(),

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
