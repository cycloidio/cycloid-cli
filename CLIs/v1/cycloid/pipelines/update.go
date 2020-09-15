package pipelines

import (
	"fmt"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

func NewUpdateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "update",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  update,
	}
	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)
	WithFlagConfig(cmd)

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	var err error

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
	varsPath, err := cmd.Flags().GetString("vars")
	if err != nil {
		return err
	}
	pipelinePath, err := cmd.Flags().GetString("pipeline")
	if err != nil {
		return err
	}
	configs, err := cmd.Flags().GetStringToString("config")
	if err != nil {
		return err
	}

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
	}
	pipeline := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
	}
	variables := string(rawVars)

	resp, err := m.UpdatePipeline(org, project, env, pipeline, variables)
	if err != nil {
		return err
	}
	fmt.Println(resp)

	//
	// PUSH CONFIG If pipeline update succeeded
	//

	if len(configs) > 0 {

		cfs := make(map[string]strfmt.Base64)

		for fp, dest := range configs {
			var c strfmt.Base64
			c, err = ioutil.ReadFile(fp)
			if err != nil {
				return fmt.Errorf("Config file reading error : %s", err.Error())
			}
			cfs[dest] = c
		}

		err = m.PushConfig(org, project, env, cfs)
		if err != nil {
			return err
		}
	}

	return nil
}
