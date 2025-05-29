package pipelines

import (
	"github.com/spf13/cobra"

	"github.com/cycloidio/cycloid-cli/cmd/cycloid/common"
	"github.com/cycloidio/cycloid-cli/cmd/cycloid/internal"
)

func NewDiffCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:    "diff",
		Hidden: true,
		Short:  "not implemented yet",
		Long:   `not implemented yet`,
		// RunE:    diff,
		RunE:    func(cmd *cobra.Command, args []string) error { panic("TODO: not implemented") },
		PreRunE: internal.CheckAPIAndCLIVersion,
	}

	common.RequiredPersistentFlag(common.WithFlagProject, cmd)
	common.RequiredPersistentFlag(common.WithFlagEnv, cmd)
	common.RequiredFlag(WithFlagPipeline, cmd)
	common.RequiredFlag(WithFlagVars, cmd)

	return cmd
}

// func diff(cmd *cobra.Command, args []string) error {
// 	api := common.NewAPI()
// 	m := middleware.NewMiddleware(api)
//
// 	var err error
//
// 	org, err := cy_args.GetOrg(cmd)
// 	if err != nil {
// 		return err
// 	}
// 	project, err := cmd.Flags().GetString("project")
// 	if err != nil {
// 		return err
// 	}
// 	env, err := cmd.Flags().GetString("env")
// 	if err != nil {
// 		return err
// 	}
// 	varsPath, err := cmd.Flags().GetString("vars")
// 	if err != nil {
// 		return err
// 	}
// 	pipelinePath, err := cmd.Flags().GetString("pipeline")
// 	if err != nil {
// 		return err
// 	}
// 	output, err := cmd.Flags().GetString("output")
// 	if err != nil {
// 		return errors.Wrap(err, "unable to get output flag")
// 	}
//
// 	// fetch the printer from the factory
// 	p, err := factory.GetPrinter(output)
// 	if err != nil {
// 		return errors.Wrap(err, "unable to get printer")
// 	}
//
// 	rawPipeline, err := os.ReadFile(pipelinePath)
// 	if err != nil {
// 		return fmt.Errorf("Pipeline file reading error : %s", err.Error())
// 	}
// 	pipelineTemplate := string(rawPipeline)
//
// 	rawVars, err := os.ReadFile(varsPath)
// 	if err != nil {
// 		return fmt.Errorf("Pipeline variables file reading error : %s", err.Error())
// 	}
// 	vars := string(rawVars)
//
// 	pd, err := m.DiffPipeline(org, project, env, pipelineTemplate, vars)
// 	if err != nil {
// 		// print the result on the standard output
// 		if err := p.Print(err, printer.Options{}, cmd.OutOrStdout()); err != nil {
// 			return errors.Wrap(err, "unable to print result")
// 		}
// 		return errors.Wrap(err, "unable obtain a diff")
// 	}
//
// 	// print the result on the standard output
// 	if err := p.Print(pd, printer.Options{}, cmd.OutOrStdout()); err != nil {
// 		return errors.Wrap(err, "unable to print result")
// 	}
//
// 	return nil
// }
