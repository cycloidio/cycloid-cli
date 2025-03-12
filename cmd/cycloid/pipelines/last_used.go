package pipelines

import (
	"slices"
	"time"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewLastUsedCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "last-used",
		Short: "output the last use of this pipeline as a timestamp.",
		RunE:  lastUsed,
	}

	cmd.PersistentFlags().Int64P("before", "b", 0, "filter pipelines that didn't ran since x days.")

	return cmd
}

type LastUsedPipeline struct {
	PipelineName string
	Project      string
	Environment  string
	LastUsed     LastUsed
}

type LastUsed struct {
	Timestamp   uint64
	DateRFC3339 string
}

func lastUsed(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}
	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return err
	}

	before, err := cmd.Flags().GetInt64("before")
	if err != nil {
		return err
	}

	if output == "table" {
		output = "json"
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	var result = []LastUsedPipeline{}
	maxTimestamp := uint64(0)
	pps, err := m.ListPipelines(org)
	if err != nil {
		return printer.SmartPrint(p, nil, err, "failed to list pipelines", printer.Options{}, cmd.ErrOrStderr())
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
			return errors.Errorf("Missing pipeline name in:\n%v", pipeline)
		}

		if pipeline.Project == nil {
			return errors.Errorf("Missing pipeline project in:\n%v", pipeline)
		}

		if pipeline.Environment == nil {
			return errors.Errorf("Missing pipeline environment in:\n%v", pipeline)
		}

		lastUsedTimestamp := maxTimestamp
		if lastestJob.FinishedBuild != nil {
			lastUsedTimestamp = lastestJob.FinishedBuild.StartTime
		}

		timestamp := time.Unix(int64(lastUsedTimestamp), 0)

		if timestamp.Unix() < (time.Now().Unix() - (before * 60 * 60 * 24)) {
			result = append(result, LastUsedPipeline{
				PipelineName: *pipeline.Name,
				Project:      *pipeline.Project.Canonical,
				Environment:  *pipeline.Environment.Canonical,
				LastUsed: LastUsed{
					Timestamp:   lastUsedTimestamp,
					DateRFC3339: timestamp.Format(time.RFC3339),
				},
			})
		}
	}

	return printer.SmartPrint(p, result, nil, "unable to list pipelines", printer.Options{}, cmd.OutOrStdout())
}
