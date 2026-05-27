package environments

import (
	stderrors "errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
)

func envTypeFromCurrent(current *models.Environment) string {
	if current != nil && current.EnvironmentType != nil && current.EnvironmentType.Canonical != nil {
		return *current.EnvironmentType.Canonical
	}
	return ""
}

func ownerFromCurrent(current *models.Environment) string {
	if current != nil && current.Owner != nil && current.Owner.Username != nil {
		return *current.Owner.Username
	}
	return ""
}

func cloudAccountsFromCurrent(current *models.Environment) []string {
	if current == nil {
		return nil
	}
	out := make([]string, 0, len(current.CloudAccounts))
	for _, account := range current.CloudAccounts {
		if account != nil && account.Canonical != nil {
			out = append(out, *account.Canonical)
		}
	}
	return out
}

func variablesFromCurrent(current *models.Environment) []*models.EnvironmentVariableItem {
	if current == nil {
		return nil
	}
	return current.Variables
}

func buildNewEnvironment(cmd *cobra.Command, name, canonical string, current *models.Environment) (*models.NewEnvironment, error) {
	body := &models.NewEnvironment{
		Canonical: canonical,
		Name:      ptr.Ptr(name),
	}

	if cyargs.IsSet(cmd, "type") {
		value, err := cyargs.GetEnvironmentType(cmd)
		if err != nil {
			return nil, err
		}
		body.Type = ptr.Ptr(value)
	} else if current != nil {
		if t := envTypeFromCurrent(current); t != "" {
			body.Type = ptr.Ptr(t)
		}
	}

	if cyargs.IsSet(cmd, "description") {
		description, err := cyargs.GetDescription(cmd)
		if err != nil {
			return nil, err
		}
		body.Description = description
	} else if current != nil {
		body.Description = current.Description
	}

	if cyargs.IsSet(cmd, "owner") {
		owner, err := cyargs.GetEnvironmentOwner(cmd)
		if err != nil {
			return nil, err
		}
		body.Owner = owner
	} else if current != nil {
		body.Owner = ownerFromCurrent(current)
	}

	if cyargs.IsSet(cmd, "cloud-account") {
		accounts, err := cyargs.GetCloudAccountCanonicals(cmd)
		if err != nil {
			return nil, err
		}
		body.CloudAccountCanonicals = accounts
	} else if current != nil {
		body.CloudAccountCanonicals = cloudAccountsFromCurrent(current)
	}

	if cyargs.IsSet(cmd, "variable") || cyargs.IsSet(cmd, "variables-file") {
		variables, err := cyargs.GetEnvironmentVariables(cmd)
		if err != nil {
			return nil, err
		}
		body.Variables = variables
	} else if current != nil {
		body.Variables = variablesFromCurrent(current)
	}

	return body, nil
}

func buildUpdateEnvironment(cmd *cobra.Command, current *models.Environment) (*models.UpdateEnvironment, error) {
	name := current.Name
	if cyargs.IsSet(cmd, "name") {
		value, err := cyargs.GetName(cmd)
		if err != nil {
			return nil, err
		}
		if value != "" {
			name = value
		}
	}

	body := &models.UpdateEnvironment{
		Name: ptr.Ptr(name),
	}

	if cyargs.IsSet(cmd, "type") {
		value, err := cyargs.GetEnvironmentType(cmd)
		if err != nil {
			return nil, err
		}
		body.Type = ptr.Ptr(value)
	} else if t := envTypeFromCurrent(current); t != "" {
		body.Type = ptr.Ptr(t)
	}

	if cyargs.IsSet(cmd, "description") {
		description, err := cyargs.GetDescription(cmd)
		if err != nil {
			return nil, err
		}
		body.Description = description
	} else {
		body.Description = current.Description
	}

	if cyargs.IsSet(cmd, "owner") {
		owner, err := cyargs.GetEnvironmentOwner(cmd)
		if err != nil {
			return nil, err
		}
		body.Owner = owner
	} else {
		body.Owner = ownerFromCurrent(current)
	}

	if cyargs.IsSet(cmd, "cloud-account") {
		accounts, err := cyargs.GetCloudAccountCanonicals(cmd)
		if err != nil {
			return nil, err
		}
		body.CloudAccountCanonicals = accounts
	} else {
		body.CloudAccountCanonicals = cloudAccountsFromCurrent(current)
	}

	if cyargs.IsSet(cmd, "variable") || cyargs.IsSet(cmd, "variables-file") {
		variables, err := cyargs.GetEnvironmentVariables(cmd)
		if err != nil {
			return nil, err
		}
		body.Variables = variables
	}

	return body, nil
}

func createOrUpdateEnvironment(cmd *cobra.Command, m middleware.Middleware, org, canonical string, allowUpdate bool) (*models.Environment, error) {
	name, err := cyargs.GetName(cmd)
	if err != nil {
		return nil, err
	}
	name, canonical, err = middleware.NameOrCanonical(&name, &canonical)
	if err != nil {
		return nil, err
	}

	current, _, err := m.GetOrgEnv(org, canonical)
	if err == nil {
		if !allowUpdate {
			return nil, fmt.Errorf("environment %q already exists (use --update to upsert)", canonical)
		}
		updateBody, err := buildUpdateEnvironment(cmd, current)
		if err != nil {
			return nil, err
		}
		updated, _, err := m.UpdateOrgEnv(org, canonical, updateBody)
		return updated, err
	}

	var apiErr *middleware.APIResponseError
	if !stderrors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
		return nil, err
	}

	createBody, err := buildNewEnvironment(cmd, name, canonical, nil)
	if err != nil {
		return nil, err
	}
	created, _, err := m.CreateOrgEnv(org, createBody)
	return created, err
}
