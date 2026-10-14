package apiclient

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// AdminDeleteUser deletes any user of the instance, including one without an organization. org has
// to be the root organization, and the caller one of its admins
func (m *apiClient) AdminDeleteUser(org, email string) (*http.Response, error) {
	return m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "admin", "users", email},
	}, nil)
}

// AdminListUsers lists the users of the whole instance, including the ones without an
// organization. org has to be the root organization, and the caller one of its admins.
//
// Supported LHS filter attributes: user_canonical, user_full_name, user_email (the primary one),
// user_created_at, user_mfa_enabled, user_has_organization.
func (m *apiClient) AdminListUsers(org string, filters ...LHSFilter) ([]*models.AdminUser, *http.Response, error) {
	result, resp, err := paginatedList[*models.AdminUser](m, Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "admin", "users"},
		LHSFilters:   filters,
	}, defaultPageSize)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
