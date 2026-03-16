package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
)

type SubscriptionPlan = string

const (
	FreeTrial    = "free_trial"
	PlatformTeam = "platform_teams"
)

var (
	AvailableSubscriptionPlans = []string{
		FreeTrial,
		PlatformTeam,
	}
)

type SubscriptionRequest struct {
	ExpiresAt     string `json:"expires_at"`
	MembersCount  uint64 `json:"members_count"`
	Overwrite     bool   `json:"overwrite"`
	PlanCanonical string `json:"plan_canonical"`
}

func (m *middleware) CreateOrUpdateSubscription(org string, plan SubscriptionPlan, expiresAt time.Time, membersCount uint64, overwrite bool) (*models.Subscription, *http.Response, error) {
	body := SubscriptionRequest{
		ExpiresAt:     expiresAt.Format(time.RFC3339),
		MembersCount:  membersCount,
		Overwrite:     overwrite,
		PlanCanonical: plan,
	}

	var result *models.Subscription
	resp, err := m.GenericRequest(Request{
		Method:       "PUT",
		Organization: &org,
		Route:        []string{"organizations", org, "subscriptions"},
		Body:         body,
	}, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update subscription: %w", err)
	}
	return result, resp, nil
}
