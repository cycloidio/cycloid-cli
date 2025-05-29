package middleware

import (
	"fmt"

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

func (m *middleware) ListEvents(org string, eventType, eventSeverity []string, begin, end uint64) ([]*models.Event, error) {
	params := organizations.NewGetEventsParams()
	params.WithOrganizationCanonical(org)
	params.Begin = &begin
	params.End = &end

	if len(eventSeverity) != 0 {
		params.WithSeverity(eventSeverity)
	}

	if len(eventType) != 0 {
		params.WithType(eventType)
	}

	resp, err := m.api.Organizations.GetEvents(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewApiError(err)
	}

	payload := resp.GetPayload()
	err = payload.Validate(strfmt.Default)
	if err != nil {
		return payload.Data, fmt.Errorf("invalid response from the API: %v", err)
	}

	return payload.Data, nil
}
