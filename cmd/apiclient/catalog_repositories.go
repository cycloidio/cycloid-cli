package apiclient

import (
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

// ListCatalogRepositories lists catalog repositories for an organization.
//
// NOTE: the backend handler for this route does not call lhs.ParseQuery, so
// LHS filters are accepted by the apiClient but silently ignored server-side.
func (m *apiClient) ListCatalogRepositories(org string, filters ...LHSFilter) ([]*models.ServiceCatalogSource, *http.Response, error) {
	var result []*models.ServiceCatalogSource
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources"},
		LHSFilters:   filters,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, *http.Response, error) {
	var result *models.ServiceCatalogSource
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources", catalogRepo},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) DeleteCatalogRepository(org, catalogRepo string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources", catalogRepo},
	}, nil)
	return resp, err
}

func (m *apiClient) CreateCatalogRepository(org, name, url, branch, cred, visibility, teamCanonical string) (*models.ServiceCatalogSource, *http.Response, error) {
	var body *models.NewServiceCatalogSource

	if len(cred) != 0 {
		body = &models.NewServiceCatalogSource{
			Branch:              &branch,
			CredentialCanonical: cred,
			Name:                &name,
			URL:                 &url,
		}
	} else {
		body = &models.NewServiceCatalogSource{
			Branch: &branch,
			Name:   &name,
			URL:    &url,
		}
	}

	switch visibility {
	case "shared", "local", "hidden":
		body.Visibility = visibility
	case "":
		break
	default:
		return nil, nil, errors.New("invalid visibility parameter for CreateCatalogRepository, accepted values are 'local', 'shared' or 'hidden'")
	}

	if teamCanonical != "" {
		body.TeamCanonical = teamCanonical
	}

	var result *models.ServiceCatalogSource
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) UpdateCatalogRepository(org, catalogRepo, name, url, branch, cred string, visibility *string) (*models.ServiceCatalogSource, *http.Response, error) {
	body := &models.UpdateServiceCatalogSource{
		Branch:              branch,
		CredentialCanonical: cred,
		Name:                &name,
		URL:                 &url,
	}

	var result *models.ServiceCatalogSource
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources", catalogRepo},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *apiClient) RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogChanges, *http.Response, error) {
	var result *models.ServiceCatalogChanges
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources", catalogRepo, "refresh"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// RefreshCatalogRepositoryVersions triggers an immediate re-index of all branches and tags
// for the given catalog repository. The backend clones the git repository, fetches all refs,
// and updates the service_catalog_source_versions table synchronously before returning.
//
// This resolves the eventual-consistency race where a freshly created catalog repository has no
// version rows yet (the background cron that populates them runs every ~10 minutes by default).
func (m *apiClient) RefreshCatalogRepositoryVersions(org, catalogRepo string) ([]*StackVersion, *http.Response, error) {
	var result []*StackVersion
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources", catalogRepo, "versions", "refresh"},
		Query: url.Values{
			"sync_presence": []string{"true"},
		},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
