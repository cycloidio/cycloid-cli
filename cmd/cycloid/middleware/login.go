package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/user"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

// Login is the method used to log the user into the Cycloid console
func (m *middleware) Login(email, password string) (*models.UserSession, error) {

	mail := strfmt.Email(email)
	p := strfmt.Password(password)

	params := user.NewLoginParams()
	body := &models.UserLogin{
		Email:    mail,
		Password: &p,
	}

	if err := body.Validate(strfmt.Default); err != nil {
		return nil, errors.Wrap(err, "unable to validate body request")
	}

	params.SetBody(body)

	res, err := m.api.User.Login(params)
	if err != nil {
		return nil, errors.Wrap(err, "unable to log user")
	}
	return res.GetPayload().Data, nil
}

func (m *middleware) LoginOrg(org, child, username, password string) (*models.UserSession, error) {

	params := user.NewRefreshTokenParams()
	params.SetOrganizationCanonical(&org)
	if len(child) != 0 {
		params.SetChildCanonical(&child)
	}
	res, err := m.api.User.RefreshToken(params, common.ClientCredentials(nil))
	if err != nil {
		return nil, errors.Wrap(err, "unable to log user")
	}
	return res.GetPayload().Data, nil
}
