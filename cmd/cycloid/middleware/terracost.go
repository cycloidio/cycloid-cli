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
	params.SetBody(&models.TerraformPlanInput{
		Tfplan: &tfplan,
	})

	res, err := m.api.CostEstimation.CostEstimateTfPlan(params, common.ClientCredentials(&org))
	if err != nil {
		return nil, fmt.Errorf("unable to estimate cost insfrastructure: %w", NewApiError(err))
	}

	return res.GetPayload().Data, nil
}
