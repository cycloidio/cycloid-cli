package pipelines

import (
	"fmt"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
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

	var err error
	var body *models.UpdatePipeline

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

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	vars := common.ReplaceCycloidVarsString(cyCtx, string(rawVars))

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines.NewDiffPipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetInpathPipelineName(pipelineName)

	body = &models.UpdatePipeline{
		PassedConfig: &pipelineTemplate,
		YamlVars:     vars,
	}

	params.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	resp, err := api.OrganizationPipelines.DiffPipeline(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	pd := p.Data

	fmt.Printf("Groups: %s\n", pd.Groups)

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", resp)
	return nil
}

// /organizations/{organization_canonical}/pipelines/{inpath_pipeline_name}/diff
// put: diffPipeline
// The diff between the provided pipeline configuration and the pipeline from the given name.
