package inventory

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
)

// addScopingFlags registers the shared project/environment/component flags
// (from cyargs) used to scope inventory queries.
func addScopingFlags(cmd *cobra.Command) {
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
}

// appendScopingFilters reads the optional --project / --env / --component flags
// and appends the matching LHS filters when they are set.
func appendScopingFilters(cmd *cobra.Command, filters []apiclient.LHSFilter) ([]apiclient.LHSFilter, error) {
	project, err := cyargs.GetProjectOrEmpty(cmd)
	if err != nil {
		return nil, err
	}
	env, err := cyargs.GetEnvOrEmpty(cmd)
	if err != nil {
		return nil, err
	}
	component, err := cyargs.GetComponentOrEmpty(cmd)
	if err != nil {
		return nil, err
	}

	if project != "" {
		filters = append(filters, apiclient.LHSFilter{Attribute: "project_canonical", Condition: "eq", Value: project})
	}
	if env != "" {
		filters = append(filters, apiclient.LHSFilter{Attribute: "environment_canonical", Condition: "eq", Value: env})
	}
	if component != "" {
		filters = append(filters, apiclient.LHSFilter{Attribute: "component_canonical", Condition: "eq", Value: component})
	}
	return filters, nil
}
