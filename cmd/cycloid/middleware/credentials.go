package middleware

import (
	"fmt"
	"regexp"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_credentials"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) CreateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, error) {
	params := organization_credentials.NewCreateCredentialParams()
	params.SetOrganizationCanonical(org)

	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		safePath := re.ReplaceAllString(name, "-")
		path = fmt.Sprintf("%s_%s", credentialType, safePath)
	}

	body := &models.NewCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &credentialType,
		Canonical:   canonical,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationCredentials.CreateCredential(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) UpdateCredential(org, name, credentialType string, rawCred *models.CredentialRaw, path, canonical, description string) (*models.Credential, error) {
	params := organization_credentials.NewUpdateCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(canonical)

	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		safePath := re.ReplaceAllString(name, "-")
		path = fmt.Sprintf("%s_%s", credentialType, safePath)
	}

	body := &models.UpdateCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &credentialType,
		Canonical:   &canonical,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	resp, err := m.api.OrganizationCredentials.UpdateCredential(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) GetCredential(org, credential string) (*models.Credential, error) {
	params := organization_credentials.NewGetCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(credential)

	resp, err := m.api.OrganizationCredentials.GetCredential(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) DeleteCredential(org, credential string) error {
	params := organization_credentials.NewDeleteCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(credential)

	_, err := m.api.OrganizationCredentials.DeleteCredential(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}

	return nil
}

func (m *middleware) ListCredentials(org, credentialType string) ([]*models.CredentialSimple, error) {
	params := organization_credentials.NewListCredentialsParams()
	params.SetOrganizationCanonical(org)

	if credentialType != "" {
		params.SetCredentialType(&credentialType)
	}

	resp, err := m.api.OrganizationCredentials.ListCredentials(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}
