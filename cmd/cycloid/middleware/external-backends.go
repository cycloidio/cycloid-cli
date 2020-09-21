package middleware

import (
	"github.com/cycloidio/youdeploy-cli/client/client/organization_external_backends"
	"github.com/cycloidio/youdeploy-cli/client/models"
	"github.com/cycloidio/youdeploy-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) ListExternalBackends(org string) ([]*models.ExternalBackend, error) {

	params := organization_external_backends.NewGetExternalBackendsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationExternalBackends.GetExternalBackends(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, err
}

func (m *middleware) DeleteExternalBackend(org string, externalBackend uint32) error {

	params := organization_external_backends.NewDeleteExternalBackendParams()
	params.SetOrganizationCanonical(org)
	params.SetExternalBackendID(externalBackend)

	_, err := m.api.OrganizationExternalBackends.DeleteExternalBackend(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) CreateExternalBackends(org, project, env, purpose string, cred uint32, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, error) {

	params := organization_external_backends.NewCreateExternalBackendParams()
	params.SetOrganizationCanonical(org)

	var body *models.NewExternalBackend

	if cred != 0 {
		body = &models.NewExternalBackend{
			ProjectCanonical:     project,
			Purpose:              &purpose,
			EnvironmentCanonical: env,
			CredentialID:         cred,
		}
	} else {
		body = &models.NewExternalBackend{
			ProjectCanonical:     project,
			EnvironmentCanonical: env,
			Purpose:              &purpose,
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

	resp, err := m.api.OrganizationExternalBackends.CreateExternalBackend(params, common.ClientCredentials(&org))

	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, err
}
