package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/organization_infrastructure_policies"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

// ValidateInfraPolicies will validate the TF plan against
// OPA policies defined on the Cycloid server
func (m *middleware) ValidateInfraPolicies(org, project, env string, plan []byte) (*models.InfraPoliciesValidationResult, error) {
	params := organization_infrastructure_policies.NewValidateProjectInfraPoliciesParams()
	params.SetOrganizationCanonical(org)
	params.SetProjectCanonical(project)
	params.SetEnvironmentCanonical(env)

	tfplan := string(plan)
	params.SetBody(organization_infrastructure_policies.ValidateProjectInfraPoliciesBody{
		Tfplan: &tfplan,
	})

	res, err := m.api.OrganizationInfrastructurePolicies.ValidateProjectInfraPolicies(params, common.ClientCredentials(&org))
	if err != nil {
		errors := err.(*organization_infrastructure_policies.ValidateProjectInfraPoliciesDefault).GetPayload().Errors
		e := make([]string, len(errors))
		for _, er := range errors {
			e = append(e, *er.Message)
		}
		return nil, fmt.Errorf("unable to validate project infra policies: %v", e)
	}

	return res.GetPayload().Data, nil
}
