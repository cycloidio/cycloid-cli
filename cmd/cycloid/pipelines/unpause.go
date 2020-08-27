package pipelines

import (
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"

	"github.com/spf13/cobra"
)

func NewUnpauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "unpause",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  unpause,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func unpause(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
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

	err = m.UnpausePipeline(org, project, env)

	return err
}
