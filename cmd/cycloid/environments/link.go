package environments

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/internal/cyargs"
	"github.com/cycloidio/cycloid-cli/internal/cyout"
)

func NewLinkCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "link",
		Args:    cobra.NoArgs,
		Short:   "link an environment to a project",
		Example: `cy --org my-org environment link --project my-proj --env my-env`,
		RunE:    link,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cmd.MarkFlagRequired("project")
	cmd.MarkFlagRequired("env")
	return cmd
}

func link(cmd *cobra.Command, args []string) error {
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

	_, err = m.LinkEnvToProject(org, project, env)
	if err != nil {
		return cyout.Print(cmd, nil, err, "failed to link environment to project")
	}
	return nil
}

func NewUnlinkCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "unlink",
		Args:    cobra.NoArgs,
		Short:   "unlink an environment from a project",
		Example: `cy --org my-org environment unlink --project my-proj --env my-env`,
		RunE:    unlink,
	}

	cyargs.AddProjectFlag(cmd)
	cyargs.AddEnvFlag(cmd)
	cmd.MarkFlagRequired("project")
	cmd.MarkFlagRequired("env")
	cyargs.AddDeleteFlags(cmd)
	return cmd
}

func unlink(cmd *cobra.Command, args []string) error {
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

	force, skipHooks, ignoreConfigFilesErr, err := cyargs.GetDeleteFlags(cmd)
	if err != nil {
		return err
	}

	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	opts := middleware.DeleteOptions{Force: force, SkipHooks: skipHooks, IgnoreConfigFilesErr: ignoreConfigFilesErr}
	_, err = m.UnlinkEnvFromProject(org, project, env, opts)
	if err != nil {
		return cyout.Print(cmd, nil, err, "failed to unlink environment from project")
	}
	return nil
}
