package middleware

import (
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetEnv(org, project, env string) (*models.Environment, *http.Response, error) {
	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateEnv(org, project, env, envName, color string) (*models.Environment, *http.Response, error) {
	envName, env, err := NameOrCanonical(&envName, &env)
	if err != nil {
		return nil, nil, err
	}

	type createOrgEnvBody struct {
		Canonical string `json:"canonical,omitempty"`
		Name      string `json:"name"`
		Type      string `json:"type"`
	}

	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "environments"},
		Body: &createOrgEnvBody{
			Canonical: env,
			Name:      envName,
			Type:      "production",
		},
	}, &result)
	if err != nil {
		return nil, resp, err
	}

	type linkEnvBody struct {
		EnvironmentCanonical string `json:"environment_canonical"`
	}

	resp, err = m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments"},
		Body:         &linkEnvBody{EnvironmentCanonical: env},
	}, nil)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

func (m *middleware) UpdateEnv(org, project, env, envName, color string) (*models.Environment, *http.Response, error) {
	type updateOrgEnvBody struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}

	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "environments", env},
		Body: &updateOrgEnvBody{
			Name: envName,
			Type: "production",
		},
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

func (m *middleware) DeleteEnv(org, project, env string, opts DeleteOptions) (*http.Response, error) {
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
