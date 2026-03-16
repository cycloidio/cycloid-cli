package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

// ValidateInfraPolicies will validate the TF plan against
// OPA policies defined on the Cycloid server
func (m *middleware) ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, *http.Response, error) {
	tfplan := string(plan)
	body := &models.TerraformPlanInput{
		Tfplan: &tfplan,
	}

	var result *models.InfraPoliciesValidationResult
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env, "validate_infra_policies"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// CreateInfraPolicy will create a new infraPolicy with the rego file supplied
func (m *middleware) CreateInfraPolicy(org, policyFile, policyCanonical, description, policyName, ownercanonical, severity string, enabled bool) (*models.InfraPolicy, *http.Response, error) {
	// Reads file content and converts it into string
	policyFileContent, err := os.ReadFile(policyFile)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read rego file: %w", err)
	}
	// If canonical empty, use the default one
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
		Owner:       &ownercanonical,
		Severity:    &severity,
	}

	var result *models.InfraPolicy
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "infra_policies"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create infra policy: %w", err)
	}
	return result, resp, nil
}

// DeleteInfraPolicy will delete a infraPolicy
func (m *middleware) DeleteInfraPolicy(org, policycanonical string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "infra_policies", policycanonical},
	}, nil)
	return resp, err
}

// ListInfraPolicies will list all infraPolicies in an organization
func (m *middleware) ListInfraPolicies(org string) ([]*models.InfraPolicy, *http.Response, error) {
	var result []*models.InfraPolicy
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "infra_policies"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetInfraPolicy will get a specific infra policy
func (m *middleware) GetInfraPolicy(org, infraPolicy string) (*models.InfraPolicy, *http.Response, error) {
	var result *models.InfraPolicy
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "infra_policies", infraPolicy},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UpdateInfraPolicy will update an existing infrapolicy with the given params
func (m *middleware) UpdateInfraPolicy(org, infraPolicy, policyFile, description, policyName, ownercanonical, severity string, enabled bool) (*models.InfraPolicy, *http.Response, error) {
	// Reads file content and converts it into string
	policyFileContent, err := os.ReadFile(policyFile)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read rego file: %w", err)
	}
	policyBody := string(policyFileContent)

	body := &models.UpdateInfraPolicy{
		Body:        &policyBody,
		Description: &description,
		Enabled:     &enabled,
		Name:        &policyName,
		Owner:       &ownercanonical,
		Severity:    &severity,
	}

	var result *models.InfraPolicy
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "infra_policies", infraPolicy},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update infra policy: %w", err)
	}
	return result, resp, nil
}
