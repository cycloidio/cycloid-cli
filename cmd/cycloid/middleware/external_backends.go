package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) GetRemoteTFExternalBackend(org string) (*models.ExternalBackend, error) {
	data, _, err := m.ListExternalBackends(org)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("couldn't find the remote terraform backend")
	}
	return data[0], nil
}

func (m *middleware) GetExternalBackend(org string, externalBackend uint32) (*models.ExternalBackend, *http.Response, error) {
	var result *models.ExternalBackend
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "external_backends", strconv.FormatUint(uint64(externalBackend), 10)},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) ListExternalBackends(org string) ([]*models.ExternalBackend, *http.Response, error) {
	var result []*models.ExternalBackend
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "external_backends"},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) DeleteExternalBackend(org string, externalBackend uint32) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "external_backends", strconv.FormatUint(uint64(externalBackend), 10)},
	}, nil)
	return resp, err
}

func (m *middleware) CreateExternalBackends(org, project, env, purpose, credential string, isDefault bool, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, *http.Response, error) {
	var body *models.NewExternalBackend

	if len(credential) != 0 {
		body = &models.NewExternalBackend{
			ProjectCanonical:     project,
			Purpose:              &purpose,
			EnvironmentCanonical: env,
			CredentialCanonical:  credential,
			Default:              isDefault,
		}
	} else {
		body = &models.NewExternalBackend{
			ProjectCanonical:     project,
			EnvironmentCanonical: env,
			Purpose:              &purpose,
			Default:              isDefault,
		}
	}

	body.SetConfiguration(ebConfig)

	var result *models.ExternalBackend
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "external_backends"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create external backend: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) UpdateExternalBackend(org string, externalBackendID uint32, purpose, credential string, isDefault bool, ebConfig models.ExternalBackendConfiguration) (*models.ExternalBackend, *http.Response, error) {
	var body *models.UpdateExternalBackend

	if len(credential) != 0 {
		body = &models.UpdateExternalBackend{
			Purpose:             &purpose,
			CredentialCanonical: credential,
			Default:             isDefault,
		}
	} else {
		body = &models.UpdateExternalBackend{
			Purpose: &purpose,
			Default: isDefault,
		}
	}

	body.SetConfiguration(ebConfig)

	var result *models.ExternalBackend
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "external_backends", strconv.FormatUint(uint64(externalBackendID), 10)},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update external backend: %w", err)
	}
	return result, resp, nil
}
