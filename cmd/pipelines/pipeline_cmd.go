package pipelines

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pipeline",
		Aliases: []string{
			"pp",
			"pipelines",
		},
		Short: "Manage the pipelines",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewBuildsCommand(),
		NewJobsCommand(),
		NewPipelineClearTaskCacheCommand(),
		NewPipelineDiffCommand(),
		NewPipelineGetCommand(),
		NewPipelineLastUsedCommand(),
		NewPipelineListCommand(),
		NewPipelinePauseCommand(),
		NewPipelineSyncedCommand(),
		NewPipelineUnpauseCommand(),
		NewPipelineUpdateCommand(),
	)

	return cmd
}

var pipelineTableOptions = printer.Options{
	Columns:    []string{"Name", "Status", "Paused", "Project", "Environment", "Component"},
	Identifier: "Name",
	Transform: func(obj interface{}) map[string]string {
		pp, ok := obj.(*models.Pipeline)
		if !ok {
			return map[string]string{}
		}
		name := ""
		if pp.Name != nil {
			name = *pp.Name
		}
		paused := ""
		if pp.Paused != nil {
			paused = strconv.FormatBool(*pp.Paused)
		}
		project := ""
		if pp.Project != nil && pp.Project.Canonical != nil {
			project = *pp.Project.Canonical
		}
		env := ""
		if pp.Environment != nil && pp.Environment.Canonical != nil {
			env = *pp.Environment.Canonical
		}
		comp := ""
		if pp.Component != nil && pp.Component.Canonical != nil {
			comp = *pp.Component.Canonical
		}
		return map[string]string{
			"Name":        name,
			"Status":      pp.Status,
			"Paused":      paused,
			"Project":     project,
			"Environment": env,
			"Component":   comp,
		}
	},
}
