package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewJobsGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "get a pipeline's job",
		Example: `cy --org my-org pp job get --project my-project --env env --component component --pipeline pipeline --job my-job -o json`,
		RunE:    getJob,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyargs.AddPipelineJob(cmd)
	cyout.RegisterModel(cmd, models.Job{})
	return cmd
}

func getJob(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	job, err := cyargs.GetPipelineJob(cmd)
	if err != nil {
		return err
	}

	outJob, _, err := m.GetJob(org, project, env, component, pipeline, job)
	return cyout.PrintWithOptions(cmd, outJob, err, "failed to fetch job: "+job, jobTableOptions)
}
