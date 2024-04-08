package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_external_backends"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListExternalBackends(org string) ([]*models.ExternalBackend, error) {

	params := organization_external_backends.NewGetExternalBackendsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationExternalBackends.GetExternalBackends(params, m.api.Credentials(&org))
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

func (m *middleware) DeleteExternalBackend(org string, externalBackend uint32) error {

	params := organization_external_backends.NewDeleteExternalBackendParams()
	params.SetOrganizationCanonical(org)
	params.SetExternalBackendID(externalBackend)

	_, err := m.api.OrganizationExternalBackends.DeleteExternalBackend(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) CreateExternalBackends(org, project, env, purpose, cred string, def bool, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error) {

	params := organization_external_backends.NewCreateExternalBackendParams()
	params.SetOrganizationCanonical(org)

	var body *models.NewExternalBackend

	if len(cred) != 0 {
		body = &models.NewExternalBackend{
			ProjectCanonical:     project,
			Purpose:              &purpose,
			EnvironmentCanonical: env,
			CredentialCanonical:  cred,
			Default:              def,
		}
	} else {
		body = &models.NewExternalBackend{
			ProjectCanonical:     project,
			EnvironmentCanonical: env,
			Purpose:              &purpose,
			Default:              def,
		}
	}

	err := ebConfig.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	body.SetConfiguration(ebConfig)
	params.SetBody(body)
	if project != "" {
		params.SetProject(&project)
	}
	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationExternalBackends.CreateExternalBackend(params, m.api.Credentials(&org))
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
