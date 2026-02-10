package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_teams"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/go-openapi/strfmt"
	"github.com/sanity-io/litter"
)

type TeamOrderByParam string

var (
	Ascending  TeamOrderByParam = "asc"
	Descending TeamOrderByParam = "desc"
)

func (m *middleware) ListTeams(org string, teamNameFilter *string, createdAtFilter *uint64, memberIDFilter *uint32, orderBy *TeamOrderByParam) ([]*models.Team, error) {
	params := organization_teams.NewGetTeamsParams()
	params.SetOrganizationCanonical(org)
	params.SetTeamName(teamNameFilter)
	params.SetTeamCreatedAt(createdAtFilter)
	params.SetMemberID(memberIDFilter)
	params.SetOrderBy((*string)(orderBy))

	resp, err := m.api.OrganizationTeams.GetTeams(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) GetTeam(org, team string) (*models.Team, error) {
	params := organization_teams.NewGetTeamParams()
	params.SetOrganizationCanonical(org)
	params.SetTeamCanonical(team)

	resp, err := m.api.OrganizationTeams.GetTeam(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) CreateTeam(org string, name, team, owner *string, roles []string) (*models.Team, error) {
	params := organization_teams.NewCreateTeamParams()
	params.SetOrganizationCanonical(org)

	teamName, canonical, err := NameOrCanonical(name, team)
	if err != nil {
		return nil, err
	}

	body := &models.NewTeam{
		Name:           &teamName,
		Canonical:      canonical,
		RolesCanonical: roles,
	}

	if owner != nil {
		body.Owner = *owner
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("failed to validate body for CreateTeam: %w", err)
	}

	params.SetBody(body)

	resp, err := m.api.OrganizationTeams.CreateTeam(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) UpdateTeam(org string, name, team, owner *string, roles []string) (*models.Team, error) {
	params := organization_teams.NewUpdateTeamParams()
	params.SetOrganizationCanonical(org)

	teamName, canonical, err := NameOrCanonical(name, team)
	if err != nil {
		return nil, err
	}
	params.SetTeamCanonical(canonical)

	body := &models.UpdateTeam{
		Name:           &teamName,
		Canonical:      &canonical,
		RolesCanonical: roles,
	}

	if owner != nil {
		body.Owner = *owner
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, fmt.Errorf("failed to validate body for UpdateTeam: %w", err)
	}

	params.SetBody(body)

	litter.Dump(params)
	resp, err := m.api.OrganizationTeams.UpdateTeam(params, m.api.Credentials(&org))
	if err != nil {
		return nil, NewAPIError(err)
	}

	payload := resp.GetPayload()

	return payload.Data, nil
}

func (m *middleware) DeleteTeam(org, team string) error {
	params := organization_teams.NewDeleteTeamParams()
	params.SetOrganizationCanonical(org)
	params.SetTeamCanonical(team)

	_, err := m.api.OrganizationTeams.DeleteTeam(params, m.api.Credentials(&org))
	if err != nil {
		return NewAPIError(err)
	}

	return nil
}
