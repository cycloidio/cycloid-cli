package pipelines

import (
	"fmt"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"

	"github.com/spf13/cobra"
)

func NewPauseCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pause",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  pause,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)

	return cmd
}

func pause(cmd *cobra.Command, args []string) error {
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

	err = m.PausePipeline(org, project, env)
	if err != nil {
		return err
	}

	// fmt.Println(resp)
	fmt.Printf("%+v\n", err)
	return nil
}
