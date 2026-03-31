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

	envBody := models.NewEnvironment{
		Name:      &envName,
		Canonical: env,
		Color:     color,
	}

	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments"},
		Body:         &envBody,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UpdateEnv(org, project, env, envName, color string) (*models.Environment, *http.Response, error) {
	envBody := models.UpdateEnvironment{
		Name:  &envName,
		Color: color,
	}

	var result *models.Environment
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env},
		Body:         &envBody,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteEnv(org, project, env string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "projects", project, "environments", env},
	}, nil)
	return resp, err
}
