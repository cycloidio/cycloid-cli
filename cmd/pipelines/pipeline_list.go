package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
)

func NewPipelineListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all pipelines of the organization",
		Args:  cobra.NoArgs,
		RunE:  list,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddPipelineStatuses(cmd)
	cyargs.AddPipeline(cmd)
	cyout.RegisterModel(cmd, models.Pipeline{})
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	// Those args are filters and optional
	project, _ := cyargs.GetProject(cmd)
	env, _ := cyargs.GetEnv(cmd)
	statuses, _ := cyargs.GetPipelineStatuses(cmd)
	pipelineName, _ := cyargs.GetPipeline(cmd)

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	pps, _, err := m.GetOrgPipelines(org, &pipelineName, &project, &env, statuses)
	return cyout.PrintWithOptions(cmd, pps, err, "unable to list pipelines", pipelineTableOptions)
}
