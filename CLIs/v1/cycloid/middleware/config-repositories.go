package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) PushConfig(org string, project string, env string, configs map[string]strfmt.Base64) error {

	cyCtx := common.CycloidContext{Env: env,
		Org:     org,
		Project: project}

	projectData, err := m.GetProject(org, project)
	if err != nil {
		return err
	}

	params := organization_config_repositories.NewCreateConfigRepositoryConfigParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryID(projectData.ConfigRepositoryID)

	var cfs []*models.ConfigFile

	for rawP, rawC := range configs {
		p := common.ReplaceCycloidVarsString(cyCtx, rawP)
		var c strfmt.Base64
		c = common.ReplaceCycloidVars(cyCtx, rawC)

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

	body := &models.SCConfig{Configs: cfs}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	params.SetBody(body)
	_, err = m.api.OrganizationConfigRepositories.CreateConfigRepositoryConfig(params, root.ClientCredentials())
	if err != nil {
		return err
	}

	return nil
}
