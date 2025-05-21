package cyargs

import (
	"fmt"
	"strings"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/spf13/cobra"
)

var (
	validPipelineStatuses = []string{
		"aborted",
		"errored",
		"failed",
		"paused",
		"pending",
		"started",
		"succeeded",
	}
)

func AddPipelineStatuses(cmd *cobra.Command) string {
	flagName := "statuses"
	cmd.Flags().StringSliceP(
		flagName, "S", validPipelineStatuses,
		fmt.Sprintf("filter per status for pipelines. valid options are [%s]", strings.Join(validPipelineStatuses, ", ")),
	)

	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		var completions []cobra.Completion
		for _, status := range validPipelineStatuses {
			if strings.HasPrefix(status, toComplete) {
				completions = append(completions, status)
			}
		}

		return completions, cobra.ShellCompDirectiveNoFileComp
	})
	return flagName
}

func GetPipelineStatuses(cmd *cobra.Command) ([]string, error) {
	statuses, err := cmd.Flags().GetStringSlice("statuses")
	if err != nil {
		return []string{}, err
	}

	return statuses, nil
}

func AddPipelineName(cmd *cobra.Command) string {
	flagName := "pipeline"
	cmd.Flags().String(flagName, "", "the name of a pipeline")
	cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
				cobra.ShellCompDirectiveNoFileComp
		}

		project, _ := GetProject(cmd)
		env, _ := GetEnv(cmd)
		statuses, _ := GetPipelineStatuses(cmd)

		api := common.NewAPI()
		m := middleware.NewMiddleware(api)
		pipelines, err := m.GetOrgPipelines(org, &toComplete, &project, &env, statuses)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "failed to fetch pipeline list for completion in org '"+org+"': "+err.Error()),
				cobra.ShellCompDirectiveNoFileComp
		}

		var names = make([]cobra.Completion, len(pipelines))
		for index, pipeline := range pipelines {
			if pipeline.Name != nil && strings.HasPrefix(*pipeline.Name, toComplete) {
				names[index] = cobra.CompletionWithDesc(*pipeline.Name,
					strings.Join([]string{*pipeline.Project.Name, pipeline.Environment.Name, *pipeline.Component.Name}, " > "),
				)
			}
		}

		return names, cobra.ShellCompDirectiveNoFileComp
	})

	return flagName
}

func GetPipelineName(cmd *cobra.Command) (string, error) {
	name, err := cmd.Flags().GetString("pipeline")
	if err != nil {
		return "", err
	}

	return name, nil
}
