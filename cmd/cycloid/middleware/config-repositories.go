package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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
	_, err = m.api.OrganizationConfigRepositories.CreateConfigRepositoryConfig(params, common.ClientCredentials(&org))
	if err != nil {
		return err
	}

	return nil
}

func (m *middleware) ListConfigRepositories(org string) ([]*models.ConfigRepository, error) {

	params := organization_config_repositories.NewGetConfigRepositoriesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationConfigRepositories.GetConfigRepositories(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}

func (m *middleware) GetConfigRepository(org string, configRepo uint32) (*models.ConfigRepository, error) {

	params := organization_config_repositories.NewGetConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryID(configRepo)

	resp, err := m.api.OrganizationConfigRepositories.GetConfigRepository(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}

func (m *middleware) DeleteConfigRepository(org string, configRepo uint32) error {
	params := organization_config_repositories.NewDeleteConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryID(configRepo)

	_, err := m.api.OrganizationConfigRepositories.DeleteConfigRepository(params, common.ClientCredentials(&org))
	return err
}

func (m *middleware) CreateConfigRepository(org, name, url, branch string, setDefault bool, cred uint32) (*models.ConfigRepository, error) {

	params := organization_config_repositories.NewCreateConfigRepositoryParams()
	params.SetOrganizationCanonical(org)

	body := &models.CreateConfigRepository{
		Branch:       &branch,
		CredentialID: &cred,
		Default:      &setDefault,
		Name:         &name,
		URL:          &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationConfigRepositories.CreateConfigRepository(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}

func (m *middleware) UpdateConfigRepository(org string, configRepo uint32, name, url, branch string, setDefault bool, cred uint32) (*models.ConfigRepository, error) {

	params := organization_config_repositories.NewUpdateConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryID(configRepo)

	body := &models.UpdateConfigRepository{
		Branch:       &branch,
		CredentialID: &cred,
		Default:      &setDefault,
		Name:         &name,
		URL:          &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationConfigRepositories.UpdateConfigRepository(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	// TODO this validate have been removed https://github.com/cycloidio/youdeploy-http-api/issues/2262
	// err = p.Validate(strfmt.Default)
	// if err != nil {
	// 	return err
	// }

	d := p.Data

	return d, err
}
