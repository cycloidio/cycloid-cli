package environmenttypes

import (
	stderrors "errors"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
	"github.com/cycloidio/cycloid-cli/utils/ptr"
)

func NewCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an environment type",
		Example: `cy --org my-org environment-type create --environment-type qa --environment-type-name QA --color staging
cy --org my-org environment-type create --environment-type prod --color production --label-selector-enforcement hard --label-selector-requirement "env-type:in:production,staging"`,
		RunE: create,
		Args: cobra.NoArgs,
	}

	cmd.MarkFlagsOneRequired(
		cyargs.AddEnvironmentTypeNameFlag(cmd),
		cyargs.AddEnvironmentTypeCanonicalFlag(cmd),
	)
	_ = cmd.MarkFlagRequired(cyargs.AddColorFlag(cmd))
	cyargs.AddUpdateFlag(cmd, "update the environment type if it already exists")
	cyargs.AddLabelSelectorCreateFlags(cmd)
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
	name, canonical, err = apiclient.NameOrCanonical(&name, &canonical)
	if err != nil {
		return err
	}

	color, err := cyargs.GetColor(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	if cyargs.GetUpdate(cmd) {
		_, _, getErr := m.GetEnvironmentType(org, canonical)
		if getErr == nil {
			return updateEnvironmentType(cmd, args)
		}
		var apiErr *apiclient.APIResponseError
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
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "failed to create environment type", environmentTypeTableOptions)
	}

	// Set label selector if flags are provided
	if cyargs.HasLabelSelectorFlags(cmd) {
		labelSelector, err := cyargs.GetLabelSelector(cmd)
		if err != nil {
			return err
		}
		result, _, err = m.SetEnvironmentTypeLabelSelector(org, canonical, labelSelector)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "failed to set label selector", environmentTypeTableOptions)
		}
	}

	return cyout.PrintWithOptions(cmd, result, nil, "", environmentTypeTableOptions)
}

func NewUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an environment type",
		Example: `cy --org my-org environment-type update --environment-type staging --color staging
cy --org my-org environment-type update --environment-type prod --label-selector-enforcement hard --label-selector-requirement "env-type:in:production"
cy --org my-org environment-type update --environment-type prod --delete-label-selector`,
		RunE: updateEnvironmentType,
		Args: cobra.NoArgs,
	}

	cyargs.AddEnvironmentTypeCanonicalFlag(cmd)
	_ = cmd.MarkFlagRequired("environment-type")
	cyargs.AddEnvironmentTypeNameFlag(cmd)
	cyargs.AddColorFlag(cmd)
	cyargs.AddLabelSelectorFlags(cmd)
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
	m := apiclient.NewAPIClient(api)

	current, _, err := m.GetEnvironmentType(org, canonical)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "environment type not found", environmentTypeTableOptions)
	}

	name := ptr.Value(current.Name)
	if cyargs.IsSet(cmd, "environment-type-name") {
		value, err := cyargs.GetEnvironmentTypeName(cmd)
		if err != nil {
			return err
		}
		if value != "" {
			name = value
		}
	}

	color := ptr.Value(current.Color)
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
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "failed to update environment type", environmentTypeTableOptions)
	}

	// Handle label selector: delete or set
	if cyargs.GetDeleteLabelSelector(cmd) {
		_, err = m.DeleteEnvironmentTypeLabelSelector(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "failed to delete label selector", environmentTypeTableOptions)
		}
		// Re-read to get updated state without label_selector
		result, _, err = m.GetEnvironmentType(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "failed to read environment type after deleting label selector", environmentTypeTableOptions)
		}
	} else if cyargs.HasLabelSelectorFlags(cmd) {
		labelSelector, err := cyargs.GetLabelSelector(cmd)
		if err != nil {
			return err
		}
		result, _, err = m.SetEnvironmentTypeLabelSelector(org, canonical, labelSelector)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, "failed to set label selector", environmentTypeTableOptions)
		}
	}

	return cyout.PrintWithOptions(cmd, result, nil, "", environmentTypeTableOptions)
}
