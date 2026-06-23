package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/apiclient"
	"github.com/cycloidio/cycloid-cli/cmd/common"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
	"github.com/cycloidio/cycloid-cli/gen/models"
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

	if compFlag, err := cyargs.GetComponent(cmd); err != nil {
		return err
	} else if compFlag != "" {
		found := false
		for _, a := range args {
			if a == compFlag {
				found = true
				break
			}
		}
		if !found {
			args = append(args, compFlag)
		}
	}

	api := common.NewAPI()
	m := apiclient.NewMiddleware(api)

	if len(args) == 1 {
		c, _, err := m.GetComponent(org, project, env, args[0])
		return cyout.PrintWithOptions(cmd, c, err, "failed to fetch state of component '"+args[0]+"'", componentTableOptions)
	}

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
