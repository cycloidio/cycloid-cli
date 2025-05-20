package middleware

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_config_repositories"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListConfigRepositories(org string) ([]*models.ConfigRepository, error) {
	params := organization_config_repositories.NewListConfigRepositoriesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationConfigRepositories.ListConfigRepositories(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) GetConfigRepository(org, configRepo string) (*models.ConfigRepository, error) {
	params := organization_config_repositories.NewGetConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryCanonical(configRepo)

	resp, err := m.api.OrganizationConfigRepositories.GetConfigRepository(params, m.api.Credentials(&org))

	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}

func (m *middleware) DeleteConfigRepository(org, configRepo string) error {
	params := organization_config_repositories.NewDeleteConfigRepositoryParams()
	params.SetOrganizationCanonical(org)
	params.SetConfigRepositoryCanonical(configRepo)

	_, err := m.api.OrganizationConfigRepositories.DeleteConfigRepository(params, m.api.Credentials(&org))
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

	resp, err := m.api.OrganizationConfigRepositories.CreateConfigRepository(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
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

	resp, err := m.api.OrganizationConfigRepositories.UpdateConfigRepository(params, m.api.Credentials(&org))

	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	return payload.Data, nil
}
