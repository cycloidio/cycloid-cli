package middleware

import (
	"strings"

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
		return NewApiError(err)
	}

	params := organization_config_repositories.NewCreateConfigRepositoryConfigParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryCanonical(projectData.ConfigRepositoryCanonical)

	var cfs []*models.ConfigFile

	for rawP, rawC := range configs {
		p := common.ReplaceCycloidVarsString(cyCtx, rawP)
		// Ensure the path doesn't start with a / as it will not be valid for the API calls
		p = strings.TrimLeft(p, "/")

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
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) ListConfigRepositories(org string) ([]*models.ConfigRepository, error) {

	params := organization_config_repositories.NewListConfigRepositoriesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationConfigRepositories.ListConfigRepositories(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) GetConfigRepository(org, configRepo string) (*models.ConfigRepository, error) {

	params := organization_config_repositories.NewGetConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryCanonical(configRepo)

	resp, err := m.api.OrganizationConfigRepositories.GetConfigRepository(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) DeleteConfigRepository(org, configRepo string) error {
	params := organization_config_repositories.NewDeleteConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryCanonical(configRepo)

	_, err := m.api.OrganizationConfigRepositories.DeleteConfigRepository(params, common.ClientCredentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
}

func (m *middleware) CreateConfigRepository(org, name, url, branch, cred string, setDefault bool) (*models.ConfigRepository, error) {

	params := organization_config_repositories.NewCreateConfigRepositoryParams()
	params.SetOrganizationCanonical(org)

	body := &models.NewConfigRepository{
		Branch:              &branch,
		CredentialCanonical: &cred,
		Default:             &setDefault,
		Name:                &name,
		URL:                 &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationConfigRepositories.CreateConfigRepository(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

func (m *middleware) UpdateConfigRepository(org, configRepo, cred, name, url, branch string, setDefault bool) (*models.ConfigRepository, error) {

	params := organization_config_repositories.NewUpdateConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryCanonical(configRepo)

	body := &models.UpdateConfigRepository{
		Branch:              &branch,
		CredentialCanonical: &cred,
		Default:             &setDefault,
		Name:                &name,
		URL:                 &url,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationConfigRepositories.UpdateConfigRepository(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}
