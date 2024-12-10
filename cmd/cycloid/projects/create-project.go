package projects

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/middleware"
	"github.com/cycloidio/cycloid-cli/printer"
	"github.com/cycloidio/cycloid-cli/printer/factory"
)

func NewCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "create a project",
		Example: `
	# create a project
	cy --org my-org project create \
		--name my-project \
		--description "an awesome project" \
		--owner "username" \
		--stack-ref my-stack-ref \
		--config-repo config-repo-canonical
`,
		RunE:    create,
		PreRunE: internal.CheckAPIAndCLIVersion,
	}
	common.RequiredFlag(WithFlagName, cmd)
	common.RequiredFlag(WithFlagStackRef, cmd)
	common.RequiredFlag(WithFlagConfigRepository, cmd)
	cmd.Flags().StringP("env", "e", "", "[deprecated] add an environment with project creation")
	cmd.Flags().MarkHidden("env")
	cmd.Flags().String("vars", "", "[deprecated] path to a variable file for the env creation")
	cmd.Flags().MarkHidden("vars")
	cmd.Flags().String("pipeline", "", "[deprecated] path to a pipeline file for the env creation")
	cmd.Flags().MarkHidden("pipeline")
	cmd.Flags().StringToString("config", nil, "[deprecated] path to a config file to inject in the config repo")
	cmd.Flags().MarkHidden("config")
	cmd.Flags().String("usecase", "", "[deprecated] the usecase for the env creation")
	cmd.Flags().MarkHidden("usecase")
	cmd.Flags().String("owner", "", "the owner username")

	WithFlagDescription(cmd)
	WithFlagCanonical(cmd)
	return cmd
}

func create(cmd *cobra.Command, args []string) error {
	api := common.NewAPI()
	m := middleware.NewMiddleware(api)

	org, err := common.GetOrg(cmd)
	if err != nil {
		return err
	}

	name, err := cmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	canonical, err := cmd.Flags().GetString("canonical")
	if err != nil {
		return err
	}

	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return err
	}

	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return err
	}

	configRepo, err := cmd.Flags().GetString("config-repo")
	if err != nil {
		return err
	}

	ownerCanonical, err := cmd.Flags().GetString("owner")
	if err != nil {
		return err
	}

	output, err := cmd.Flags().GetString("output")
	if err != nil {
		return errors.Wrap(err, "unable to get output flag")
	}

	// fetch the printer from the factory
	p, err := factory.GetPrinter(output)
	if err != nil {
		return errors.Wrap(err, "unable to get printer")
	}

	// Deprecated flags, we still collect them to give a correct error message to the user
	_, err = cmd.Flags().GetStringToString("config")
	if err != nil {
		return err
	}

	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	usecase, err := cmd.Flags().GetString("usecase")
	if err != nil {
		return err
	}

	varsPath, err := cmd.Flags().GetString("vars")
	if err != nil {
		return err
	}

	pipelinePath, err := cmd.Flags().GetString("pipeline")
	if err != nil {
		return err
	}

	// Handle deprecation
	if env+usecase+varsPath+pipelinePath != "" {
		// If any of the env provisioning vars is not empty, create the project with an env
		return errors.New("Creating an environment when creating a project is not possible anymore. Please create your env separately using the 'cy project create-env' command.")
	}

	project, err := m.CreateProject(org, name, canonical, description, stackRef, configRepo, ownerCanonical)
	return printer.SmartPrint(p, project, err, "", printer.Options{}, cmd.OutOrStdout())
}
