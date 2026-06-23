package pipelines

import (
	"fmt"
	"slices"
	"time"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewPipelineLastUsedCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "last-used",
		Short:   "output the last trigger date of all pipelines in the selected organization.",
		Example: `cy pipelines last-used --since-days 10`,
		RunE:    lastUsed,
		Args:    cobra.NoArgs,
	}

	cmd.PersistentFlags().Int64P("since-days", "s", 0, "filter pipelines that didn't ran since x days.")
	return cmd
}

type LastUsedPipeline struct {
	PipelineName string
	Project      string
	Environment  string
	Component    string
	LastUsed     LastUsed
}

type LastUsed struct {
	Timestamp   uint64
	DateRFC3339 string
}

func lastUsed(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}
	sinceDays, err := cmd.Flags().GetInt64("since-days")
	if err != nil {
		return err
	}

	result := []LastUsedPipeline{}
	maxTimestamp := uint64(0)
	pps, _, err := m.GetOrgPipelines(org, nil, nil, nil, nil)
	if err != nil {
		return cyout.PrintWithOptions(cmd, nil, err, "failed to list pipelines", printer.Options{})
	}

	for _, pipeline := range pps {
		lastestJob := slices.MaxFunc(pipeline.Jobs, func(a, b *models.Job) int {
			if a.FinishedBuild == nil && b.FinishedBuild == nil {
				return int(maxTimestamp)
			} else if a.FinishedBuild == nil && b.FinishedBuild != nil {
				return int(b.FinishedBuild.StartTime)
			} else if a.FinishedBuild != nil && b.FinishedBuild == nil {
				return int(a.FinishedBuild.StartTime)
			}

			if a.FinishedBuild.StartTime > b.FinishedBuild.StartTime {
				return int(a.FinishedBuild.StartTime)
			}

			return int(b.FinishedBuild.StartTime)
		})

		if pipeline.Name == nil {
			return fmt.Errorf("missing pipeline name in:\n%v", pipeline)
		}

		if pipeline.Project == nil {
			return fmt.Errorf("missing pipeline project in:\n%v", pipeline)
		}

		if pipeline.Environment == nil {
			return fmt.Errorf("missing pipeline environment in:\n%v", pipeline)
		}

		if pipeline.Component == nil {
			return fmt.Errorf("missing pipeline component in:\n%v", pipeline)
		}

		lastUsedTimestamp := maxTimestamp
		if lastestJob.FinishedBuild != nil {
			lastUsedTimestamp = lastestJob.FinishedBuild.StartTime
		}

		timestamp := time.Unix(int64(lastUsedTimestamp), 0)

		if timestamp.Unix() < (time.Now().Unix() - (sinceDays * 60 * 60 * 24)) {
			result = append(result, LastUsedPipeline{
				PipelineName: *pipeline.Name,
				Project:      *pipeline.Project.Canonical,
				Environment:  *pipeline.Environment.Canonical,
				Component:    *pipeline.Component.Canonical,
				LastUsed: LastUsed{
					Timestamp:   lastUsedTimestamp,
					DateRFC3339: timestamp.Format(time.RFC3339),
				},
			})
		}
	}

	return cyout.PrintWithOptions(cmd, result, nil, "unable to list pipelines", printer.Options{})
}
