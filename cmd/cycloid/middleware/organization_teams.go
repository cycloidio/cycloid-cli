package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/cycloidio/cycloid-cli/client/models"
)

type TeamOrderByParam string

var (
	Ascending  TeamOrderByParam = "asc"
	Descending TeamOrderByParam = "desc"
)

func (m *middleware) ListTeams(org string, teamNameFilter *string, createdAtFilter *uint64, memberIDFilter *uint32, orderBy *TeamOrderByParam) ([]*models.Team, *http.Response, error) {
	query := url.Values{}
	if teamNameFilter != nil {
		query.Set("team_name", *teamNameFilter)
	}
	if createdAtFilter != nil {
		query.Set("team_created_at", strconv.FormatUint(*createdAtFilter, 10))
	}
	if memberIDFilter != nil {
		query.Set("member_id", strconv.FormatUint(uint64(*memberIDFilter), 10))
	}
	if orderBy != nil {
		query.Set("order_by", string(*orderBy))
	}

	var result []*models.Team
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "teams"},
		Query:        query,
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) GetTeam(org, team string) (*models.Team, *http.Response, error) {
	var result *models.Team
	resp, err := m.GenericRequest(Request{
		Method:       "GET",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", team},
	}, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

func (m *middleware) CreateTeam(org string, name, team, owner *string, roles []string) (*models.Team, *http.Response, error) {
	teamName, canonical, err := NameOrCanonical(name, team)
	if err != nil {
		return nil, nil, err
	}

	body := &models.NewTeam{
		Name:           &teamName,
		Canonical:      canonical,
		RolesCanonical: roles,
	}

	if owner != nil {
		body.Owner = *owner
	}

	var result *models.Team
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "teams"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create team: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) UpdateTeam(org string, name, team, owner *string, roles []string) (*models.Team, *http.Response, error) {
	teamName, canonical, err := NameOrCanonical(name, team)
	if err != nil {
		return nil, nil, err
	}

	body := &models.UpdateTeam{
		Name:           &teamName,
		Canonical:      &canonical,
		RolesCanonical: roles,
	}

	if owner != nil {
		body.Owner = *owner
	}

	var result *models.Team
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", canonical},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update team: %w", err)
	}
	return result, resp, nil
}

func (m *middleware) DeleteTeam(org, team string) (*http.Response, error) {
	resp, err := m.GenericRequest(Request{
		Method:       "DELETE",
		Organization: &org,
		Route:        []string{"organizations", org, "teams", team},
	}, nil)
	return resp, err
}
