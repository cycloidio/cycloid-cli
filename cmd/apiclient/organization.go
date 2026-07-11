package apiclient

import (
	"net/http"
	"net/url"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

func (m *apiClient) CreateOrganization(name string) (*models.Organization, *http.Response, error) {
	body := &models.NewOrganization{
		Name: &name,
	}

	var result *models.Organization
	resp, err := m.GenericRequest(Request{
		Method: "POST",
		Route:  []string{"organizations"},
		Body:   body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// UpdateOrganizationOpts holds optional fields for organization update.
type UpdateOrganizationOpts struct {
	CanChildrenManageOidcMapping *bool
}

func (m *apiClient) UpdateOrganization(can, name string, opts ...UpdateOrganizationOpts) (*models.Organization, *http.Response, error) {
	body := &models.UpdateOrganization{
		Name: &name,
	}

	if len(opts) > 0 && opts[0].CanChildrenManageOidcMapping != nil {
		body.CanChildrenManageOidcMapping = *opts[0].CanChildrenManageOidcMapping
	}

	var result *models.Organization
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &can,
		Route:        []string{"organizations", can},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) GetOrganization(org string) (*models.Organization, *http.Response, error) {
	var result *models.Organization
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) ListOrganizationWorkers(org string) ([]*models.Worker, *http.Response, error) {
	var result []*models.Worker
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "workers"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) ListOrganizations() ([]*models.Organization, *http.Response, error) {
	var result []*models.Organization
	resp, err := m.GenericRequest(Request{
		Method: "GET",
		Route:  []string{"organizations"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) ListOrganizationChildrens(org string) ([]*models.Organization, *http.Response, error) {
	query := url.Values{"order_by": []string{"organization_canonical:asc"}}

	var result []*models.Organization
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "children"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) CreateOrganizationChild(org, childOrg string, childOrgName *string) (*models.Organization, *http.Response, error) {
	if childOrgName == nil {
		childOrgName = &childOrg
	}

	body := &models.NewOrganization{
		Name:      childOrgName,
		Canonical: childOrg,
	}

	var result *models.Organization
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "children"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) DeleteOrganization(org string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org},
	}, nil)
	return resp, err
}
