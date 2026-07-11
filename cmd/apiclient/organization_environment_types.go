package apiclient

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

func (m *apiClient) ListEnvironmentTypes(org string) ([]*models.EnvironmentType, *http.Response, error) {
	result, resp, err := paginatedList[*models.EnvironmentType](m, Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "environment_types"},
	}, defaultPageSize)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) GetEnvironmentType(org, canonical string) (*models.EnvironmentType, *http.Response, error) {
	types, resp, err := m.ListEnvironmentTypes(org)
	if err != nil {
		return nil, resp, err
	}
	for _, envType := range types {
		if envType.Canonical != nil && *envType.Canonical == canonical {
			return envType, resp, nil
		}
	}
	return nil, resp, &APIResponseError{
		StatusCode: http.StatusNotFound,
		Status:     "404 Not Found",
		Body:       []byte(fmt.Sprintf("environment type %q not found", canonical)),
	}
}

func (m *apiClient) CreateEnvironmentType(org string, body *models.NewEnvironmentType) (*models.EnvironmentType, *http.Response, error) {
	var result *models.EnvironmentType
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "environment_types"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) UpdateEnvironmentType(org, canonical string, body *models.UpdateEnvironmentType) (*models.EnvironmentType, *http.Response, error) {
	var result *models.EnvironmentType
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "environment_types", canonical},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) DeleteEnvironmentType(org, canonical string) (*http.Response, error) {
	return m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "environment_types", canonical},
	}, nil)
}
