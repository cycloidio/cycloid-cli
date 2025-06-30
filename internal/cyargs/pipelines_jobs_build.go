package cyargs

import (
	"fmt"
	"strconv"
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

func AddPipeline(cmd *cobra.Command) string {
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

func GetPipeline(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("pipeline")
}

func AddPipelineConfig(cmd *cobra.Command) string {
	flagName := "pipeline-config"
	cmd.Flags().StringP(flagName, "C", "pipeline.yml", "path to the pipeline config file.")
	cmd.MarkFlagFilename("pipeline-config", "yml", "yaml")
	return flagName
}

func GetPipelineConfig(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("pipeline-config")
}

func AddPipelineVars(cmd *cobra.Command) string {
	flagName := "pipeline-vars"
	cmd.Flags().StringP(flagName, "V", "variables.sample.yml", "path to the pipeline variable file.")
	cmd.MarkFlagFilename("pipeline-vars", "yml", "yaml")
	return flagName
}

func GetPipelineVars(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("pipeline-vars")
}

func AddPipelineJob(cmd *cobra.Command) string {
	flagName := "job"
	cmd.Flags().StringP(flagName, "j", "", "job canonical.")
	cmd.RegisterFlagCompletionFunc("job", func(cmd *cobra.Command, args []cobra.Completion, toComplete string) ([]string, cobra.ShellCompDirective) {
		org, err := GetOrg(cmd)
		if err != nil {
			return cobra.AppendActiveHelp(nil, "missing org parameter for completion"),
				cobra.ShellCompDirectiveNoFileComp
		}

		project, _ := GetProject(cmd)
		env, _ := GetEnv(cmd)
		component, _ := GetEnv(cmd)
		pipeline, _ := GetPipeline(cmd)

		api := common.NewAPI()
		m := middleware.NewMiddleware(api)
		jobs, err := m.GetJobs(org, project, env, component, pipeline)
		if err != nil {
			return cobra.AppendActiveHelp(nil, fmt.Sprintf(
					"failed to fetch job list for completion with context org: %s, project: %s, env: %s, component: %s, pipeline: %s, err: %s",
					org, project, env, component, pipeline, err,
				)),
				cobra.ShellCompDirectiveNoFileComp
		}

		jobNames := make([]cobra.Completion, len(jobs))
		for index, job := range jobs {
			if job.Name != nil && strings.HasPrefix(*job.Name, toComplete) {
				jobNames[index] = cobra.CompletionWithDesc(*job.Name, strconv.Itoa(int(*job.ID)))
			}
		}

		return jobNames, cobra.ShellCompDirectiveNoFileComp
	})

	return flagName
}

func GetPipelineJob(cmd *cobra.Command) (string, error) {
	return cmd.Flags().GetString("job")
}
