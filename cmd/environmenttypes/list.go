package environmenttypes

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List environment types",
		RunE:  list,
		Args:  cobra.NoArgs,
	}

	cyout.RegisterModel(cmd, models.EnvironmentType{})
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	types, _, err := m.ListEnvironmentTypes(org)
	return cyout.PrintWithOptions(cmd, types, err, "failed to list environment types", environmentTypeTableOptions)
}

func NewGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get [environment_type_canonical...]",
		Short:             "Get environment type details",
		RunE:              get,
		Args:              cyargs.RequireArgsOrFlag("environment-type"),
		ValidArgsFunction: cyargs.CompleteEnvironmentTypeCanonical,
	}

	cyargs.AddEnvironmentTypeCanonicalFlag(cmd)
	cyout.RegisterModel(cmd, models.EnvironmentType{})
	return cmd
}

func get(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if flag, err := cyargs.GetEnvironmentTypeCanonical(cmd); err != nil {
		return err
	} else if flag != "" {
		args = appendUniqueArg(args, flag)
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	if len(args) == 1 {
		envType, _, err := m.GetEnvironmentType(org, args[0])
		return cyout.PrintWithOptions(cmd, envType, err, "failed to get environment type", environmentTypeTableOptions)
	}

	results := make([]*models.EnvironmentType, 0, len(args))
	for _, canonical := range args {
		envType, _, err := m.GetEnvironmentType(org, canonical)
		if err != nil {
			return cyout.PrintWithOptions(cmd, nil, err, fmt.Sprintf("failed to get environment type %q", canonical), environmentTypeTableOptions)
		}
		results = append(results, envType)
	}
	return cyout.PrintWithOptions(cmd, results, nil, "", environmentTypeTableOptions)
}

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "delete [environment_type_canonical...]",
		Short:             "Delete environment types",
		RunE:              deleteEnvironmentType,
		Args:              cyargs.RequireArgsOrFlag("environment-type"),
		ValidArgsFunction: cyargs.CompleteEnvironmentTypeCanonical,
	}

	cyargs.AddEnvironmentTypeCanonicalFlag(cmd)
	return cmd
}

func deleteEnvironmentType(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	if flag, err := cyargs.GetEnvironmentTypeCanonical(cmd); err != nil {
		return err
	} else if flag != "" {
		args = appendUniqueArg(args, flag)
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	for _, canonical := range args {
		if _, err := m.DeleteEnvironmentType(org, canonical); err != nil {
			return cyout.Print(cmd, nil, err, "failed to delete environment type "+canonical)
		}
	}
	return nil
}

func appendUniqueArg(args []string, value string) []string {
	for _, existing := range args {
		if existing == value {
			return args
		}
	}
	return append(args, value)
}
