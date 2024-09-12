package middleware

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/cycloidio/cycloid-cli/client/client/organizations"
	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) error {

	params := organizations.NewSendEventParams()
	params.SetOrganizationCanonical(org)

	var ts []*models.Tag
	var err error

	for k, v := range tags {
		_k := k
		_v := v
		tag := &models.Tag{
			Key:   &_k,
			Value: &_v,
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

	_, err = m.api.Organizations.SendEvent(params, m.api.Credentials(&org))
	if err != nil {
		return NewApiError(err)
	}

	return nil
}
