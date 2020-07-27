package projects

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/youdeploy-cli/client/client/organization_projects"

	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "...",
		Long:  `........ . . .... .. .. ....`,
		RunE:  create,
	}
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagStackRef, cmd)
	common.RequiredFlag(WithFlagConfigRepository, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)

	WithFlagUsecase(cmd)
	WithFlagDescription(cmd)
	WithFlagCanonical(cmd)
	WithFlagCloudProvider(cmd)
	WithFlagConfig(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := root.NewAPI()

	var err error
	var body *models.NewProject
	var pipelines []*models.NewPipeline

	org, err := cmd.Flags().GetString("org")
	if err != nil {
		return err
	}
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}
	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}
	cloudProvider, err := cmd.Flags().GetString("cloud-provider")
	if err != nil {
		return err
	}
	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return err
	}
	configRepo, err := cmd.Flags().GetUint32("config-repo")
	if err != nil {
		return err
	}
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}
	usecase, err := cmd.Flags().GetString("usecase")
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

	if canonical == "" {
		re := regexp.MustCompile(`[^a-z0-9\-_]`)
		canonical = re.ReplaceAllString(strings.ToLower(name), "-")
	}

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: canonical}

	params := organization_projects.NewCreateProjectParams()
	params.SetOrganizationCanonical(org)

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

	pipelineName := fmt.Sprintf("%s-%s", canonical, env)

	pipeline := &models.NewPipeline{
		Environment:  &env,
		PipelineName: &pipelineName,
		UseCase:      usecase,
		PassedConfig: &pipelineTemplate,
		YamlVars:     vars,
	}
	err = pipeline.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	pipelines = append(pipelines, pipeline)

	body = &models.NewProject{
		Name:               &name,
		Description:        description,
		Canonical:          &canonical,
		CloudProvider:      cloudProvider,
		ServiceCatalogRef:  &stackRef,
		ConfigRepositoryID: &configRepo,
		Pipelines:          pipelines,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	params.SetBody(body)
	resp, err := api.OrganizationProjects.CreateProject(params, root.ClientCredentials())
	// TODO create a error handeling function to format our error with a better display
	if err != nil {
		return err
	}

	// If project creation succeeded we push the config files
	paramsC := organization_config_repositories.NewCreateConfigRepositoryConfigParams()
	paramsC.SetOrganizationCanonical(org)
	paramsC.SetConfigRepositoryID(configRepo)

	if len(configs) > 0 {
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

// /organizations/{organization_canonical}/projects
// post: createProject
// Create a new project with envs and pipelines in the organization.
