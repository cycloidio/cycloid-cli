package middleware

import (
	"fmt"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_components"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetComponentConfig(org, project, env, component string) (models.FormVariables, error) {
	params := organization_components.NewGetComponentConfigParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	resp, err := m.api.OrganizationComponents.GetComponentConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) GetComponent(org, project, env, component string) (*models.Component, error) {
	params := organization_components.NewGetComponentParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)

	resp, err := m.api.OrganizationComponents.GetComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) ListComponents(org, project, env string) ([]*models.Component, error) {
	params := organization_components.NewGetComponentsParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)

	resp, err := m.api.OrganizationComponents.GetComponents(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) CreateComponent(org, project, env, component, description string, componentName, serviceCatalogRef *string, cloudProviderCanonical string) (*models.Component, error) {
	params := organization_components.NewCreateComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)

	body := &models.NewComponent{
		Name:                   componentName,
		Canonical:              component,
		Description:            description,
		ServiceCatalogRef:      serviceCatalogRef,
		CloudProviderCanonical: cloudProviderCanonical,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("createComponent parameter validation failed, body:\n%v\nerr: %v", body, err)
	}
	params.WithBody(body)

	resp, err := m.api.OrganizationComponents.CreateComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) CreateAndConfigureComponent(org, project, env, component, description string, componentName *string, serviceCatalogRef, useCase, cloudProviderCanonical string, vars models.FormVariables) (*models.Component, error) {
	params := organization_components.NewCreateAndConfigureComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)

	body := &models.NewAndConfiguredComponent{
		Name:                   componentName,
		Canonical:              component,
		Description:            description,
		ServiceCatalogRef:      serviceCatalogRef,
		UseCase:                useCase,
		Vars:                   vars,
		CloudProviderCanonical: cloudProviderCanonical,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("createAndConfigureComponent body validation failed, body:\n%v\nerr: %v", body, err)
	}

	params.WithBody(body)

	resp, err := m.api.OrganizationComponents.CreateAndConfigureComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil

}

func (m *middleware) UpdateComponent(org, project, env, component, description string, componentName *string) (*models.Component, error) {
	params := organization_components.NewUpdateComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	body := &models.UpdateComponent{
		Name:        componentName,
		Description: description,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("updateComponent parameter validation failed, body:\n%v\nerr: %v", body, err)
	}
	params.WithBody(body)

	resp, err := m.api.OrganizationComponents.UpdateComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) ConfigureComponent(org, project, env, component, useCase string, vars models.FormVariables) error {
	params := organization_components.NewConfigureComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	body := &models.ConfigureComponent{
		UseCase: &useCase,
		Vars:    vars,
	}

	params.WithBody(body)
	_, err := m.api.OrganizationComponents.ConfigureComponent(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}

	return nil
}

func (m *middleware) MigrateComponent(org, project, env, component, targetProject, targetEnv, newCanonical, newName string) (*models.Component, error) {
	params := organization_components.NewMigrateComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)
	body := models.MigrateComponent{
		DestinationProjectCanonical:     targetProject,
		DestinationEnvironmentCanonical: targetEnv,
		DestinationComponentCanonical:   newCanonical,
		DestinationComponentName:        newName,
	}

	params.WithBody(&body)

	resp, err := m.api.OrganizationComponents.MigrateComponent(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) DeleteComponent(org, project, env, component string) error {
	params := organization_components.NewDeleteComponentParams()
	params.WithOrganizationCanonical(org)
	params.WithProjectCanonical(project)
	params.WithEnvironmentCanonical(env)
	params.WithComponentCanonical(component)

	_, err := m.api.OrganizationComponents.DeleteComponent(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}

	return nil
}

func (m *middleware) GetComponentStackConfig(org, project, env, component, useCase string) (models.ServiceCatalogConfigs, error) {
	params := organization_components.NewGetComponentStackConfigurationParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)
	params.SetComponentCanonical(component)
	params.SetUseCase(&useCase)

	resp, err := m.api.OrganizationComponents.GetComponentStackConfiguration(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}
