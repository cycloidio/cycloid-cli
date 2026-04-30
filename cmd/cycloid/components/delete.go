package components

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewDeleteComponentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [canonical...]",
		Args:  cyargs.RequireArgsOrFlag("component"),
		Short: "Delete a component",
		Example: `
	# delete a component using flags
	cy --org my-org component delete -p my-proj -e my-env -c my-comp

	# delete multiple components using positional args
	cy --org my-org component delete -p my-proj -e my-env comp-a comp-b
`,
		RunE: deleteComponent,
	}
	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cyargs.AddComponentFlag(cmd)
	return cmd
}

func deleteComponent(cmd *cobra.Command, args []string) error {
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
	m := middleware.NewMiddleware(api)

	components := args

	for _, component := range components {
		_, err = m.DeleteComponent(org, project, env, component)
		if err != nil {
			return cyout.Print(cmd, nil, err, "failed to delete component '"+component+"'")
		}
	}
	return nil
}
