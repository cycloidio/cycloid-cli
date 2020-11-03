package pipelines

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
)

func NewUnpauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "unpause",
		Short: "unpause a pipeline",
		Example: `
	# unpause pipeline my-project-env
	cy --org my-org pipeline unpause --project my-project --env env
`,
		RunE:    unpause,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func unpause(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	project, err := cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	if err := m.UnpausePipeline(org, project, env); err != nil {
		return errors.Wrap(err, "unable to unpause pipeline")
	}

	return nil
}
