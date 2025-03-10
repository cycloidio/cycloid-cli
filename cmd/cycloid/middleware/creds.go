package middleware

import (
	"fmt"
	"regexp"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_credentials"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) CreateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, can, description string) (*models.Credential, error) {

	params := organization_credentials.NewCreateCredentialParams()
	params.SetOrganizationCanonical(org)

	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		safePath := re.ReplaceAllString(name, "-")
		path = fmt.Sprintf("%s_%s", cType, safePath)
	}

	body := &models.NewCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &cType,
		Canonical:   can,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationCredentials.CreateCredential(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) UpdateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, can, description string) (*models.Credential, error) {

	params := organization_credentials.NewUpdateCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(can)

	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		safePath := re.ReplaceAllString(name, "-")
		path = fmt.Sprintf("%s_%s", cType, safePath)
	}

	body := &models.UpdateCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &cType,
		Canonical:   &can,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationCredentials.UpdateCredential(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) GetCredential(org, cred string) (*models.Credential, error) {

	params := organization_credentials.NewGetCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(cred)

	resp, err := m.api.OrganizationCredentials.GetCredential(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	p := resp.GetPayload()

	d := p.Data
	return d, nil
}

func (m *middleware) DeleteCredential(org, cred string) error {

	params := organization_credentials.NewDeleteCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(cred)

	_, err := m.api.OrganizationCredentials.DeleteCredential(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (m *middleware) ListCredentials(org, cType string) ([]*models.CredentialSimple, error) {

	params := organization_credentials.NewListCredentialsParams()
	params.SetOrganizationCanonical(org)

	if cType != "" {
		params.SetCredentialType(&cType)
	}

	resp, err := m.api.OrganizationCredentials.ListCredentials(params, m.api.Credentials(&org))
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
