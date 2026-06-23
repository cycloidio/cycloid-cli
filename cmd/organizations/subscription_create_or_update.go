package organizations

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateOrUpdateSubscriptionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Add a subscription to an organization, requires an Admin API Key from a parent org.",
		Long:    "Related docs: https://docs.cycloid.io/reference/organizations/concepts/licencing",
		Aliases: []string{"update"},
		Args:    cobra.NoArgs,
		RunE:    createOrUpdateSubscription,
	}

	cmd.Flags().BoolP("update", "u", false, "allow update if a subscription exists")
	cmd.Flags().Uint64P("member-count", "c", 5, "number of member to attribute to this organization, default to 5.")
	defaultDate := time.Now().AddDate(0, 3, 0)
	cmd.Flags().TimeP("expires-at", "t", defaultDate, []string{time.RFC3339}, "Add an expiration time for the subscription, default in three month ("+defaultDate.Format(time.RFC3339)+")")
	cmd.MarkFlagRequired("expire-at")
	cmd.Flags().StringP("plan", "p", apiclient.PlatformTeam, "Select a plan, default to `platform_teams`")
	cmd.RegisterFlagCompletionFunc("plan", func(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
		completions := []cobra.Completion{}
		for _, p := range apiclient.AvailableSubscriptionPlans {
			if strings.HasPrefix(p, toComplete) {
				completions = append(completions, p)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	})

	return cmd
}

func createOrUpdateSubscription(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	update, err := cmd.Flags().GetBool("update")
	if err != nil {
		return err
	}
	if cmd.Name() == "update" {
		update = true
	}

	memberCount, err := cmd.Flags().GetUint64("member-count")
	if err != nil {
		return err
	}

	expiresAt, err := cmd.Flags().GetTime("expires-at")
	if err != nil {
		return err
	}

	plan, err := cmd.Flags().GetString("plan")
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	subscription, _, err := m.CreateOrUpdateSubscription(org, plan, expiresAt, memberCount, update)
	if err != nil {
		return fmt.Errorf("failed to update subscription for org %q: %w", org, err)
	}

	return cyout.PrintWithOptions(cmd, subscription, nil, "", printer.Options{})
}
