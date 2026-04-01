package middleware

import (
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/client/models"
)

// CostEstimation will consume the backend API endpoint for cost estimation
func (m *middleware) CostEstimation(org string, plan []byte) (*models.CostEstimationResult, *http.Response, error) {
	tfplan := string(plan)
	body := &models.TerraformPlanInput{
		Tfplan: &tfplan,
	}

	var result *models.CostEstimationResult
	resp, err := m.GenericRequest(Request{
		Method:       "POST",
		Organization: &org,
		Route:        []string{"organizations", org, "cost_estimation", "tfplan"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("unable to estimate cost infrastructure: %w", err)
	}
	return result, resp, nil
}
