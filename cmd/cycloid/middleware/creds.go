package middleware

import (
	"fmt"
	"regexp"

	"github.com/cycloidio/youdeploy-cli/client/client/organization_credentials"
	"github.com/cycloidio/youdeploy-cli/client/models"
	root "github.com/cycloidio/youdeploy-cli/cmd/cycloid"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) CreateCredential(org, name, cType string, rawCred *models.CredentialRaw, path, description string) error {

	params := organization_credentials.NewCreateCredentialParams()
	params.SetOrganizationCanonical(org)

	if path == "" {
		re := regexp.MustCompile(`[^a-zA-z0-9_\-./]`)
		safePath := re.ReplaceAllString(name, "-")
		path = fmt.Sprintf("%s_%s", cType, safePath)
	}

	body := &models.CreateCredential{
		Description: description,
		Name:        &name,
		Path:        &path,
		Raw:         rawCred,
		Type:        &cType,
	}

	params.SetBody(body)
	err := body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	_, err = m.api.OrganizationCredentials.CreateCredential(params, root.ClientCredentials())

	return err
}

func (m *middleware) GetCredential(org string, cred uint32) (*models.Credential, error) {

	params := organization_credentials.NewGetCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialID(cred)

	resp, err := m.api.OrganizationCredentials.GetCredential(params, root.ClientCredentials())
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	d := p.Data
	return d, err
}

func (m *middleware) DeleteCredential(org string, cred uint32) error {

	params := organization_credentials.NewDeleteCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialID(cred)

	_, err := m.api.OrganizationCredentials.DeleteCredential(params, root.ClientCredentials())

	return err
}

func (m *middleware) ListCredentials(org, cType string) ([]*models.CredentialSimple, error) {

	params := organization_credentials.NewGetCredentialsParams()
	params.SetOrganizationCanonical(org)

	if cType != "" {
		params.SetCredentialType(&cType)
	}

	resp, err := m.api.OrganizationCredentials.GetCredentials(params, root.ClientCredentials())
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
