package apiclient

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/gen/models"
)

func (m *middleware) ListOrgEnvs(org string, filters ...LHSFilter) ([]*models.Environment, *http.Response, error) {
	result, resp, err := paginatedList[*models.Environment](m, Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "environments"},
		LHSFilters:   filters,
	}, defaultPageSize)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetOrgEnv(org, env string) (*models.Environment, *http.Response, error) {
	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "environments", env},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateOrgEnv(org string, body *models.NewEnvironment) (*models.Environment, *http.Response, error) {
	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "environments"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdateOrgEnv(org, env string, body *models.UpdateEnvironment) (*models.Environment, *http.Response, error) {
	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "environments", env},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteOrgEnv(org, env string) (*http.Response, error) {
	return m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "environments", env},
	}, nil)
}

func (m *middleware) LinkEnvToProject(org, project, env string) (*http.Response, error) {
	envs, _, err := m.ListProjectEnvs(org, project)
	if err != nil {
		return nil, err
	}
	for _, linked := range envs {
		if linked.Canonical != nil && *linked.Canonical == env {
			return nil, nil
		}
	}

	type linkEnvBody struct {
		EnvironmentCanonical string `json:"environment_canonical"`
	}

	return m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments"},
		Body:         &linkEnvBody{EnvironmentCanonical: env},
	}, nil)
}

func (m *middleware) UnlinkEnvFromProject(org, project, env string, opts DeleteOptions) (*http.Response, error) {
	req := Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env},
	}
	if q := opts.Resolve(); q.SkipHooks || q.IgnoreConfigFilesErr {
		req.Query = q
	}
	return m.GenericRequest(req, nil)
}
