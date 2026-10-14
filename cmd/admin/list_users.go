package admin

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

var userTableOptions = printer.Options{
	Columns:    []string{"Username", "Email", "Verified", "CreatedAt", "LastLoginAt", "Organizations"},
	Identifier: "Username",
	Transform: func(obj interface{}) map[string]string {
		u, ok := obj.(*models.AdminUser)
		if !ok || u == nil {
			return map[string]string{}
		}
		orgs := make([]string, 0, len(u.Organizations))
		for _, o := range u.Organizations {
			orgs = append(orgs, ptr.Value(o.Canonical)+" ("+ptr.Value(o.Role)+")")
		}
		return map[string]string{
			"Username":      ptr.Value(u.Username),
			"Email":         ptr.Value(u.Email).String(),
			"Verified":      strconv.FormatBool(ptr.Value(u.EmailVerified)),
			"CreatedAt":     formatMillis(ptr.Value(u.CreatedAt)),
			"LastLoginAt":   formatMillis(ptr.Value(u.LastLoginAt)),
			"Organizations": strings.Join(orgs, ", "),
		}
	},
}

// formatMillis renders an API timestamp, blank when absent
func formatMillis(ms uint64) string {
	if ms == 0 {
		return ""
	}
	return time.UnixMilli(int64(ms)).UTC().Format(time.RFC3339)
}

func NewListUsersCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		Short:   "List the users of the instance, including the ones without an organization",
		Long: `List the users of the instance, including the ones without an organization, with the
organizations they belong to and their role there. --org has to be the root organization, and the
caller one of its admins.`,
		Example: `
	# List all the users
	cy --org my-root-org admin users list

	# List the users without an organization, created since October 1st, whose email contains example.com
	cy --org my-root-org admin users list --without-organization --created-after 2026-10-01 --email example.com
`,
		RunE: listUsers,
	}
	cyout.RegisterModel(cmd, models.AdminUser{})

	cmd.Flags().String("email", "", "keep the users whose primary email contains this value")
	cmd.Flags().Bool("without-organization", false, "keep only the users who belong to no organization")
	cmd.Flags().String("created-after", "", "keep the users created after this date or RFC3339 time")

	return cmd
}

func listUsers(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	filters, err := userFilters(cmd)
	if err != nil {
		return err
	}

	users, _, err := m.AdminListUsers(org, filters...)
	return cyout.PrintWithOptions(cmd, users, err, "unable to list users", userTableOptions)
}

func userFilters(cmd *cobra.Command) ([]apiclient.LHSFilter, error) {
	var filters []apiclient.LHSFilter

	email, err := cmd.Flags().GetString("email")
	if err != nil {
		return nil, err
	}
	if email != "" {
		// Quoted: a pasted address has '.' and often '+', which a regex would not take literally
		filters = append(filters, apiclient.LHSFilter{Attribute: "user_email", Condition: "rlike", Value: regexp.QuoteMeta(email)})
	}

	withoutOrg, err := cmd.Flags().GetBool("without-organization")
	if err != nil {
		return nil, err
	}
	if withoutOrg {
		filters = append(filters, apiclient.LHSFilter{Attribute: "user_has_organization", Condition: "eq", Value: "false"})
	}

	createdAfter, err := cmd.Flags().GetString("created-after")
	if err != nil {
		return nil, err
	}
	if createdAfter != "" {
		t, err := time.Parse(time.RFC3339, createdAfter)
		if err != nil {
			t, err = time.Parse(time.DateOnly, createdAfter)
		}
		if err != nil {
			return nil, fmt.Errorf("invalid --created-after %q: expected YYYY-MM-DD or an RFC3339 time", createdAfter)
		}
		filters = append(filters, apiclient.LHSFilter{
			Attribute: "user_created_at",
			Condition: "gt",
			Value:     strconv.FormatInt(t.UnixMilli(), 10),
		})
	}

	return filters, nil
}
