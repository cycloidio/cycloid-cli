package environmenttypes

import (
	stderrors "errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/internal/ptr"
	"github.com/cycloidio/cycloid-cli/printer"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create an environment type",
		Example: `cy --org my-org environment-type create --environment-type qa --environment-type-name QA --color staging`,
		RunE:    create,
		Args:    cobra.NoArgs,
	}

	cmd.MarkFlagsOneRequired(
		cyargs.AddEnvironmentTypeNameFlag(cmd),
		cyargs.AddEnvironmentTypeCanonicalFlag(cmd),
	)
	cmd.MarkFlagRequired(cyargs.AddColorFlag(cmd))
	cyargs.AddUpdateFlag(cmd, "update the environment type if it already exists")
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cyargs.GetEnvironmentTypeName(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetEnvironmentTypeCanonical(cmd)
	if err != nil {
		return err
	}
	name, canonical, err = middleware.NameOrCanonical(&name, &canonical)
	if err != nil {
		return err
	}

	color, err := cyargs.GetColor(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	if cyargs.GetUpdate(cmd) {
		_, _, getErr := m.GetEnvironmentType(org, canonical)
		if getErr == nil {
			return updateEnvironmentType(cmd, args)
		}
		var apiErr *middleware.APIResponseError
		if !stderrors.As(getErr, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
			return cyout.PrintWithOptions(cmd, nil, getErr, "failed to check existing environment type", printer.Options{})
		}
	}

	body := &models.NewEnvironmentType{
		Canonical: canonical,
		Name:      ptr.Ptr(name),
		Color:     ptr.Ptr(color),
	}
	result, _, err := m.CreateEnvironmentType(org, body)
	return cyout.PrintWithOptions(cmd, result, err, "failed to create environment type", environmentTypeTableOptions)
}

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Short:   "Update an environment type",
		RunE:    updateEnvironmentType,
		Args:    cobra.NoArgs,
	}

	cyargs.AddEnvironmentTypeCanonicalFlag(cmd)
	cmd.MarkFlagRequired("environment-type")
	cyargs.AddEnvironmentTypeNameFlag(cmd)
	cyargs.AddColorFlag(cmd)
	return cmd
}

func updateEnvironmentType(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	canonical, err := cyargs.GetEnvironmentTypeCanonical(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	current, _, err := m.GetEnvironmentType(org, canonical)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "environment type not found", environmentTypeTableOptions)
	}

	name := ptrValue(current.Name)
	if cyargs.IsSet(cmd, "environment-type-name") {
		value, err := cyargs.GetEnvironmentTypeName(cmd)
		if err != nil {
			return err
		}
		if value != "" {
			name = value
		}
	}

	color := ptrValue(current.Color)
	if cyargs.IsSet(cmd, "color") {
		value, err := cyargs.GetColor(cmd)
		if err != nil {
			return err
		}
		color = value
	}

	body := &models.UpdateEnvironmentType{
		Name:  ptr.Ptr(name),
		Color: ptr.Ptr(color),
	}
	result, _, err := m.UpdateEnvironmentType(org, canonical, body)
	return cyout.PrintWithOptions(cmd, result, err, "failed to update environment type", environmentTypeTableOptions)
}

func ptrValue(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
