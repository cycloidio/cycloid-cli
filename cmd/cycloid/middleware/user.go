package middleware

import (
	"net/http"
	"net/url"

	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) UserSignup(username, email, password, fullName string) (*http.Response, error) {
	// Use a raw struct to match the actual API field names.
	// The swagger model uses given_name+family_name but the backend still accepts full_name.
	body := map[string]string{
		"username":  username,
		"email":     email,
		"password":  password,
		"full_name": fullName,
	}

	resp, err := m.GenericRequest(Request{
		Method: "POST",
		NoAuth: true,
		Route:  []string{"user"},
		Body:   body,
	}, nil)
	return resp, err
}

func (m *middleware) RefreshToken(org, childOrg *string, token string) (*models.UserSession, *http.Response, error) {
	query := url.Values{}
	if org != nil {
		query.Set("organization_canonical", *org)
	}
	if childOrg != nil {
		query.Set("child_canonical", *childOrg)
	}

	// RefreshToken uses the provided token directly, not from config
	var result *models.UserSession
	resp, err := m.GenericRequest(Request{
		Method:  "GET",
		Route:   []string{"user", "refresh_token"},
		Query:   query,
		Headers: map[string]string{"Authorization": "Bearer " + token},
		NoAuth:  true, // We manually set the Authorization header above
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UserLogin(org, email *string, password string) (*models.UserSession, *http.Response, error) {
	body := models.UserLogin{
		Password: (*strfmt.Password)(&password),
	}

	if email != nil {
		emailFmt := strfmt.Email(*email)
		body.Email = emailFmt
	}

	if org != nil {
		body.OrganizationCanonical = *org
	}

	var result *models.UserSession
	resp, err := m.GenericRequest(Request{
		Method: "POST",
		NoAuth: true,
		Route:  []string{"user", "login"},
		Body:   &body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) UserLoginToOrg(org, email, password string) (*models.UserSession, *http.Response, error) {
	emailFmt := strfmt.Email(email)
	body := models.UserLogin{
		Email:                 emailFmt,
		OrganizationCanonical: org,
		Password:              (*strfmt.Password)(&password),
	}

	var result *models.UserSession
	resp, err := m.GenericRequest(Request{
		Method: "POST",
		NoAuth: true,
		Route:  []string{"user", "login", org},
		Body:   &body,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
