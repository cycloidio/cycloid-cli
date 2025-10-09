package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/client/service_catalogs"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetStack(org, ref string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewGetServiceCatalogParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalog(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) ListStacks(org string) ([]*models.ServiceCatalog, error) {
	params := service_catalogs.NewListServiceCatalogsParams()
	params.SetOrganizationCanonical(org)

	resp, err := m.api.ServiceCatalogs.ListServiceCatalogs(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) ListStackUseCases(org, ref string) ([]*models.StackUseCase, error) {
	params := service_catalogs.NewGetServiceCatalogUseCasesParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(ref)

	resp, err := m.api.ServiceCatalogs.GetServiceCatalogUseCases(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

// ListBlueprints will list stacks that are flagged as blueprint. Uses the same route as ListStack.
// TODO: Merge this route with ListStack once we find a way to add LHS filter params to the client.
func (m *middleware) ListBlueprints(org string) ([]*models.ServiceCatalog, error) {
	// This method use a custom request because we use the (undocumented)
	// LHS filter param like the frontend does: `service_catalog_blueprint[eq]=true`
	url := fmt.Sprintf("%s/organizations/%s/service_catalogs?organization_canonical=%s&service_catalog_blueprint%%5Beq%%5D=true",
		m.api.Config.URL, org, org)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	token := m.api.GetToken(&org)
	if token == "" {
		return nil, errors.New("missing API key, please provide valid authentication using CY_API_KEY env var")
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var stacks struct {
		Data []*models.ServiceCatalog `json:"data"`
	}

	err = json.Unmarshal(body, &stacks)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	var validBlueprints []*models.ServiceCatalog
	for _, catalog := range stacks.Data {
		if catalog.Blueprint {
			validBlueprints = append(validBlueprints, catalog)
		}
	}

	// Don't validate payload on this route, not supported atm.
	return validBlueprints, nil
}

func (m *middleware) CreateStackFromBlueprint(org, blueprintRef, name, stack, catalogRepository, useCase string) (*models.ServiceCatalog, error) {
	params := service_catalogs.NewCreateServiceCatalogFromTemplateParams()
	params.SetOrganizationCanonical(org)
	params.SetServiceCatalogRef(blueprintRef)
	body := &models.NewServiceCatalogFromTemplate{
		Name:                          &name,
		Canonical:                     &stack,
		ServiceCatalogSourceCanonical: &catalogRepository,
		UseCase:                       &useCase,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, errors.Wrap(err, "validation failed for createStackFromBlueprint input")
	}
	params.WithBody(body)

	resp, err := m.api.ServiceCatalogs.CreateServiceCatalogFromTemplate(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}
	payload := resp.GetPayload()

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
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

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
