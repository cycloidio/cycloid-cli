package middleware

import (
	"fmt"

	"github.com/cycloidio/cycloid-cli/client/client/cost_estimation"
	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
)

// CostEstimation will consume the backend API endpoint for cost estimation
func (m *middleware) CostEstimation(org string, plan []byte) (*models.CostEstimationResult, error) {
	params := cost_estimation.NewCostEstimateTfPlanParams()
	params.SetOrganizationCanonical(org)

	tfplan := string(plan)
	params.SetBody(cost_estimation.CostEstimateTfPlanBody{
		Tfplan: &tfplan,
	})

	res, err := m.api.CostEstimation.CostEstimateTfPlan(params, common.ClientCredentials(&org))
	if err != nil {
		if _, ok := err.(*cost_estimation.CostEstimateTfPlanDefault); ok {
			errors := err.(*cost_estimation.CostEstimateTfPlanDefault).GetPayload().Errors
			e := make([]string, len(errors))
			for _, er := range errors {
				e = append(e, *er.Message)
			}
			return nil, fmt.Errorf("unable to estimate cost insfrastructure: %v", e)
		}
		return nil, fmt.Errorf("unable to estimate cost insfrastructure: %w", err)
	}

	return res.GetPayload().Data, nil
}
