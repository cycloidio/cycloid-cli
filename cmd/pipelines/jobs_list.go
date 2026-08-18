package pipelines

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewJobsListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list a pipeline's jobs",
		Example: `cy --org my-org pp job list --project my-project --env env --component component --pipeline pipeline -o json`,
		RunE:    listJobs,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cyout.RegisterModel(cmd, models.Job{})
	return cmd
}

func listJobs(cmd *cobra.Command, args []string) error {
	org, project, env, component, err := cyargs.GetCyContext(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := apiclient.NewAPIClient(api)

	jobs, _, err := m.GetJobs(org, project, env, component, pipeline)
	errMsg := fmt.Sprintf("failed to fetch jobs for pipeline %q", pipeline)
	return cyout.PrintWithOptions(cmd, jobs, err, errMsg, jobTableOptions)
}
