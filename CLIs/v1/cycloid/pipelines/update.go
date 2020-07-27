package pipelines

import (
	"fmt"
	"io/ioutil"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_pipelines"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
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
	api := root.NewAPI()
	m := middleware.NewMiddleware(api)

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
	configs, err := cmd.Flags().GetStringToString("config")
	if err != nil {
		return err
	}

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	pipelineName := fmt.Sprintf("%s-%s", project, env)

	params := organization_pipelines.NewUpdatePipelineParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetInpathPipelineName(pipelineName)

	rawPipeline, err := ioutil.ReadFile(pipelinePath)
	if err != nil {
		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
	}
	pipelineTemplate := string(rawPipeline)

	rawVars, err := ioutil.ReadFile(varsPath)
	if err != nil {
		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
	}
	vars := common.ReplaceCycloidVarsString(cyCtx, string(rawVars))

	body = &models.UpdatePipeline{
		PassedConfig: &pipelineTemplate,
		YamlVars:     vars,
	}
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	params.SetBody(body)
	resp, err := api.OrganizationPipelines.UpdatePipeline(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	//
	// PUSH CONFIG If pipeline update succeeded
	//

	if len(configs) > 0 {
		projectData, err := m.GetProject(org, project)
		if err != nil {
			return err
		}

		paramsC := organization_config_repositories.NewCreateConfigRepositoryConfigParams()
		paramsC.SetOrganizationCanonical(org)
		paramsC.SetConfigRepositoryID(projectData.ConfigRepositoryID)

		var cfs []*models.ConfigFile

		for fp, dest := range configs {
			var c strfmt.Base64
			p := common.ReplaceCycloidVarsString(cyCtx, dest)
			c, err = ioutil.ReadFile(fp)
			if err != nil {
				return fmt.Errorf("Config file reading error : %s", err.Error())
			}
			c = common.ReplaceCycloidVars(cyCtx, c)

			cf := &models.ConfigFile{
				Content: &c,
				Path:    &p,
			}
			err = cf.Validate(strfmt.Default)
			if err != nil {
				return err
			}

			cfs = append(cfs, cf)
		}

		bodyC := &models.SCConfig{Configs: cfs}

		err = bodyC.Validate(strfmt.Default)
		if err != nil {
			return err
		}

		paramsC.SetBody(bodyC)
		_, err = api.OrganizationConfigRepositories.CreateConfigRepositoryConfig(paramsC, root.ClientCredentials())
		if err != nil {
			return err
		}
	}

	fmt.Println(resp)

	return nil
}
