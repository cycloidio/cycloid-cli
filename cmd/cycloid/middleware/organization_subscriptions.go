package middleware

import (
	"fmt"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
)

type SubscriptionPlan = string

const (
	FreeTrial    = "free_trial"
	PlatformTeam = "platform_team"
)

var (
	AvailableSubscriptionPlans = []string{
		FreeTrial,
		PlatformTeam,
	}
)

func (m *middleware) CreateOrUpdateSubscription(org string, plan *SubscriptionPlan, expiresAt time.Time, membersCount uint64, overwrite bool) (*models.Subscription, error) {
	var subscription *models.Subscription

	var body = map[string]any{
		"expires_at":    expiresAt.Format(time.RFC3339),
		"members_count": membersCount,
		"overwrite":     overwrite,
	}

	if plan != nil {
		body["plan_canonical"] = *plan
	}

	_, err := m.GenericRequest("PUT", &org, nil, nil, body, subscription, "organizations", org, "subscriptions")
	if err != nil {
		return nil, fmt.Errorf("failed to update subscription: %w", err)
	}

	return subscription, nil
}
