package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

func (m *middleware) SendEvent(org, eventType, title, message, severity string, tags map[string]string, color string) (*http.Response, error) {
	var ts []*models.Tag

	for k, v := range tags {
		_k := k
		_v := v
		tag := &models.Tag{
			Key:   &_k,
			Value: &_v,
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

	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "events"},
		Body:         body,
	}, nil)
	return resp, err
}

func (m *middleware) ListEvents(org string, eventType, eventSeverity []string, begin, end uint64) ([]*models.Event, *http.Response, error) {
	query := url.Values{}
	query.Set("begin", strconv.FormatUint(begin, 10))
	query.Set("end", strconv.FormatUint(end, 10))
	for _, s := range eventSeverity {
		query.Add("severity", s)
	}
	for _, t := range eventType {
		query.Add("type", t)
	}

	var result []*models.Event
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "events"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list events: %w", err)
	}
	return result, resp, nil
}
