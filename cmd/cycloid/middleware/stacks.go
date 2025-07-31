package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/service_catalogs"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

// skipValidationForBlueprint skips validation for blueprints since they may contain templating strings
func (m *middleware) skipValidationForBlueprint(data *models.ServiceCatalog) bool {
	return data.Blueprint
}

func (m *middleware) GetStack(org, ref string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewGetServiceCatalogParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) ListStacks(org string) ([]*models.ServiceCatalog, error) {
	params := service_catalogs.NewListServiceCatalogsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.ServiceCatalogs.ListServiceCatalogs(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}

func (m *middleware) ListBlueprints(org string) ([]*models.ServiceCatalog, error) {
	// Create a custom request with the correct query parameter for blueprint filtering
	// The frontend uses: service_catalog_blueprint[eq]=true
	// We need to add this as a custom query parameter

	// Build the URL with the correct query parameter using the configured API URL
	baseURL := m.api.Config.URL
	url := fmt.Sprintf("%s/organizations/%s/service_catalogs?organization_canonical=%s&service_catalog_blueprint%%5Beq%%5D=true",
		baseURL, org, org)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add authentication headers
	token := m.api.GetToken(&org)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the response
	var response struct {
		Data []*models.ServiceCatalog `json:"data"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// For blueprints, we skip validation entirely since they may contain templating strings
	var validBlueprints []*models.ServiceCatalog
	for _, catalog := range response.Data {
		if catalog.Blueprint {
			validBlueprints = append(validBlueprints, catalog)
		}
	}

	return validBlueprints, nil
}

func (m *middleware) CreateStackFromBlueprint(org, blueprintRef, name, canonical, serviceCatalogSourceCanonical, useCase string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewCreateServiceCatalogFromTemplateParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(blueprintRef)

	body := &models.NewServiceCatalogFromTemplate{
		Name:                          ptr.Ptr(name),
		Canonical:                     ptr.Ptr(canonical),
		ServiceCatalogSourceCanonical: ptr.Ptr(serviceCatalogSourceCanonical),
		UseCase:                       ptr.Ptr(useCase),
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "validation failed for createServiceCatalogFromTemplate input")
	}

	params.WithBody(body)

	resp, err := m.api.ServiceCatalogs.CreateServiceCatalogFromTemplate(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	// Skip validation for blueprints since they may contain templating strings
	if !m.skipValidationForBlueprint(payload.Data) {
		err = payload.Validate(strfmt.Default)
		if err != nil {
			return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
		}
	}

	return payload.Data, nil
}

func (m *middleware) UpdateStack(
	org, ref, teamCanonical string,
	visibility *string,
) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewUpdateServiceCatalogParams()
	params.WithOrganizationCanonical(org)
	params.WithServiceCatalogRef(ref)

	body := &models.UpdateServiceCatalog{
		TeamCanonical: teamCanonical,
		Visibility:    visibility,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "validation failed for updateServiceCatalog input")
	}

	params.WithBody(body)

	resp, err := m.api.ServiceCatalogs.UpdateServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	// TODO: This is a local fix for https://github.com/cycloidio/youdeploy-http-api/issues/5020
	// Remove this condition when backend will be fixed
	// If the team attribute is nil, this means that the backend did not found the maitainer canonical
	if teamCanonical != "" && payload.Data.Team == nil {
		return payload.Data, errors.Errorf(
			"maintainer with canonical '%s' may not exists, maintainer on stack ref '%s' has been removed, please check you team canonical argument and ensure that the team exists.",
			teamCanonical, ref,
		)
	}

	return payload.Data, nil
}

func (m *middleware) GetStackConfig(org, ref string) (models.ServiceCatalogConfigs, error) {
	params := service_catalogs.NewGetServiceCatalogConfigParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogConfig(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}
