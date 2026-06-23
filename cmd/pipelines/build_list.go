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

func NewBuildListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "list builds of a pipeline",
		Example: `cy --org my-org pp list-builds --project my-project --env env --component component --job my-job`,
		RunE:    listBuilds,
		Args:    cobra.NoArgs,
	}

	cyargs.AddCyContext(cmd)
	cyargs.AddPipeline(cmd)
	cmd.MarkFlagRequired(cyargs.AddPipelineJob(cmd))
	cyout.RegisterModel(cmd, models.Build{})
	return cmd
}

func listBuilds(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	pipeline, err := cyargs.GetPipeline(cmd)
	if err != nil {
		return err
	}

	// Those args are optional filters, we don't care about err
	_, project, env, component, _ := cyargs.GetCyContext(cmd)
	job, _ := cyargs.GetPipelineJob(cmd)

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	builds, _, err := m.GetBuilds(org, project, env, component, pipeline, job)
	errMsg := fmt.Sprintf("failed to get builds of pipeline '%s', in project '%s', in env '%s', in component '%s'",
		pipeline, project, env, component)
	return cyout.PrintWithOptions(cmd, builds, err, errMsg, buildTableOptions)
}
