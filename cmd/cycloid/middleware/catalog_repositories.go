package middleware

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) ListCatalogRepositories(org string) ([]*models.ServiceCatalogSource, *http.Response, error) {
	var result []*models.ServiceCatalogSource
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogSource, *http.Response, error) {
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

func (m *middleware) DeleteCatalogRepository(org, catalogRepo string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "service_catalog_sources", catalogRepo},
	}, nil)
	return resp, err
}

func (m *middleware) CreateCatalogRepository(org, name, url, branch, cred, visibility, teamCanonical string) (*models.ServiceCatalogSource, *http.Response, error) {
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

func (m *middleware) UpdateCatalogRepository(org, catalogRepo string, name, url, branch, cred string, visibility *string) (*models.ServiceCatalogSource, *http.Response, error) {
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

func (m *middleware) RefreshCatalogRepository(org, catalogRepo string) (*models.ServiceCatalogChanges, *http.Response, error) {
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
