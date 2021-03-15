package middleware

import (
	"encoding/json"

	"github.com/cycloidio/cycloid-cli/client/client/organization_kpis"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	strfmt "github.com/go-openapi/strfmt"
)

func (m *middleware) CreateKpi(name, kpiType, widget, org, project, job, env, config string) (*models.KPI, error) {
	var body *models.NewKPI
	var err error

	params := organization_kpis.NewCreateKpiParams()
	params.SetOrganizationCanonical(org)
	pipeline := ""
	if project != "" {
		params.SetProject(&project)
	}
	if env != "" {
		params.SetEnvironment(&env)
	}

	if project != "" && env != "" {
		pipeline = common.GetPipelineName(project, env)
	}

	var uConfig interface{}
	if config != "" {
		err = json.Unmarshal([]byte(config), &uConfig)
		if err != nil {
			return nil, err
		}
	}

	body = &models.NewKPI{
		EnvironmentCanonical: env,
		Config:               uConfig,
		JobName:              job,
		Name:                 &name,
		PipelineName:         pipeline,
		ProjectCanonical:     project,
		Type:                 &kpiType,
		Widget:               &widget,
	}

	err = body.Validate(strfmt.Default)
	if err != nil {
		return nil, err
	}

	params.SetBody(body)
	resp, err := m.api.OrganizationKpis.CreateKpi(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	d := p.Data
	return d, err
}

func (m *middleware) ListKpi(org, project, env string) ([]*models.KPI, error) {
	params := organization_kpis.NewGetKpisParams()
	params.SetOrganizationCanonical(org)
	if project != "" {
		params.SetProject(&project)
	}
	if env != "" {
		params.SetEnvironment(&env)
	}

	resp, err := m.api.OrganizationKpis.GetKpis(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, err
	}

	p := resp.GetPayload()

	d := p.Data
	return d, err
}

func (m *middleware) DeleteKpi(org, kpi string) error {
	params := organization_kpis.NewDeleteKpiParams()
	params.SetOrganizationCanonical(org)
	params.SetKpiCanonical(kpi)

	_, err := m.api.OrganizationKpis.DeleteKpi(params, common.ClientCredentials(&org))
	return err
}
