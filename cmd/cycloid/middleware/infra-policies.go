package middleware

import (
	"fmt"
	"os"

	"github.com/cycloidio/cycloid-cli/client/client/organization_infrastructure_policies"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

// ValidateInfraPolicies will validate the TF plan against
// OPA policies defined on the Cycloid server
func (m *middleware) ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, error) {
	params := organization_infrastructure_policies.NewValidateProjectInfraPoliciesParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	tfplan := string(plan)
	params.SetBody(&models.TerraformPlanInput{
		Tfplan: &tfplan,
	})

	resp, err := m.api.OrganizationInfrastructurePolicies.ValidateProjectInfraPolicies(params, m.api.Credentials(&org))
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

// CreateInfraPoliciy will create a new infraPolicy
// with the rego file suplied
func (m *middleware) CreateInfraPolicy(org, policyFile, policyCanonical, description, policyName, ownerCannonical, severity string, enabled bool) (*models.InfraPolicy, error) {
	params := organization_infrastructure_policies.NewCreateInfraPolicyParams()
	params.SetOrganizationCanonical(org)
	// Reads file content and converts it into string
	policyFileContent, err := os.ReadFile(policyFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to read rego file: %v", err)
	}
	// If cannonical empty,use the default one
	if policyCanonical == "" {
		policyCanonical = common.GenerateCanonical(policyName)
	}

	policyBody := string(policyFileContent)
	body := &models.NewInfraPolicy{
		Body:        &policyBody,
		Canonical:   policyCanonical,
		Description: description,
		Enabled:     enabled,
		Name:        &policyName,
		Owner:       &ownerCannonical,
		Severity:    &severity,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("InfraPolicy invalid: %v", err)
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationInfrastructurePolicies.CreateInfraPolicy(params, m.api.Credentials(&org))
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

// DeleteInfraPolicy will delete a infraPolicy
func (m *middleware) DeleteInfraPolicy(org, policyCannonical string) error {
	params := organization_infrastructure_policies.NewDeleteInfraPolicyParams()
	params.SetOrganizationCanonical(org)
	params.SetInfraPolicyCanonical(policyCannonical)

	_, err := m.api.OrganizationInfrastructurePolicies.DeleteInfraPolicy(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}
	return nil
}

// ListInfraPolicies will list all infraPolicies in an organisation
func (m *middleware) ListInfraPolicies(org string) ([]*models.InfraPolicy, error) {

	params := organization_infrastructure_policies.NewGetInfraPoliciesParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.OrganizationInfrastructurePolicies.GetInfraPolicies(params, m.api.Credentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()
	err = p.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	d := p.Data

	return d, nil
}

// GetInfraPolicy will list all infraPolicies in an organisation
func (m *middleware) GetInfraPolicy(org, infraPolicy string) (*models.InfraPolicy, error) {

	params := organization_infrastructure_policies.NewGetInfraPolicyParams()
	params.SetOrganizationCanonical(org)
	params.SetInfraPolicyCanonical(infraPolicy)

	resp, err := m.api.OrganizationInfrastructurePolicies.GetInfraPolicy(params, m.api.Credentials(&org))
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

// UpdateInfraPolicy will update an existing infrapolicy with the given params
func (m *middleware) UpdateInfraPolicy(org, infraPolicy, policyFile, description, policyName, ownerCannonical, severity string, enabled bool) (*models.InfraPolicy, error) {

	params := organization_infrastructure_policies.NewUpdateInfraPolicyParams()
	params.SetOrganizationCanonical(org)
	params.SetInfraPolicyCanonical(infraPolicy)

	// Reads file content and converts it into string
	policyFileContent, err := os.ReadFile(policyFile)
	if err != nil {
		return nil, fmt.Errorf("Unable to read rego file: %v", err)
	}
	policyBody := string(policyFileContent)

	body := &models.UpdateInfraPolicy{
		Body:        &policyBody,
		Description: &description,
		Enabled:     &enabled,
		Name:        &policyName,
		Owner:       &ownerCannonical,
		Severity:    &severity,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("InfraPolicy invalid: %v", err)
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationInfrastructurePolicies.UpdateInfraPolicy(params, m.api.Credentials(&org))
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
