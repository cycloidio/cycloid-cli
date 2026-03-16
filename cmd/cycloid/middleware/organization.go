package middleware

import (
	"net/http"
	"net/url"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) CreateOrganization(name string) (*models.Organization, *http.Response, error) {
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

func (m *middleware) UpdateOrganization(can, name string) (*models.Organization, *http.Response, error) {
	body := &models.UpdateOrganization{
		Name: &name,
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

func (m *middleware) GetOrganization(org string) (*models.Organization, *http.Response, error) {
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

func (m *middleware) ListOrganizationWorkers(org string) ([]*models.Worker, *http.Response, error) {
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

func (m *middleware) ListOrganizations() ([]*models.Organization, *http.Response, error) {
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

func (m *middleware) ListOrganizationChildrens(org string) ([]*models.Organization, *http.Response, error) {
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

func (m *middleware) CreateOrganizationChild(org, childOrg string, childOrgName *string) (*models.Organization, *http.Response, error) {
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

func (m *middleware) DeleteOrganization(org string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org},
	}, nil)
	return resp, err
}
