package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/client/models"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewGetComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [canonical...]",
		Args:  cyargs.RequireArgsOrFlag("component"),
		Short: "Get the state of a current component.",
		Example: `
	# get a component using flags
	cy --org my-org component get -p my-proj -e my-env -c my-comp

	# get multiple components using positional args
	cy --org my-org component get -p my-proj -e my-env comp-a comp-b
`,
		RunE: getComponent,
	}
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	cyout.RegisterModel(cmd, models.Component{})
	return cmd
}

func getComponent(cmd *cobra.Command, args []string) error {
	org, err := cyargs.GetOrg(cmd)
	if err != nil {
		return err
	}

	project, err := cyargs.GetProject(cmd)
	if err != nil {
		return err
	}

	env, err := cyargs.GetEnv(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	// Multi-arg
	if len(args) > 1 {
		results := make([]*models.Component, 0, len(args))
		for _, component := range args {
			c, _, err := m.GetComponent(org, project, env, component)
			if err != nil {
				return cyout.PrintWithOptions(cmd, nil, err, "failed to fetch component '"+component+"'", componentTableOptions)
			}
			results = append(results, c)
		}
		return cyout.PrintWithOptions(cmd, results, nil, "", componentTableOptions)
	}

	var component string
	if len(args) == 1 {
		component = args[0]
	} else {
		component, err = cyargs.GetComponent(cmd)
		if err != nil {
			return err
		}
	}

	c, _, err := m.GetComponent(org, project, env, component)
	return cyout.PrintWithOptions(cmd, c, err, "failed to fetch state of component '"+component+"'", componentTableOptions)
}
