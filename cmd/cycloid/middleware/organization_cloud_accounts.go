package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListCloudAccounts(org string) ([]*models.CloudAccountDetail, *http.Response, error) {
	result, resp, err := paginatedList[*models.CloudAccountDetail](m, Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "cloud_accounts"},
	}, defaultPageSize)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetCloudAccount(org, canonical string) (*models.CloudAccountDetail, *http.Response, error) {
	var result *models.CloudAccountDetail
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "cloud_accounts", canonical},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateCloudAccount(org string, body *models.NewCloudAccount) (*models.CloudAccount, *http.Response, error) {
	var result *models.CloudAccount
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "cloud_accounts"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateCloudAccountWithCredentials(org string, body *models.NewCloudAccountWithCredentials) (*models.CloudAccount, *http.Response, error) {
	var result *models.CloudAccount
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "cloud_accounts", "with_credentials"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdateCloudAccount(org, canonical string, body *models.UpdateCloudAccount) (*models.CloudAccount, *http.Response, error) {
	var result *models.CloudAccount
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "cloud_accounts", canonical},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteCloudAccount(org, canonical string) (*http.Response, error) {
	return m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "cloud_accounts", canonical},
	}, nil)
}
