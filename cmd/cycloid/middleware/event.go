package middleware

import (
	"github.com/cycloidio/cycloid-cli/client/client/organizations"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error {

	params := organizations.NewSendOrgEventParams()
	params.SetOrganizationCanonical(org)

	var ts []*models.Tag
	var err error

	for k, v := range tags {
		tag := &models.Tag{
			Key:   &k,
			Value: &v,
		}
		err = tag.Validate(strfmt.Default)
		if err != nil {
			continue
		}

		ts = append(ts, tag)
	}

	body := &models.NewEvent{
		Tags:     ts,
		Type:     &eventType,
		Title:    &title,
		Color:    color,
		Severity: &severity,
		Message:  &message,
	}

	params.SetBody(body)
	err = body.Validate(strfmt.Default)
	if err != nil {
		return err
	}

	_, err = m.api.Organizations.SendOrgEvent(params, common.ClientCredentials(&org))

	return err
}
