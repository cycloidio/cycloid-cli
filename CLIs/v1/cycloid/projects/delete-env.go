package projects

import (
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"

	"github.com/spf13/cobra"
)

func NewDeleteEnvCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "delete-env",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  deleteEnv,
	}
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)

	return cmd
}

func deleteEnv(cmd *cobra.Command, args []string) error {
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

	err = m.DeleteProjectEnv(org, project, env)

	return err
}
