package middleware

import (
	"fmt"
	"regexp"

	"github.com/cycloidio/cycloid-cli/client/client/organization_credentials"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
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

	_, err = m.api.OrganizationCredentials.CreateCredential(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) GetCredential(org, cred string) (*models.Credential, error) {

	params := organization_credentials.NewGetCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(cred)

	resp, err := m.api.OrganizationCredentials.GetCredential(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	d := p.Data
	return d, err
}

func (m *middleware) DeleteCredential(org, cred string) error {

	params := organization_credentials.NewDeleteCredentialParams()
	params.SetOrganizationCanonical(org)
	params.SetCredentialCanonical(cred)

	_, err := m.api.OrganizationCredentials.DeleteCredential(params, common.ClientCredentials(&org))

	return err
}

func (m *middleware) ListCredentials(org, cType string) ([]*models.CredentialSimple, error) {

	params := organization_credentials.NewGetCredentialsParams()
	params.SetOrganizationCanonical(org)

	if cType != "" {
		params.SetCredentialType(&cType)
	}

	resp, err := m.api.OrganizationCredentials.GetCredentials(params, common.ClientCredentials(&org))
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
