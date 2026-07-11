package roles

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
)

// runRoleWrite creates or updates a role. upsert: if true, update when role exists; if false, error when it exists.
func runRoleWrite(cmd *cobra.Command, upsert bool) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	roleName, err := cyargs.GetRoleName(cmd)
	if err != nil {
		return err
	}

	roleCan, err := cyargs.GetRoleCanonical(cmd)
	if err != nil {
		return err
	}

	deprecatedCan, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	if roleCan == "" && deprecatedCan != "" {
		roleCan = deprecatedCan
	}

	name, role, err := apiclient.NameOrCanonical(&roleName, &roleCan)
	if err != nil {
		return err
	}

	rulesJSON, err := cyargs.GetRoleRulesJSON(cmd)
	if err != nil {
		return err
	}

	rulesFiles, err := cyargs.GetRoleRulesFiles(cmd)
	if err != nil {
		return err
	}

	rules := append(rulesJSON, rulesFiles...)

	description, err := cyargs.GetDescription(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	_, _, err = m.GetRole(org, role)
	exists := err == nil
	if err != nil {
		var apiErr *apiclient.APIResponseError
		if stderrors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
			exists = false
		} else {
			return cyout.PrintWithOptions(cmd, nil, err, "unable to check if role exists", printer.Options{})
		}
	}

	if exists && !upsert {
		return cyout.PrintWithOptions(cmd, nil,
			fmt.Errorf("role %q already exists; use --update or `cy roles update` to update it", role),
			"failed to create role", printer.Options{})
	}

	if exists {
		outRole, _, err := m.UpdateRole(org, role, &name, &role, &description, rules)
		return cyout.PrintWithOptions(cmd, outRole, err, "failed to update role", printer.Options{})
	}

	outRole, _, err := m.CreateRole(org, &name, &role, &description, rules)
	return cyout.PrintWithOptions(cmd, outRole, err, "failed to create role", printer.Options{})
}
