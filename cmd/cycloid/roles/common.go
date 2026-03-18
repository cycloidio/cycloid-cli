package roles

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
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

	name, role, err := middleware.NameOrCanonical(&roleName, &roleCan)
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

	output, err := cyargs.GetOutput(cmd)
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	_, _, err = m.GetRole(org, role)
	exists := err == nil
	if err != nil {
		var apiErr *middleware.APIResponseError
		if stderrors.As(err, &apiErr) && apiErr.StatusCode == http.StatusNotFound {
			exists = false
		} else {
			return printer.SmartPrint(p, nil, err, "unable to check if role exists", printer.Options{}, cmd.OutOrStderr())
		}
	}

	if exists && !upsert {
		return printer.SmartPrint(p, nil,
			fmt.Errorf("role %q already exists; use --update or `cy roles update` to update it", role),
			"failed to create role", printer.Options{}, cmd.OutOrStderr())
	}

	if exists {
		outRole, _, err := m.UpdateRole(org, role, &name, &role, &description, rules)
		if err != nil {
			return printer.SmartPrint(p, nil, err, "failed to update role", printer.Options{}, cmd.OutOrStderr())
		}
		return printer.SmartPrint(p, outRole, nil, "", printer.Options{}, cmd.OutOrStdout())
	}

	outRole, _, err := m.CreateRole(org, &name, &role, &description, rules)
	if err != nil {
		return fmt.Errorf("failed to create the role: %w", err)
	}

	return printer.SmartPrint(p, outRole, nil, "", printer.Options{}, cmd.OutOrStdout())
}
