package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/user"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) UserSignup(username, email, password, givenName, familyName string) error {
	body := models.NewUserAccount{
		Username:   &username,
		Email:      (*strfmt.Email)(&email),
		Password:   (*strfmt.Password)(&password),
		GivenName:  &givenName,
		FamilyName: &familyName,
	}

	err := body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	params := user.NewSignUpParams()
	params.WithBody(&body)

	_, err = m.api.User.SignUp(params)
	if err != nil {
		return NewApiError(err)
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
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) UserLogin(email, username, org *string, password string) (*models.UserSession, error) {
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

	err := body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.WithBody(&body)
	resp, err := m.api.User.Login(params)
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}
