package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_components"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/go-openapi/strfmt"
)

func (m *middleware) GetComponentConfig(org, project, env, component string) (*models.FormVariables, error) {
	params := organization_components.NewGetComponentConfigParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	resp, err := m.api.OrganizationComponents.GetComponentConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return &resp.GetPayload().Data, nil
}

func (m *middleware) GetComponent(org, project, env, component string) (*models.Component, error) {
	params := organization_components.NewGetComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	resp, err := m.api.OrganizationComponents.GetComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.GetPayload(), nil
}

func (m *middleware) GetComponents(org, project, env string) ([]*models.Component, error) {
	params := organization_components.NewGetComponentsParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)

	resp, err := m.api.OrganizationComponents.GetComponents(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.GetPayload().Data, nil
}

func (m *middleware) CreateComponent(org, project, env, component, description string, componentName, serviceCatalogRef, useCase, cloudProviderCanonical *string, vars *models.FormVariables) (*models.Component, error) {
	params := organization_components.NewCreateComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	body := &models.NewComponent{
		Name:              componentName,
		Canonical:         component,
		Description:       description,
		ServiceCatalogRef: serviceCatalogRef,
		UseCase:           useCase,
	}

	if vars != nil {
		body.Vars = *vars
	}

	if cloudProviderCanonical != nil {
		body.CloudProviderCanonical = *cloudProviderCanonical
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("createComponent parameter validation failed, body:\n%v\nerr: %v", body, err)
	}
	params.WithBody(body)

	resp, err := m.api.OrganizationComponents.CreateComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.GetPayload().Data, nil
}

func (m *middleware) UpdateComponent(org, project, env, component, description string, componentName, useCase *string, vars *models.FormVariables) (*models.Component, error) {
	params := organization_components.NewUpdateComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)
	body := &models.UpdateComponent{
		Name:        componentName,
		Description: description,
		UseCase:     useCase,
	}

	if vars != nil {
		body.Vars = *vars
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("createComponent parameter validation failed, body:\n%v\nerr: %v", body, err)
	}
	params.WithBody(body)

	resp, err := m.api.OrganizationComponents.UpdateComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.GetPayload(), nil
}

func (m *middleware) MigrateComponent(org, project, env, component, targetProject, targetEnv string) (*models.Component, error) {
	params := organization_components.NewMigrateComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)
	body := models.MigrateComponent{
		DestinationProjectCanonical:     targetProject,
		DestinationEnvironmentCanonical: targetEnv,
	}
	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("migrateComponent body validation failed, body:\n%v\nerr: %v", body, err)
	}

	params.WithBody(&body)

	resp, err := m.api.OrganizationComponents.MigrateComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	return resp.GetPayload(), nil
}

func (m *middleware) DeleteComponent(org, project, env, component string) error {
	params := organization_components.NewDeleteComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	_, err := m.api.OrganizationComponents.DeleteComponent(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}
