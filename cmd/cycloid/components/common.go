package components

import "github.com/spf13/cobra"

func addComponentNameFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("name", "n", "", "set the name of the component")
}

func getComponentName(cmd *cobra.Command) (*string, error) {
	componentName, err := cmd.Flags().GetString("name")
	if err != nil {
		return nil, err
	}

	return &componentName, nil
}

func addDescriptionFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("description", "d", "", "set the description of the component")
}

func getDescription(cmd *cobra.Command) (*string, error) {
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, err
	}

	return &description, nil
}

func addCloudProviderFlag(cmd *cobra.Command) {
	cmd.Flags().String("cloud-provider", "", "set the cloud provider of the component")
}

func getCloudProvider(cmd *cobra.Command) (*string, error) {
	cloudProvider, err := cmd.Flags().GetString("cloud-provider")
	if err != nil {
		return nil, err
	}

	return &cloudProvider, nil
}

func addStackRefFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("stack-ref", "s", "", "set the stack ref of the component in format org:stack-canonical")
	cmd.MarkFlagRequired("stack-ref")
}

func getStackRef(cmd *cobra.Command) (*string, error) {
	stackRef, err := cmd.Flags().GetString("stack-ref")
	if err != nil {
		return nil, err
	}

	return &stackRef, nil
}

func addUseCaseFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("use-case", "u", "", "set the use-case of the component")
	cmd.MarkFlagRequired("use-case")
}

func getUseCase(cmd *cobra.Command) (*string, error) {
	useCase, err := cmd.Flags().GetString("use-case")
	if err != nil {
		return nil, err
	}

	return &useCase, nil
}
