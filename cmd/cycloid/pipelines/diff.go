package pipelines

import (
	"fmt"
	"io/ioutil"

	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

func NewDiffCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "diff",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  diff,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)

	return cmd
}

func diff(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()
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

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
	}
	vars := string(rawVars)

	pd, err := m.DiffPipeline(org, project, env, pipelineTemplate, vars)
	if err != nil {
		return err
	}

	fmt.Printf("Groups: %s\n", pd.Groups)

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", pd)
	return nil
}

// /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff
// put: diffPipeline
// The diff between the provided pipeline configuration and the pipeline from the given name.
