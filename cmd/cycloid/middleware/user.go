package middleware

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/user"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) UserSignup(username, email, password, fullName string) error {
	params := user.NewSignUpParams()
	body := &models.NewUserAccount{
		Email:    (*strfmt.Email)(&email),
		Password: (*strfmt.Password)(&password),
		FullName: &fullName,
	}

	// Username is optional
	if username != "" {
		body.Username = username
	}

	params.WithBody(body)

	_, err := m.api.User.SignUp(params)
	if err != nil {
		return NewAPIError(err)
	}

	return nil
}

func (m *middleware) RefreshToken(org, childOrg *string, token string) (*models.UserSession, error) {
	params := user.NewRefreshTokenParams()
	if org != nil {
		params.WithOrganizationCanonical(org)
	}

	if childOrg != nil {
		params.WithChildCanonical(childOrg)
	}

	resp, err := m.api.User.RefreshToken(params,
		runtime.ClientAuthInfoWriterFunc(
			func(r runtime.ClientRequest, _ strfmt.Registry) error {
				r.SetHeaderParam("Authorization", "Bearer "+token)
				return nil
			},
		),
	)
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	return payload.Data, nil
}

func (m *middleware) UserLogin(org, email *string, password string) (*models.UserSession, error) {
	params := user.NewLoginParams()
	body := models.UserLogin{
		Password: (*strfmt.Password)(&password),
	}

	if email != nil {
		body.Email = (*strfmt.Email)(email)
	}

	if org != nil {
		body.OrganizationCanonical = *org
	}

	params.WithBody(&body)
	resp, err := m.api.User.Login(params)
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	return payload.Data, nil
}

func (m *middleware) UserLoginToOrg(org, email, password string) (*models.UserSession, error) {
	params := user.NewLoginToOrgParams()
	params.WithOrganizationCanonical(org)
	body := models.UserLogin{
		Email:                 (*strfmt.Email)(&email),
		OrganizationCanonical: org,
		Password:              (*strfmt.Password)(&password),
	}
	params.WithBody(&body)

	resp, err := m.api.User.LoginToOrg(params)
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()
	return payload.Data, nil
}
