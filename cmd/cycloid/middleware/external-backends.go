package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organization_external_backends"
	"github.com/cycloidio/cycloid-cli/client/models"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"errors"
)

const (
	defaultEB = true
)

func (m *middleware) GetRemoteTFExternalBackend(org string) (*models.ExternalBackend, error) {
	params := organization_external_backends.NewGetExternalBackendsParams()
	params.SetOrganizationCanonical(org)
	params.SetExternalBackendDefault(ptr.Ptr(defaultEB))

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
	if len(d) == 0 {
		return nil, NewApiError(errors.New("couldn't find the remote terraform backend"))
	}

	return d[0], nil
}

func (m *middleware) GetExternalBackend(org string, externalBackend uint32) (*models.ExternalBackend, error) {
	params := organization_external_backends.NewGetExternalBackendParams()
	params.SetOrganizationCanonical(org)
	params.SetExternalBackendID(externalBackend)

	resp, err := m.api.OrganizationExternalBackends.GetExternalBackend(params, m.api.Credentials(&org))
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

func (m *middleware) UpdateExternalBackend(org string, externalBackendID uint32, purpose, cred string, def bool, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error) {
	params := organization_external_backends.NewUpdateExternalBackendParams()
	params.SetOrganizationCanonical(org)
	params.SetExternalBackendID(externalBackendID)

	var body *models.UpdateExternalBackend

	if len(cred) != 0 {
		body = &models.UpdateExternalBackend{
			Purpose:             &purpose,
			CredentialCanonical: cred,
			Default:             def,
		}
	} else {
		body = &models.UpdateExternalBackend{
			Purpose: &purpose,
			Default: def,
		}
	}

	err := ebConfig.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	body.SetConfiguration(ebConfig)
	params.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationExternalBackends.UpdateExternalBackend(params, m.api.Credentials(&org))
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
